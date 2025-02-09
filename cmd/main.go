// @title Task Management API
// @version 1.0
// @description A simple task management API
// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	docs "github.com/sampathreddy22/task-management-api/cmd/docs"
	"github.com/sampathreddy22/task-management-api/internal/config"
	"github.com/sampathreddy22/task-management-api/internal/handlers"
	"github.com/sampathreddy22/task-management-api/internal/repositories"
	"github.com/sampathreddy22/task-management-api/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	api := router.Group("/api/v1")

	//Initialize repositories and services
	taskRepo := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	attachmentRepo := repositories.NewAttachmentRepository(db)
	attachmentService := services.NewAttachmentService(attachmentRepo)
	attachmentHandler := handlers.NewAttachmentHandler(attachmentService)

	//Setup routes
	tasks := api.Group("/tasks")
	{
		tasks.POST("/", taskHandler.CreateTask)
		tasks.GET("/:id", taskHandler.GetTaskByID)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.DeleteTask)
		tasks.GET("/", taskHandler.GetTasks)
	}

	users := api.Group("/users")
	{
		users.POST("/", userHandler.CreateUser)
		users.GET("/:id", userHandler.GetUserByID)
	}

	attachments := api.Group("/attachments")
	{
		attachments.POST("/", attachmentHandler.CreateAttachment)
		attachments.GET("/:id", attachmentHandler.GetAttachment)
		attachments.GET("/task/:taskId", attachmentHandler.GetTaskAttachments)
		attachments.PUT("/:id", attachmentHandler.UpdateAttachment)
		attachments.DELETE("/:id", attachmentHandler.DeleteAttachment)
	}

	//Add swagger documentation
	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DocExpansion("none")))

	return router

}

func main() {
	//Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	//Initialize database
	db, err := config.InitializeDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// setup router with database connectivity
	router := setupRouter(db)

	// Start the server
	port := cfg.Server.Port
	if port == "" {
		port = "8080"
	}

	router.Run(fmt.Sprintf(":%s", port))

}
