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
	var learning models.Learning

	// id
	learningId := uuid.New().String()

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

	// ==========================================>>>>>>>>>>>>>>>> CHAPTER
	var chapters []models.Chapter

	chapterId := uuid.New().String()
	// chapter no
	chapterNo := c.PostForm("chapter_no")
	if chapterNo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chapter number is required"})
		return
	}

	chapterNoParsed, err := strconv.ParseInt(moduleNo, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chapter number must be a valid number"})
		return
	}

	if chapterNoParsed <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chapter number must be greater than zero"})
		return
	}

	// chapter name
	chapterName := c.PostForm("chapter_name")
	if chapterName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chapter name is required"})
		return
	}

	// chapter duration
	chapterDuration := c.PostForm("chapter_duration")
	if chapterDuration == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chapter duration is required"})
		return
	}

	chapterDurationParsed, err := strconv.ParseFloat(chapterDuration, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duration"})
		return
	}

	// chapter thumbnail
	var chapterThumbnailFile *models.ChapterThumbnail
	chapterImage, err := c.FormFile("chapter_thumpnail")
	if err == nil {
		chapterImageUploadDir := "uploads/chapter/images"
		if err := os.MkdirAll(chapterImageUploadDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		chapterFileName := uuid.New().String() + chapterImage.Filename
		chapterFilePath := filepath.Join(chapterImageUploadDir, chapterFileName)
		if err := c.SaveUploadedFile(chapterImage, chapterFilePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save metadata to server"})
			return
		}

		chapterThumbnailFile = &models.ChapterThumbnail{
			ChapterFilename: chapterFileName,
			ChapterFilePath: chapterFilePath,
		}
	} else {
		chapterThumbnailFile = &models.ChapterThumbnail{
			ChapterFilename: "",
			ChapterFilePath: "",
		}
	}

	// chapter video
	var chapterVideoFile models.ChapterVideo
	chapterVideo, err := c.FormFile("chapter_video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chapter video not found"})
		return
	}

	// validate file type
	ext := filepath.Ext(chapterVideo.Filename)
	if !isValidExtensions(ext) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chater video extension"})
		return
	}

	chapterVideoUploadDir := "uploads/chapter/videos"
	if err := os.MkdirAll(chapterVideoUploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create chapter video upload directory"})
		return
	}

	// save file to the server
	chapterVideoFileName := uuid.New().String() + chapterVideo.Filename
	chapterVideoFilePath := filepath.Join(chapterVideoUploadDir, chapterVideoFileName)

	if err := c.SaveUploadedFile(chapterVideo, chapterVideoFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save chapter video file"})
		return
	}

	chapterVideoFile = models.ChapterVideo{
		ChapterVideoFilename: chapterVideoFileName,
		ChapterVideoFilePath: chapterVideoFilePath,
		ChapterVideoSize:     chapterVideo.Size,
	}

	// tutor name
	tutorName := c.PostForm("tutor_name")

	// chapter model
	chapterModel := models.Chapter{
		ChapterId:        chapterId,
		LearningId:       learningId,
		ChapterNo:        int(chapterNoParsed),
		ChapterName:      chapterName,
		ChapterDuration:  float32(chapterDurationParsed),
		ChapterThumbnail: chapterThumbnailFile,
		ChapterVideo:     chapterVideoFile,
		TutorName:        tutorName,
	}

	chapters = append(chapters, chapterModel)

	// ==========================================>>>>>>>>>>>>>>>> STUDY MATERIAL
	var studyMaterials []models.StudyMaterial

	// module id
	materialId := uuid.New().String()

	// material no
	materialNo := c.PostForm("material_no")
	materialNoParsed, err := strconv.ParseInt(materialNo, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Material number must be a valid number"})
		return
	}

	if materialNoParsed <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Material number must be greater than zero"})
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

	studyMaterials = append(studyMaterials, materialModel)

	// ==========================================>>>>>>>>>>>>>>>> SAVE
	// set learning model
	learning = models.Learning{
		Id:             learningId,
		ModuleNo:       int(moduleNoParsed),
		ModuleName:     moduleName,
		TotalDuration:  float32(totalDurationParsed),
		Category:       category,
		Thumbnail:      thumbnailFile,
		Chapters:       &chapters,
		StudyMaterials: &studyMaterials,
	}

	if err := config.DB.Create(&learning).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{"message": "Learning created", "data": learning})
}
