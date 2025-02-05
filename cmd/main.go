package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sampathreddy22/task-management-api/internal/config"
	"github.com/sampathreddy22/task-management-api/internal/handlers"
	"github.com/sampathreddy22/task-management-api/internal/repositories"
	"github.com/sampathreddy22/task-management-api/internal/services"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	//Initialize repositories and services
	taskRepo := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	//Setup routes
	router.POST("/tasks", taskHandler.CreateTask)
	router.GET("/tasks/:id", taskHandler.GetTaskByID)
	router.PUT("/tasks/:id", taskHandler.UpdateTask)
	router.DELETE("/tasks/:id", taskHandler.DeleteTask)
	router.GET("/tasks", taskHandler.GetTasks)

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
