package controllers

import (
	"learning/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// create learning [dashboard:postman]
func CreateLearning(c *gin.Context) {
	var learning models.Learning

	// module no
	moduleNo := c.PostForm("module_no")
	if moduleNo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Module number is required"})
		return
	}

	moduleNoParsed, err := strconv.ParseInt(moduleNo, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Module number must be a valid number"})
		return
	}

	if moduleNoParsed <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Module number must be greater than zero"})
		return
	}

	// module name
	moduleName := c.PostForm("module_name")
	if moduleName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Module name is required"})
		return
	}

	// total duration
	totalDuration := c.PostForm("total_duration")
	if totalDuration == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Total duration is required"})
		return
	}

	totalDurationParsed, err := strconv.ParseFloat(totalDuration, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duration"})
		return
	}

	// category
	category := c.PostForm("category")
	if !isValidCategory(category) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}
	
	// thumbnail image
	thumbnailImage, err := c.FormFile("thumbnail_url")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Thumbnail url is required"})
		return
	}

	uploadDir := "uploads/module"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
		return
	}

	thumbnailName := uuid.New().String() + thumbnailImage.Filename
	thumbnailPath := filepath.Join(uploadDir, thumbnailName)
	if err := c.SaveUploadedFile(thumbnailImage, thumbnailPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thumbnail file server"})
		return
	}

	// chapters
	
}