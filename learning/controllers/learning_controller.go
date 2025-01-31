package controllers

import (
	"learning/config"
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
	
    // id
	learningId := uuid.New().String()

	// module no
	moduleNo := c.PostForm("module_no")
	moduleNoParsed, err := strconv.Atoi(moduleNo)
	if err != nil || moduleNoParsed <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Module number must be a valid number"})
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
	totalDurationParsed, err := strconv.ParseFloat(totalDuration, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid total duration"})
		return
	}

	// category
	category := c.PostForm("category")
	if !isValidCategory(category) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	// thumbnail image
	var thumbnailFile models.Thumbnail
	thumbnailImage, err := c.FormFile("thumbnail")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Thumbnail is required"})
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

	thumbnailFile = models.Thumbnail{
		Filename: thumbnailName,
		FilePath: thumbnailPath,
	}

	// set learning model
	learning := models.Learning{
		Id:             learningId,
		ModuleNo:       int(moduleNoParsed),
		ModuleName:     moduleName,
		TotalDuration:  float32(totalDurationParsed),
		Category:       category,
		Thumbnail:      thumbnailFile,
	}

	if err := config.DB.Create(&learning).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{"message": "Learning created", "data": learning})
}
