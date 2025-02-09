package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sampathreddy22/task-management-api/internal/models"
	"github.com/sampathreddy22/task-management-api/internal/services"
)

type AttachmentHandler struct {
	service *services.AttachmentService
}

func NewAttachmentHandler(service *services.AttachmentService) *AttachmentHandler {
	return &AttachmentHandler{service: service}
}

// @Summary Create a new attachment
// @Description Create a new attachment with the given input
// @Accept json
// @Produce json
// @Param attachment body models.AttachmentInput true "The input for the attachment"
// @Success 201 {object} models.Attachment
// @Router /attachments [post]
// @tags attachments
func (h *AttachmentHandler) CreateAttachment(c *gin.Context) {
	var input models.AttachmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attachment, err := h.service.CreateAttachment(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, attachment)
}

// @Summary Get an attachment by ID
// @Description Get an attachment by ID
// @Accept json
// @Produce json
// @Param id path string true "The ID of the attachment to get"
// @Success 200 {object} models.Attachment
// @Router /attachments/{id} [get]
// @tags attachments
func (h *AttachmentHandler) GetAttachment(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	attachment, err := h.service.GetAttachment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attachment not found"})
		return
	}

	c.JSON(http.StatusOK, attachment)
}

// @Summary Get all attachments for a task
// @Description Get all attachments for a task
// @Accept json
// @Produce json
// @Param taskId path string true "The ID of the task to get attachments for"
// @Success 200 {object} []models.Attachment
// @Router /attachments/task/{taskId} [get]
// @tags attachments
func (h *AttachmentHandler) GetTaskAttachments(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("taskId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}

	attachments, err := h.service.GetTaskAttachments(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attachments)
}

// @Summary Update an attachment
// @Description Update an attachment with the given input
// @Accept json
// @Produce json
// @Param id path string true "The ID of the attachment to update"
// @Param attachment body models.AttachmentInput true "The input for the attachment"
// @Success 200 {object} models.Attachment
// @Router /attachments/{id} [put]
// @tags attachments
func (h *AttachmentHandler) UpdateAttachment(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input models.AttachmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attachment, err := h.service.UpdateAttachment(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attachment)
}

// @Summary Delete an attachment
// @Description Delete an attachment with the given ID
// @Accept json
// @Produce json
// @Param id path string true "The ID of the attachment to delete"
// @Success 200
// @Router /attachments/{id} [delete]
// @tags attachments
func (h *AttachmentHandler) DeleteAttachment(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.service.DeleteAttachment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attachment deleted successfully"})
}
