package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/newbee/app/models"
)

type TaskRepo struct{}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{}
}

func (r *TaskRepo) SaveTask(ctx context.Context, db *sqlx.DB, task *models.Task) error {
	return nil
}

func (r *TaskRepo) UpdateTask(ctx context.Context, db *sqlx.DB, task *models.Task) error {
	return nil
}

func (r *TaskRepo) GetTask(ctx context.Context, db *sqlx.DB, id string) (*models.Task, error) {
	return nil, nil
}

func (r *TaskRepo) GetTaskWithMaterials(ctx context.Context, db *sqlx.DB, id string) (*models.TaskWithMaterials, error) {
	return nil, nil
}

func (r *TaskRepo) ListTask(ctx context.Context, db *sqlx.DB) ([]*models.Task, error) {
	return nil, nil
}

func (r *TaskRepo) DeleteTask(ctx context.Context, db *sqlx.DB, id string) error {
	return nil
}
