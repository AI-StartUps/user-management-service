package services

import (
	"time"

	"github.com/AI-StartUps/user-management-service/internal/core/domain"
	"github.com/AI-StartUps/user-management-service/internal/core/ports"
	"github.com/google/uuid"
)

type userService struct {
	repo ports.UserRepository
}

type roleService struct {
	repo ports.RoleRepository
}

type userRoleService struct {
	repo ports.UserRoleRepository
}

func NewUserService(repo ports.UserRepository) *userService {
	service := userService{
		repo: repo,
	}
	return &service
}

func NewRoleService(repo ports.RoleRepository) *roleService {
	service := roleService{
		repo: repo,
	}
	return &service
}

func NewUserRoleService(repo ports.UserRoleRepository) *userRoleService {
	service := userRoleService{
		repo: repo,
	}
	return &service
}

func (svc userService) CreateUser(user *domain.User) error {
	user.UserId = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return svc.repo.CreateUser(*user)
}

func (svc userService) GetUserById(userId string) (*domain.User, error) {
	return svc.repo.GetUserById(userId)
}

func (svc userService) GetUsers() ([]domain.User, error) {
	return svc.repo.GetUsers()
}
func (svc userService) GetUsersWithRole(roleName string) ([]*domain.User, error) {
	return svc.repo.GetUsersWithRole(roleName)
}

func (svc userService) UpdateUser(user domain.User) error {
	return svc.repo.UpdateUser(user)
}

func (svc userService) DeleteUser(userId string) error {
	return svc.repo.DeleteUser(userId)
}

func (svc roleService) CreateRole(role *domain.Role) error {
	role.RoleId = uuid.New().String()
	return svc.repo.CreateRole(*role)
}

func (svc roleService) GetRoleById(roleId string) (*domain.Role, error) {
	return svc.repo.GetRoleById(roleId)
}

func (svc roleService) UpdateRole(role domain.Role) error {
	return svc.repo.UpdateRole(role)
}

func (svc roleService) DeleteRole(roleId string) error {
	return svc.repo.DeleteRole(roleId)
}

func (svc userRoleService) AddUserRole(userRole domain.UserRole) error {
	return svc.repo.AddUserRole(userRole)
}

func (svc userRoleService) RemoveUserRole(userRole domain.UserRole) error {
	return svc.repo.RemoveUserRole(userRole)
}
