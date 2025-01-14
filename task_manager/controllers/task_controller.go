package controllers

import (
	"fmt"
	"net/http"
	"task_manager/config"
	"task_manager/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// create task
func CreateTask(c *gin.Context) {
	var task models.Task

	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	description := c.PostForm("description")

	status := c.PostForm("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
		return
	}

	if !isValidStatus(status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	category := c.PostForm("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
		return
	}

	if !isValidCategory(category) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	startDate := c.PostForm("start_date")
	dueDate := c.PostForm("due_date")

	const dateLayout = "02-01-2006"
	startDateParsed, err := time.Parse(dateLayout, startDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	dueDateParsed, err := time.Parse(dateLayout, dueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due date format"})
		return
	}

	task.Id = uuid.New().String()
	task.Title = title
	task.Description = description
	task.Status = status
	task.Category = category
	task.StartDate = startDateParsed
	task.DueDate = dueDateParsed

	if err := config.TaskDB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task Created", "data": task})
}

// get tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task

	if err := config.TaskDB.Order("created_at desc").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": fmt.Sprintf("%d items", len(tasks)), "data": tasks})
}

// get task by id
func GetTaskDetails(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := config.TaskDB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}