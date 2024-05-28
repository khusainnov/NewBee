package models

type TaskStatus string

var (
	TaskStatusTODO       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in progress"
	TaskStatusReview     TaskStatus = "review"
	TaskStatusFinished   TaskStatus = "finished"
)

type Task struct {
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	UserID      string     `json:"user_id,omitempty"`
	DueTo       string     `json:"due_to"`
	CreatedAt   string     `json:"created_at,omitempty"`
	UpdatedAt   string     `json:"updated_at,omitempty"`
}

type TaskWithMaterials struct {
	Task               *Task
	EducationMaterials []*EducationMaterial
}
