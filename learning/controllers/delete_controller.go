package controllers

import (
	"learning/config"
	"learning/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// delete learning
func DeleteLearning(c *gin.Context) {
	var id = c.Param("id")
	var learning models.Learning
	var chapter models.Chapter
	var studyMaterial models.StudyMaterial

	if err := config.DB.Preload("Chapters").Preload("StudyMaterials").Where("id = ?", id).First(&learning).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	// remove material pdf
	if studyMaterial.MaterialPdf != nil {
		if !deleteServerFile(c, studyMaterial.MaterialPdf.MaterialFilePath, "Failed to delete material pdf from the server") {
			return
		}
	}

	// remove chapter thumbnail, chapter video
	if chapter.ChapterThumbnail != nil {
		if !deleteServerFile(c, chapter.ChapterThumbnail.ChapterFilePath, "Failed to delete chapter thumbnail from the server") {
			return
		}
	}

	if !deleteServerFile(c, chapter.ChapterVideo.ChapterVideoFilePath, "Failed to delete chapter video from the server") {
		return
	}

	// remove module thumbnail
	if !deleteServerFile(c, learning.Thumbnail.FilePath, "Failed to delete thumpnail from server") {
		return
	}

	// delete material
	if !deleteRecord(c, config.DB, "learning_id = ?", learning.Id, &studyMaterial, "Failed to delete material record") {
		return
	}

	// delete chapter
	if !deleteRecord(c, config.DB, "learning_id = ?", learning.Id, &chapter, "Failed to delete chapter") {
		return
	}

	// delete learning
	if err := config.DB.Delete(&learning).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

// utilities

func deleteServerFile(c *gin.Context, filePath string, errMessage string) bool {
	if filePath != "" {
		if err := os.Remove(filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMessage})
			return false
		}
	}
	return true
}

func deleteRecord(c *gin.Context, db *gorm.DB, condition string, value string, model interface{}, errMessage string) bool {
	if err := db.Where(condition, value).Delete(model).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMessage + ":" + err.Error()})
		return false
	}
	return true
}
