package models

type ChapterVideo struct {
	ChapterVideoFilename string `json:"chapter_video_filename"`
	ChapterVideoFilePath string `json:"chapter_video_file_path"`
	ChapterVideoSize     int64  `json:"chapter_video_size"`
}

type ChapterThumbnail struct {
	ChapterFilename string `json:"chapter_thumpnail_filename"`
	ChapterFilePath string `json:"chapter_thumpnail_file_path"`
}

type Chapter struct {
	ChapterId        string            `json:"chapter_id" gorm:"primaryKey"`
	LearningId       string            `json:"learning_id"`
	ChapterNo        int               `json:"chapter_no"`
	ChapterName      string            `json:"chapter_name"`
	ChapterDuration  float32           `json:"chapter_duration"`
	ChapterThumbnail *ChapterThumbnail `json:"chapter_thumpnail,omitempty" gorm:"embedded;embeddedPrefix=chapter_thumbnail_"`
	ChapterVideo     ChapterVideo      `json:"chapter_video" gorm:"embedded;embeddedPrefix=chapter_video_"`
	TutorName        string            `json:"tutor_name"`
}