package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/newbee/app/errors"
	"github.com/khusainnov/newbee/app/models"
)

type TaskRepo struct{}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{}
}

func (r *TaskRepo) SaveTask(ctx context.Context, db *sqlx.DB, task *models.Task) error {
	query := `INSERT INTO task (name, description, status, user_id, due_to) VALUES ($1, $2, $3, $4, $5);`

	_, err := db.ExecContext(ctx, query, task.Name, task.Description, task.Status, task.UserID, task.DueTo)

	return errors.WrapIfErr(err, "failed to insert task")
}

func (r *TaskRepo) UpdateTask(ctx context.Context, db *sqlx.DB, task *models.Task) error {
	query := `UPDATE task 
				SET name=$2, description=$3, status=$4, updated_at=now()
					WHERE id=$1;`

	_, err := db.ExecContext(ctx, query, task.ID, task.Name, task.Description, task.Status)

	return errors.WrapIfErr(err, "failed to update task with id:"+task.ID)
}

func (r *TaskRepo) GetTask(ctx context.Context, db *sqlx.DB, id string) (*models.Task, error) {
	query := `SELECT id, name, description, status, user_id, due_to, created_at, updated_at 
				FROM task
					WHERE id=$1;`

	var task models.Task
	if err := db.GetContext(ctx, &task, query, id); err != nil {
		return nil, errors.WrapIfErr(err, "failed to get task")
	}

	return &task, nil
}

func (r *TaskRepo) GetTaskWithMaterials(ctx context.Context, db *sqlx.DB, id string) (*models.TaskWithMaterials, error) {
	task, err := r.GetTask(ctx, db, id)
	if err != nil {
		return nil, errors.WrapIfErr(err, "GetTaskWithMaterials")
	}

	materials, err := r.getMaterials(ctx, db, id)
	if err != nil {
		return nil, errors.WrapIfErr(err, "GetTaskWithMaterials")
	}

	return &models.TaskWithMaterials{
		Task:               task,
		EducationMaterials: materials,
	}, nil
}

func (r *TaskRepo) ListTask(ctx context.Context, db *sqlx.DB) ([]*models.Task, error) {
	query := `SELECT id, name, description, status, user_id, due_to, created_at, updated_at 
				FROM task;`

	var tasks []*models.Task
	if err := db.SelectContext(ctx, &tasks, query); err != nil {
		return nil, errors.WrapIfErr(err, "failed to get list of tasks")
	}

	return tasks, nil
}

func (r *TaskRepo) DeleteTask(ctx context.Context, db *sqlx.DB, id string) error {
	query := `DELETE FROM task WHERE id=$1;`

	_, err := db.ExecContext(ctx, query, id)

	return errors.WrapIfErr(err, "failed to delete task with id:"+id)
}

func (r *TaskRepo) getMaterials(ctx context.Context, db *sqlx.DB, taskID string) ([]*models.EducationMaterial, error) {
	query := `SELECT id, name, description, category, material_url, task_id, created_at, updated_at
				FROM education_material
					WHERE task_id = $1;`

	var materials []*models.EducationMaterial
	if err := db.SelectContext(ctx, &materials, query, taskID); err != nil {
		return nil, errors.WrapIfErr(err, "failed to get materials for task with id"+taskID)
	}

	return materials, nil
}
