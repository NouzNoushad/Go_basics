package models

import "time"

type StudyMaterial struct {
	MaterialNo   int    `json:"material_no"`
	MaterialName string `json:"material_name"`
	MaterialPdf  string `json:"material_pdf"`
}

type ChapterThumbnail struct {
	ChapterFilename string `json:"chapter_thumpnail_filename"`
	ChapterFilePath string `json:"chapter_thumpnail_file_path"`
}

type Chapter struct {
	ChapterNo        int               `json:"chapter_no"`
	ChapterName      string            `json:"chapter_name"`
	ChapterDuration  float32           `json:"chapter_duration"`
	ChapterThumbnail *ChapterThumbnail `json:"chapter_thumpnail,omitempty" gorm:"embedded;embeddedPrefix=chapter_thumbnail_"`
	ChapterVideo     string            `json:"chapter_video"`
	TutorName        string            `json:"tutor_name"`
}

type ThumbnailUrl struct {
	Filename string `json:"thumbnail_filename"`
	FilePath string `json:"thumbnail_file_path"`
}

type Learning struct {
	Id             string           `json:"id" gorm:"primaryKey"`
	ModuleNo       int              `json:"module_no"`
	ModuleName     string           `json:"module_name"`
	TotalDuration  float32          `json:"total_duration"`
	Category       string           `json:"category"`
	ThumbnailUrl   ThumbnailUrl     `json:"thumbnail_url" gorm:"embedded;embeddedPrefix=thumbnail_"`
	Chapters       []Chapter        `json:"chapters" gorm:"embedded;embeddedPrefix=chapter_"`
	StudyMaterials *[]StudyMaterial `json:"study_materials,omitempty" gorm:"embedded;embeddedPrefix=material_"`
	CreatedAt      time.Time        `json:"created_at"`
}
