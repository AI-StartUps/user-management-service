package repository

import (
	"database/sql"
	"fmt"

	"github.com/AI-StartUps/user-management-service/config"
	"github.com/AI-StartUps/user-management-service/internal/core/domain"
	"github.com/AI-StartUps/user-management-service/internal/core/ports"
	_ "github.com/lib/pq"
)

type postgresClient struct {
	db        *sql.DB
	logger    ports.LoggerService
	tablename string
}

func NewUserPostgresClient(config config.Config, logger ports.LoggerService) (*postgresClient, error) {
	dbname := config.POSTGRES_DB
	tablename := config.USER_TABLE
	user := config.POSTGRES_USER
	password := config.POSTGRES_PASSWORD
	port := config.POSTGRES_PORT
	host := config.POSTGRES_HOST

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to the database: %v", err))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to ping the database: %v", err))
		return nil, err
	}

	queryString := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            user_id VARCHAR(255) PRIMARY KEY UNIQUE,
            username VARCHAR(255) NOT NULL,
            password_hash VARCHAR(255) NOT NULL,
            email VARCHAR(255) NOT NULL,
            fullname VARCHAR(255) NOT NULL,
            phone_number VARCHAR(255),
            avatar VARCHAR(255),
            address VARCHAR(255),
            created_at TIMESTAMP,
            updated_at TIMESTAMP
        )
    `, tablename)

	_, err = db.Exec(queryString)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create user table: %v", err))
		return nil, err
	}
	logger.Info("Connected to the database successfully")
	return &postgresClient{db: db, tablename: tablename, logger: logger}, nil
}

func NewRolePostgresClient(config config.Config, logger ports.LoggerService) (*postgresClient, error) {
	dbname := config.POSTGRES_DB
	tablename := config.ROLE_TABLE
	user := config.POSTGRES_USER
	password := config.POSTGRES_PASSWORD
	port := config.POSTGRES_PORT
	host := config.POSTGRES_HOST

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to the database: %v", err))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to ping the database: %v", err))
		return nil, err
	}

	roleQueryString := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            role_id VARCHAR(255) PRIMARY KEY UNIQUE,
            name VARCHAR(255) NOT NULL,
            description VARCHAR(255)
        )
    `, tablename)

	_, err = db.Exec(roleQueryString)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create role table: %v", err))
		return nil, err
	}
	logger.Info("Connected to the database successfully")
	return &postgresClient{db: db, tablename: tablename, logger: logger}, nil
}

func NewUserRolePostgresClient(config config.Config, logger ports.LoggerService) (*postgresClient, error) {
	dbname := config.POSTGRES_DB
	tablename := config.USER_ROLE_TABLE
	user := config.POSTGRES_USER
	password := config.POSTGRES_PASSWORD
	port := config.POSTGRES_PORT
	host := config.POSTGRES_HOST

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to the database: %v", err))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to ping the database: %v", err))
		return nil, err
	}

	// Create the roles table if it doesn't exist
	roleQueryString := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS roles (
            role_id VARCHAR(255) PRIMARY KEY UNIQUE,
            name VARCHAR(255) NOT NULL,
            description VARCHAR(255)
        )
    `)
	_, err = db.Exec(roleQueryString)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create roles table: %v", err))
		return nil, err
	}

	userRoleQueryString := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            user_id VARCHAR(255),
            role_id VARCHAR(255),
            PRIMARY KEY (user_id, role_id),
            CONSTRAINT fk_user_roles_user_id FOREIGN KEY (user_id) REFERENCES %s(user_id),
            CONSTRAINT fk_user_roles_role_id FOREIGN KEY (role_id) REFERENCES roles(role_id)
        )
    `, tablename, config.USER_TABLE)

	_, err = db.Exec(userRoleQueryString)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create user_role table: %v", err))
		return nil, err
	}

	alterQueryString := fmt.Sprintf(`
        ALTER TABLE %s
        ADD CONSTRAINT unique_user_role UNIQUE (user_id, role_id);
    `, tablename)

	_, err = db.Exec(alterQueryString)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to add unique constraint to user_role table: %v", err))
		return nil, err
	}

	logger.Info("Connected to the database successfully")
	return &postgresClient{db: db, tablename: tablename, logger: logger}, nil
}

func (svc postgresClient) CreateUser(user domain.User) error {
	query := `
        INSERT INTO users (user_id, username, password_hash, email, fullname, phone_number, avatar, address, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    `
	_, err := svc.db.Exec(query,
		user.UserId,
		user.Username,
		user.PasswordHash,
		user.Email,
		user.FullName,
		user.PhoneNumber,
		user.Avatar,
		user.Address,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to create user: %s`, err.Error()))
		return err
	}
	return nil
}

func (svc postgresClient) GetUserById(userId string) (*domain.User, error) {
	query := `
        SELECT user_id, username, password_hash, email, fullname, phone_number, avatar, address, created_at, updated_at
        FROM users
        WHERE user_id = $1
    `
	row := svc.db.QueryRow(query, userId)
	user := &domain.User{}
	err := row.Scan(&user.UserId, &user.Username, &user.PasswordHash, &user.Email, &user.FullName, &user.PhoneNumber, &user.Avatar, &user.Address, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to get user: %s`, err.Error()))
		return nil, err
	}
	return user, nil
}

