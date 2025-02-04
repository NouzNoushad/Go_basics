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

// update learning
func UpdateLearning(c *gin.Context) {
	id := c.Param("id")
	var learning models.Learning

	if err := config.DB.Where("id = ?", id).First(&learning).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
		return
	}

	// module no
	moduleNo := c.PostForm("module_no")
	if moduleNo != "" {
		moduleNoParsed, err := strconv.Atoi(moduleNo)
		if err != nil || moduleNoParsed <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Module no must be valid number"})
			return
		}
		learning.ModuleNo = moduleNoParsed
 	}

	// module name
	moduleName := c.PostForm("module_name")
	if moduleName != "" {
		learning.ModuleName = moduleName
	}

	// total duration
	totalDuration := c.PostForm("total_duration")
	if totalDuration != "" {
		totalDurationParsed, err := strconv.ParseFloat(totalDuration, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid total duration"})
			return
		}
		learning.TotalDuration = float32(totalDurationParsed)
	}

	// category
	category := c.PostForm("category")
	if category != "" {
		// check category
		if !isValidCategory(category) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
			return
		}
		learning.Category = category
	}

	// thumbnail image
	var thumbnailFile models.Thumbnail
	thumbnailImage, err := c.FormFile("thumbnail")
	if err == nil {
		// delete old thumbnail
		if err := os.Remove(learning.Thumbnail.FilePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old thumbnail"})
			return
		}

		uploadDir := "uploads/module"
		thumbnailName := uuid.New().String() + thumbnailImage.Filename
		thumbnailPath := filepath.Join(uploadDir, thumbnailName)

		if err := c.SaveUploadedFile(thumbnailImage, thumbnailPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Faile to create thumbnail file server"})
			return
		}

		thumbnailFile = models.Thumbnail{
			Filename: thumbnailName,
			FilePath: thumbnailPath,
		}

		learning.Thumbnail = thumbnailFile
	}

	// update module
	if err := config.DB.Save(&learning).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{"message": "Module updated", "data": learning})
}
