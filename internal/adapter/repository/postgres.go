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
	

	userRoleTablename := config.USER_ROLE_TABLE
	userRoleQueryString := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            user_id VARCHAR(255),
            role_id VARCHAR(255),
            PRIMARY KEY (user_id, role_id),
            FOREIGN KEY (user_id) REFERENCES %s(user_id),
            FOREIGN KEY (role_id) REFERENCES %s(role_id)
        )
    `, userRoleTablename, tablename, tablename)

	_, err = db.Exec(userRoleQueryString)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create user_role table: %v", err))
		return nil, err
	}

	logger.Info("Connected to the database successfully")
	return &postgresClient{db: db, tablename: tablename, logger: logger}, nil
}

func (svc postgresClient) CreateUser(user *domain.User) error {
	return nil
}

func (svc postgresClient) GetUserById(userId string) (*domain.User, error) {
	return nil, nil
}

func (svc postgresClient) UpdateUser(user domain.User) error {
	return nil
}

func (svc postgresClient) DeleteUser(userId string) error {
	return nil
}

func (svc postgresClient) CreateRole(role *domain.Role) error {
	return nil
}

func (svc postgresClient) GetRoleById(roleId string) (*domain.Role, error) {
	return nil, nil
}

func (svc postgresClient) UpdateRole(role domain.Role) error {
	return nil
}

func (svc postgresClient) DeleteRole(roleId string) error {
	return nil
}

func (svc postgresClient) AddUserRole(userRole domain.UserRole) error {
	return nil
}

func (svc postgresClient) RemoveUserRole(userRole domain.UserRole) error {
	return nil
}
