package models

type TaskStatus string

var (
	TaskStatusTODO       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in progress"
	TaskStatusReview     TaskStatus = "review"
	TaskStatusFinished   TaskStatus = "finished"
)

type Task struct {
	ID          string     `json:"id,omitempty" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	Status      TaskStatus `json:"status" db:"status"`
	UserID      string     `json:"user_id,omitempty" db:"user_id"`
	DueTo       string     `json:"due_to" db:"due_to"`
	CreatedAt   string     `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   string     `json:"updated_at,omitempty" db:"updated_at"`
}

type TaskWithMaterials struct {
	Task               *Task
	EducationMaterials []*EducationMaterial
}
