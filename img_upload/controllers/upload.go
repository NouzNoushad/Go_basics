package controllers

import (
	"fmt"
	"img_upload/config"
	"img_upload/models"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadImage(c *gin.Context) {
	var upload models.Upload

	description := c.PostForm("description")
	if description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
		return
	}

	// get file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch file"})
		return
	}

	// save file to server
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	filePath := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save metadata to database"})
		return
	}

	// create upload model
	upload.Id = uuid.New().String()
	upload.Image = models.Image{
		Filename: file.Filename,
		FilePath: filePath,
	}
	upload.Description = description

	if err := config.DB.Create(&upload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File Uploaded"})
}

// get uploads
func GetUploads(c *gin.Context) {
	var uploads []models.Upload

	if err := config.DB.Order("created_at desc").Find(&uploads).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": fmt.Sprintf("%d items", len(uploads)), "data": uploads})
}

// get upload by id
func GetUploadById(c *gin.Context) {

	id := c.Param("id")
	var upload models.Upload

	if err := config.DB.Where("id = ?", id).First(&upload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": upload})
}

// delete upload
func DeleteUpload(c *gin.Context) {
    id := c.Param("id")
    var upload models.Upload

    if err := config.DB.Where("id = ?", id).First(&upload).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Upload not found"})
        return
    }

    if err := os.Remove(upload.Image.FilePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file from server"})
        return
    }

    if err := config.DB.Delete(&upload).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Upload deleted"})
}

// update upload
func UpdateUpload(c *gin.Context) {
    id := c.Param("id")
    var upload models.Upload

    if err := config.DB.Where("id = ?", id).First(&upload).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Upload not found"})
        return
    }

    descripton := c.PostForm("description")
    if descripton != "" {
        upload.Description = descripton
    }

    file, err := c.FormFile("image")
    if err == nil {
        // remove old file
        if err := os.Remove(upload.Image.FilePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old file"})
            return
        }

        // save new file
        uploadDir := "uploads"
        filePath := filepath.Join(uploadDir, file.Filename)
        if err := c.SaveUploadedFile(file, filePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save new file"})
            return
        }

        // update file
        upload.Image = models.Image{
            Filename: file.Filename,
            FilePath: filePath,
        }
    }

    if err := config.DB.Save(&upload).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Upload updated", "data": upload})
}