func (svc postgresClient) GetUsers() ([]*domain.User, error) {
	query := `
        SELECT user_id, username, password_hash, email, fullname, phone_number, avatar, address, created_at, updated_at
        FROM users
    `
	rows, err := svc.db.Query(query)
	if err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to get users: %s`, err.Error()))
		return nil, err
	}
	defer rows.Close()

	users := []*domain.User{}
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.UserId, &user.Username, &user.PasswordHash, &user.Email, &user.FullName, &user.PhoneNumber, &user.Avatar, &user.Address, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			svc.logger.Error(fmt.Sprintf(`Unable to get users: %s`, err.Error()))
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to get users: %s`, err.Error()))
		return nil, err
	}

	return users, nil
}

func (svc postgresClient) GetUsersWithRole(roleName string) ([]*domain.User, error) {
	query := `
        SELECT u.user_id, u.username, u.password_hash, u.email, u.fullname, u.phone_number, u.avatar, u.address, u.created_at, u.updated_at
        FROM users u
        JOIN user_roles ur ON u.user_id = ur.user_id
        JOIN roles r ON ur.role_id = r.role_id
        WHERE r.name = $1
    `
	rows, err := svc.db.Query(query, roleName)
	if err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to get users with role %s: %s`, roleName, err.Error()))
		return nil, err
	}
	defer rows.Close()

	users := []*domain.User{}
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.UserId, &user.Username, &user.PasswordHash, &user.Email, &user.FullName, &user.PhoneNumber, &user.Avatar, &user.Address, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			svc.logger.Error(fmt.Sprintf(`Unable to get users with role %s: %s`, roleName, err.Error()))
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to get users with role %s: %s`, roleName, err.Error()))
		return nil, err
	}

	return users, nil
}

func (svc postgresClient) UpdateUser(user domain.User) error {
	query := `
        UPDATE users
        SET username=$2, password_hash=$3, email=$4, fullname=$5, phone_number=$6, avatar=$7, address=$8, updated_at=$9
        WHERE user_id=$1
    `
	_, err := svc.db.Exec(query, user.UserId, user.Username, user.PasswordHash, user.Email, user.FullName, user.PhoneNumber, user.Avatar, user.Address, user.UpdatedAt)
	if err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to update user: %s`, err.Error()))
		return err
	}
	return nil
}

func (svc postgresClient) DeleteUser(userId string) error {
	query := `
        DELETE FROM users
        WHERE user_id=$1
    `
	_, err := svc.db.Exec(query, userId)
	if err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to delete user: %s`, err.Error()))
		return err
	}
	return nil
}

func (svc postgresClient) CreateRole(role domain.Role) error {
	query := `
        INSERT INTO roles (role_id, name, description)
        VALUES ($1, $2, $3)
    `
	_, err := svc.db.Exec(query, role.RoleId, role.Name, role.Description)
	if err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to create role: %s`, err.Error()))
		return err
	}
	return nil
}

func (svc postgresClient) GetRoleById(roleId string) (*domain.Role, error) {
	query := `
        SELECT role_id, name, description
        FROM roles
        WHERE role_id = $1
    `
	row := svc.db.QueryRow(query, roleId)
	role := &domain.Role{}
	err := row.Scan(&role.RoleId, &role.Name, &role.Description)
	if err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to get role: %s`, err.Error()))
		return nil, err
	}
	return role, nil
}

func (svc postgresClient) UpdateRole(role domain.Role) error {
	query := `
        UPDATE roles
        SET name=$2, description=$3
        WHERE role_id=$1
    `
	_, err := svc.db.Exec(query, role.RoleId, role.Name, role.Description)
	if err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to update role: %s`, err.Error()))
		return err
	}
	return nil
}

func (svc postgresClient) DeleteRole(roleId string) error {
	query := `
        DELETE FROM roles
        WHERE role_id=$1
    `
	_, err := svc.db.Exec(query, roleId)
	if err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to delete role: %s`, err.Error()))
		return err
	}
	return nil
}

func (svc postgresClient) AddUserRole(userRole domain.UserRole) error {
	query := `
        INSERT INTO user_roles (user_id, role_id)
        VALUES ($1, $2)
    `
	_, err := svc.db.Exec(query, userRole.UserId, userRole.RoleId)
	if err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to add user role: %s`, err.Error()))
		return err
	}
	return nil
}

func (svc postgresClient) RemoveUserRole(userRole domain.UserRole) error {
	query := `
        DELETE FROM user_roles
        WHERE user_id=$1 AND role_id=$2
    `
	_, err := svc.db.Exec(query, userRole.UserId, userRole.RoleId)
	if err != nil {
		svc.logger.Error(fmt.Sprintf(`Unable to remove user role: %s`, err.Error()))
		return err
	}
	return nil
}
