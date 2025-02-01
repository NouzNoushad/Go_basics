package controllers

import (
	"learning/config"
	"learning/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLearningList(c *gin.Context) {
	var learnings []models.Learning

	if err := config.DB.Preload("Chapters").Preload("StudyMaterials").Find(&learnings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": learnings})
}

// get learning by id
func GetLearningById(c *gin.Context) {
	var id = c.Param("id")
	var learning models.Learning

	if err := config.DB.Preload("Chapters").Preload("StudyMaterials").Where("id = ?", id).First(&learning).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": learning})
}
