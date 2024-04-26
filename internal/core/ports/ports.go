package ports

import (
	"github.com/AI-StartUps/user-management-service/internal/core/domain"
)

type UserService interface {
	CreateUser(user domain.User) error
	GetUsers() ([]domain.User, error)
	GetUsersWithRole(roleName string) ([]*domain.User, error)
	GetUserById(userId string) (*domain.User, error)
	UpdateUser(user domain.User) error
	DeleteUser(userId string) error
}

type RoleService interface {
	CreateRole(role domain.Role) error
	GetRoleById(roleId string) (*domain.Role, error)
	UpdateRole(role domain.Role) error
	DeleteRole(roleId string) error
}

type UserRoleService interface {
	AddUserRole(userRole domain.UserRole) error
	RemoveUserRole(userRole domain.UserRole) error
}

type UserRepository interface {
	CreateUser(user domain.User) error
	GetUserById(userId string) (*domain.User, error)
	GetUsers() ([]domain.User, error)
	GetUsersWithRole(roleName string) ([]*domain.User, error)
	UpdateUser(user domain.User) error
	DeleteUser(userId string) error
}

type RoleRepository interface {
	CreateRole(role domain.Role) error
	GetRoleById(roleId string) (*domain.Role, error)
	UpdateRole(role domain.Role) error
	DeleteRole(roleId string) error
}

type UserRoleRepository interface {
	AddUserRole(userRole domain.UserRole) error
	RemoveUserRole(userRole domain.UserRole) error
}

type LoggerService interface {
	Info(message string)
	Warning(message string)
	Error(message string)
}
