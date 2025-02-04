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

// create chapter
func CreateChapter(c *gin.Context) {

	// chapter id
	chapterId := uuid.New().String()

	// learnig id
	learningId := c.PostForm("learning_id")

	// chapter no
	chapterNo := c.PostForm("chapter_no")
	chapterNoParsed, err := strconv.Atoi(chapterNo)
	if err != nil || chapterNoParsed <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chapter number must be a valid number"})
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
	chapterDurationParsed, err := strconv.ParseFloat(chapterDuration, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chpater duration"})
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

	if err := config.DB.Create(&chapterModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{"message": "Chapter created", "data": chapterModel})
}

// update chapter
func UpdateChapter(c *gin.Context) {
	id := c.Param("id")
	var chapter models.Chapter

	// check the chapter
	if err := config.DB.Where("id = ?", id).First(&chapter).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
		return
	}

	// learning id
	learningId := c.PostForm("learning_id")
	if learningId != "" {
		chapter.LearningId = learningId
	}

	// chapter no
	chapterNo := c.PostForm("chapter_no")
	if chapterNo != "" {
		chapterNoParsed, err := strconv.Atoi(chapterNo)
		if err != nil || chapterNoParsed <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Chapter number must be a valid number"})
			return
		}
		chapter.ChapterNo = chapterNoParsed
	}

	// chapter name
	chapterName := c.PostForm("chapter_name")
	if chapterName != "" {
		chapter.ChapterName = chapterName
	}

	// chapter duration
	chapterDuration := c.PostForm("chapter_duration")
	if chapterDuration != "" {
		chapterDurationParsed, err := strconv.ParseFloat(chapterDuration, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chpater duration"})
			return
		}
		chapter.ChapterDuration = float32(chapterDurationParsed)
	}

	// chapter thumbnail
	var chapterThumbnailFile *models.ChapterThumbnail
	chapterImage, err := c.FormFile("chapter_thumpnail")
	if err == nil {

		// delete old thumbnail
		if err := os.Remove(chapter.ChapterThumbnail.ChapterFilePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old thumbnail"})
			return
		}

		chapterImageUploadDir := "uploads/chapter/images"
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

		chapter.ChapterThumbnail = chapterThumbnailFile
	}

	// chapter video
	var chapterVideoFile models.ChapterVideo
	chapterVideo, err := c.FormFile("chapter_video")
	if err == nil {
		// validate file type
		ext := filepath.Ext(chapterVideo.Filename)
		if !isValidExtensions(ext) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chater video extension"})
			return
		}

		// remove old video
		if err := os.Remove(chapter.ChapterVideo.ChapterVideoFilePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove old chapter video upload directory"})
			return
		}

		// save file to the server
		chapterVideoUploadDir := "uploads/chapter/videos"
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

		chapter.ChapterVideo = chapterVideoFile
	}

	// tutor name
	tutorName := c.PostForm("tutor_name")
	if tutorName != "" {
		chapter.TutorName = tutorName
	}

	// update chapter
	if err := config.DB.Save(&chapter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{"message": "Chapter updated", "data": chapter})
}
