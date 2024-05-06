package services

import (
	"testing"

	"github.com/AI-StartUps/user-management-service/config"
	"github.com/AI-StartUps/user-management-service/internal/adapter/logger"
	"github.com/AI-StartUps/user-management-service/internal/adapter/repository"
)

type testingservices struct {
	userService     UserService
	roleService     RoleService
	userRoleService UserRoleService
}

func initServices() testingservices {
	logger, err := logger.NewDefaultLogger()
	if err != nil {
		panic(err)
	}

	config, err := config.NewConfig(logger)
	if err != nil {
		panic(err)
	}

	roleRepo, _ := repository.NewRolePostgresClient(*config, logger)
	userRepo, _ := repository.NewUserPostgresClient(*config, logger)
	userRoleRepo, _ := repository.NewUserRolePostgresClient(*config, logger)

	userService := NewUserService(userRepo)
	roleService := NewRoleService(roleRepo)
	userRoleService := NewUserRoleService(userRoleRepo)

	return testingservices{
		userService:     *userService,
		roleService:     *roleService,
		userRoleService: *userRoleService,
	}

}

func TestUserManagementService(t *testing.T) {
	// 1. Testing CreateUser
	t.Run("Testing CreateUser", func(t *testing.T) {

	})
	// 2. Testing GetUsers
	t.Run("Testing GetUsers", func(t *testing.T) {

	})
	// 3. Testing GetUsersWithRole
	t.Run("Testing GetUsersWithRole", func(t *testing.T) {

	})
	// 4. Testing GetUserById
	t.Run("Testing GetUserById", func(t *testing.T) {

	})
	// 5. Testing UpdateUser
	t.Run("Testing UpdateUser", func(t *testing.T) {

	})
	// 6. Testing DeleteUser
	t.Run("Testing DeleteUser", func(t *testing.T) {

	})

}
