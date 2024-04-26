package app

import (
	"net/http"

	"github.com/AI-StartUps/user-management-service/internal/core/domain"
	"github.com/AI-StartUps/user-management-service/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type GinHandler interface {
	CreateUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUsersWithRole(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	CreateRole(ctx *gin.Context)
	GetRoleById(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)
	AddUserRole(ctx *gin.Context)
	RemoveUserRole(ctx *gin.Context)
}

type handler struct {
	userService     ports.UserService
	roleService     ports.RoleService
	userRoleService ports.UserRoleService
}

func NewGinHandler(userService ports.UserService, roleService ports.RoleService, userRoleService ports.UserRoleService) GinHandler {
	routerHandler := handler{
		userService:     userService,
		roleService:     roleService,
		userRoleService: userRoleService,
	}
	return routerHandler
}

func (h handler) CreateUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h handler) GetUsers(ctx *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h handler) GetUsersWithRole(ctx *gin.Context) {
	roleName := ctx.Param("role_name")
	users, err := h.userService.GetUsersWithRole(roleName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h handler) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	user, err := h.userService.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h handler) UpdateUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h handler) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	err := h.userService.DeleteUser(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h handler) CreateRole(ctx *gin.Context) {
	var role domain.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.roleService.CreateRole(role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h handler) GetRoleById(ctx *gin.Context) {
	roleID := ctx.Param("roleId")
	role, err := h.roleService.GetRoleById(roleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if role == nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, role)
}

func (h handler) UpdateRole(ctx *gin.Context) {
	roleID := ctx.Param("roleId")
	var role domain.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role.RoleId = roleID
	if err := h.roleService.UpdateRole(role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h handler) DeleteRole(ctx *gin.Context) {
	roleID := ctx.Param("roleId")
	if err := h.roleService.DeleteRole(roleID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h handler) AddUserRole(ctx *gin.Context) {
	var userRole domain.UserRole
	if err := ctx.ShouldBindJSON(&userRole); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userRoleService.AddUserRole(userRole)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h handler) RemoveUserRole(ctx *gin.Context) {
	var userRole domain.UserRole
	if err := ctx.ShouldBindJSON(&userRole); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userRoleService.RemoveUserRole(userRole)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
