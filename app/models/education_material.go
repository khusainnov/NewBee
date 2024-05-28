package models

type EducationMaterial struct {
	ID           string `json:"id,omitempty" db:"id"`
	TaskID       string `json:"task_id" db:"task_id"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	Category     string `json:"category" db:"category"`
	MaterialLink string `json:"material_link" db:"material_url"`
	CreatedAt    string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    string `json:"updated_at,omitempty" db:"updated_at"`
}
