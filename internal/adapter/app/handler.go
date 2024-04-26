package app

import (
	"fmt"
	"log"

	"github.com/AI-StartUps/user-management-service/config"
	"github.com/AI-StartUps/user-management-service/internal/core/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitGinRoutes(userService ports.UserService, roleService ports.RoleService, userRoleService ports.UserRoleService, config config.Config) {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	handler := NewGinHandler(
		userService,
		roleService,
		userRoleService,
	)

	userRoutes := router.Group("/users/v1")
	{
		userRoutes.POST("/", handler.CreateUser)
		userRoutes.GET("/:user_id", handler.GetUserById)
		userRoutes.GET("/", handler.GetUsers)
		userRoutes.GET("/roles/:role_name", handler.GetUsersWithRole)
		userRoutes.PUT(":user_id", handler.UpdateUser)
		userRoutes.DELETE("/:user_id", handler.DeleteUser)
	}
	roleRoutes := router.Group("/roles/v1")
	{
		roleRoutes.POST("/", handler.CreateRole)
		roleRoutes.GET("/:role_id", handler.GetRoleById)
		roleRoutes.PUT(":role_id", handler.UpdateRole)
		roleRoutes.DELETE("/:role_id", handler.DeleteRole)
	}

	userRoleRoutes := router.Group("/user_roles/v1")
	{
		userRoleRoutes.POST("/", handler.AddUserRole)
		userRoleRoutes.GET("/:role_id", handler.RemoveUserRole)
	}
	log.Printf("Server running on port 0.0.0.0:%s", config.SERVER_PORT)
	router.Run(fmt.Sprintf(":%s", config.SERVER_PORT))
}
