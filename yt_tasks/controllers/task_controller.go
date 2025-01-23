package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"yt_tasks/config"
	"yt_tasks/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// create task
func CreateTask(c *gin.Context) {
	var task models.Task

	taskName := c.PostForm("name")
	if taskName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task name is required"})
		return
	}

	taskDescription := c.PostForm("description")

	category := c.PostForm("category")
	if !isValidCategory(category) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	priority := c.PostForm("priority")
	if !isValidPriority(priority) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid priority"})
		return
	}

	status := c.PostForm("status")
	if !isValidStatus(status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	const dateLayout = "02/01/2006"
	dueDate := c.PostForm("due_date")
	dueDateParsed, err := time.Parse(dateLayout, dueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due date"})
		return
	}

	userName := c.PostForm("username")
	if userName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Assignee name is required"})
		return
	}

	var image *models.Image
	file, err := c.FormFile("image")
	if err == nil {
		uploadDir := "uploads"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		filePath := filepath.Join(uploadDir, file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save metadata to server"})
			return
		}

		image = &models.Image{
			Filename: file.Filename,
			FilePath: filePath,
		}
	} else {
		image = &models.Image{
			Filename: "",
			FilePath: "",
		}
	}

	task.Id = uuid.New().String()
	task.Name = taskName
	task.Description = taskDescription
	task.Cateogry = category
	task.Priority = priority
	task.Status = status
	task.Assignee = models.Assignee{
		Username: userName,
		Image:    image,
	}
	task.DueDate = dueDateParsed

	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "New task created", "data": task})
}

// get tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task

	if err := config.DB.Order("created_at desc").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": fmt.Sprintf("%d items", len(tasks)), "data": tasks})
}

// get task by id
func GetTaskDetails(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := config.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// update task
func UpdateTaskDetails(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := config.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	taskName := c.PostForm("name")
	if taskName != "" {
		task.Name = taskName
	}

	taskDescription := c.PostForm("description")
	if taskDescription != "" {
		task.Description = taskDescription
	}

	category := c.PostForm("category")
	if category != "" {
		if !isValidCategory(category) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
			return
		}
		task.Cateogry = category
	}

	priority := c.PostForm("priority")
	if priority != "" {
		if !isValidPriority(priority) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid priority"})
			return
		}
		task.Priority = priority
	}

	status := c.PostForm("status")
	if status != "" {
		if !isValidStatus(status) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
			return
		}
		task.Status = status
	}

	const dateLayout = "02/01/2006"
	dueDate := c.PostForm("due_date")
	if dueDate != "" {
		dueDateParsed, err := time.Parse(dateLayout, dueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due date"})
			return
		}
		task.DueDate = dueDateParsed
	}

	username := c.PostForm("username")
	if username != "" {
		task.Assignee.Username = username
	}

	file, err := c.FormFile("image")
	if err == nil {
		// remove old file
		if err := os.Remove(task.Assignee.Image.FilePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old image"})
			return
		}

		// save new file
		uploadDir := "uploads"
		filePath := filepath.Join(uploadDir, file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save new file"})
			return
		}

		task.Assignee.Image = &models.Image{
			Filename: file.Filename,
			FilePath: filePath,
		}
	}

	// update task
	if err := config.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated", "data": task})
}

// delete task
func DeleteTaskDetails(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := config.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// remove assignee image from server
	if task.Assignee.Image != nil && task.Assignee.Image.FilePath != "" {
		if err := os.Remove(task.Assignee.Image.FilePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image from server"})
			return
		}
	}

	if err := config.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
