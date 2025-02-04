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

// create study materials
func CreateStudyMaterial(c *gin.Context) {

	// module id
	materialId := uuid.New().String()

	// learnig id
	learningId := c.PostForm("learning_id")
	if learningId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Module id is required"})
		return
	}

	// material no
	materialNo := c.PostForm("material_no")
	materialNoParsed, err := strconv.Atoi(materialNo)
	if err != nil || materialNoParsed <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Material number must be a valid number"})
		return
	}

	// material name
	materialName := c.PostForm("material_name")

	// material pdf
	var materialPdfFile *models.MaterialPdf
	materialPdf, err := c.FormFile("material_pdf")
	if err == nil {
		materialUploadDir := "uploads/material"
		if err := os.MkdirAll(materialUploadDir, os.ModePerm); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create material upload directory"})
			return
		}

		materialFileName := uuid.New().String() + materialPdf.Filename
		materialFilePath := filepath.Join(materialUploadDir, materialFileName)
		if err := c.SaveUploadedFile(materialPdf, materialFilePath); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save metadata to server"})
			return
		}

		materialPdfFile = &models.MaterialPdf{
			MaterialFilename: materialFileName,
			MaterialFilePath: materialFilePath,
		}
	} else {
		materialPdfFile = &models.MaterialPdf{
			MaterialFilename: "",
			MaterialFilePath: "",
		}
	}

	// material model
	materialModel := models.StudyMaterial{
		MaterialId:   materialId,
		LearningId:   learningId,
		MaterialNo:   int(materialNoParsed),
		MaterialName: materialName,
		MaterialPdf:  materialPdfFile,
	}

	// save
	if err := config.DB.Create(&materialModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{"message": "Study material created", "data": materialModel})
}

// update study material
func UpdateStudyMaterial(c *gin.Context) {
	id := c.Param("id")
	var studyMaterial models.StudyMaterial

    // check the material
    if err := config.DB.Where("id = ?", id).First(&studyMaterial).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
        return
    }

	// material id
	materialId := c.PostForm("material_id")
	if materialId != "" {
		studyMaterial.MaterialId = materialId
	}

	// learning id
	learningId := c.PostForm("learning_id")
	if learningId != "" {
		studyMaterial.LearningId = learningId
	}

	// material no
	materialNo := c.PostForm("material_no")
	if materialNo != "" {
		materialNoParsed, err := strconv.Atoi(materialNo)
		if err != nil || materialNoParsed <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Material number must be a valid number"})
			return
		}
		studyMaterial.MaterialNo = materialNoParsed
	}

	// material name
	materialName := c.PostForm("material_name")
	if materialName != "" {
		studyMaterial.MaterialName = materialName
	}

	// material pdf
    var materialPdfFile *models.MaterialPdf
	materialPdf, err := c.FormFile("material_pdf")
	if err == nil {
		
        // remove old material
		if err := os.Remove(studyMaterial.MaterialPdf.MaterialFilePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old material"})
            return
		}

        materialUploadDir := "uploads/material"
        materialFileName := uuid.New().String() + materialPdf.Filename
        materialFilePath := filepath.Join(materialUploadDir, materialFileName)

        if err := c.SaveUploadedFile(materialPdf, materialFilePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save meta data to the server"})
            return
        }

        materialPdfFile = &models.MaterialPdf{
            MaterialFilename: materialFileName,
            MaterialFilePath: materialFilePath,
        }

        studyMaterial.MaterialPdf = materialPdfFile
	}

    // update
    if err := config.DB.Save(&studyMaterial).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // success
    c.JSON(http.StatusOK, gin.H{"message": "Study material updated", "data": studyMaterial})
}
