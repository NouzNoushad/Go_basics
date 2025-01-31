package models

import "time"

type Thumbnail struct {
	Filename string `json:"thumbnail_filename"`
	FilePath string `json:"thumbnail_file_path"`
}

type Learning struct {
	Id             string           `json:"id" gorm:"primaryKey"`
	ModuleNo       int              `json:"module_no"`
	ModuleName     string           `json:"module_name"`
	TotalDuration  float32          `json:"total_duration"`
	Category       string           `json:"category"`
	Thumbnail      Thumbnail        `json:"thumbnail_url" gorm:"embedded;embeddedPrefix=thumbnail_"`
	Chapters       []Chapter        `json:"chapters" gorm:"foreignKey:LearningId"`
	StudyMaterials *[]StudyMaterial `json:"study_materials,omitempty" gorm:"foreignKey:LearningId"`
	CreatedAt      time.Time        `json:"created_at" gorm:"autoCreateTime"`
}
