package models

type MaterialPdf struct {
	MaterialFilename string `json:"material_filename"`
	MaterialFilePath string `json:"material_file_path"`
}

type StudyMaterial struct {
	MaterialId   string       `json:"material_id" gorm:"primaryKey"`
	LearningId   string       `json:"learning_id"`
	MaterialNo   int          `json:"material_no"`
	MaterialName string       `json:"material_name"`
	MaterialPdf  *MaterialPdf `json:"material_pdf,omitempty" gorm:"embedded;embeddedPrefix=material_pdf_"`
}