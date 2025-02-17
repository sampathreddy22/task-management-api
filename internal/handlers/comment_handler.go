package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sampathreddy22/task-management-api/internal/models"
	"github.com/sampathreddy22/task-management-api/internal/services"
)

type CommentHandler struct {
	commentService *services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

// @Summary Add a new comment to a task
// @Description Add a new comment to a task
// @Accept json
// @Produce json
// @Param comment body models.Comment true "Comment to add"
// @Success 201 {object} models.Comment
// @Router /comments [post]
// @tags comments
func (h *CommentHandler) AddComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment.ID = uuid.New()
	comment.CreatedAt = time.Now()

	if err := h.commentService.CreateComment(c.Request.Context(), &comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// @Summary Get comments by task ID
// @Description Get comments by task ID
// @Accept json
// @Produce json
// @Param taskID path string true "Task ID"
// @Success 200 {array} models.Comment
// @Router /comments/{taskID} [get]
// @tags comments
func (h *CommentHandler) GetCommentsByTaskID(c *gin.Context) {
	taskID := c.Param("taskID")
	comments, err := h.commentService.GetCommentsByTaskID(c.Request.Context(), taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	if err := h.commentService.DeleteComment(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
