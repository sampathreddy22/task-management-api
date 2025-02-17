package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sampathreddy22/task-management-api/internal/models"
	"github.com/sampathreddy22/task-management-api/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// @Summary Create a new user
// @Description Create a new user with the given name and email
// @Accept json
// @Produce json
// @Param user body models.UserInput true "UserInput to create user"
// @Success 201 {object} models.User
// @Router /users [post]
// @tags users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var userInput models.UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	user.FromUserInput(userInput)

	if err := h.userService.CreateUser(context.Background(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "User created successfully")
}

// @Summary Get a user by ID
// @Description Get a user by ID
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
// @tags users
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userService.GetUserByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
