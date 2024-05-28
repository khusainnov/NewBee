package task

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/newbee/app/models"
	"go.uber.org/zap"
)

type TaskProcessor interface {
	InsertTask(ctx context.Context, db *sqlx.DB, task *models.Task) error
	UpdateTask(ctx context.Context, db *sqlx.DB, task *models.Task) error
	GetTask(ctx context.Context, db *sqlx.DB, id string) (*models.Task, error)
	GetTaskWithMaterials(ctx context.Context, db *sqlx.DB, id string) (*models.TaskWithMaterials, error)
	ListTask(ctx context.Context, db *sqlx.DB) ([]*models.Task, error)
	DeleteTask(ctx context.Context, db *sqlx.DB, id string) error
}

type TaskRepo interface {
	SaveTask(ctx context.Context, db *sqlx.DB, task *models.Task) error
	UpdateTask(ctx context.Context, db *sqlx.DB, task *models.Task) error
	GetTask(ctx context.Context, db *sqlx.DB, id string) (*models.Task, error)
	GetTaskWithMaterials(ctx context.Context, db *sqlx.DB, id string) (*models.TaskWithMaterials, error)
	ListTask(ctx context.Context, db *sqlx.DB) ([]*models.Task, error)
	DeleteTask(ctx context.Context, db *sqlx.DB, id string) error
}

type Processor struct {
	log      *zap.Logger
	taskRepo TaskRepo
}

func New(log *zap.Logger, taskRepo TaskRepo) *Processor {
	return &Processor{
		log:      log,
		taskRepo: taskRepo,
	}
}

func (p *Processor) InsertTask(ctx context.Context, db *sqlx.DB, task *models.Task) error {
	return nil
}

func (p *Processor) UpdateTask(ctx context.Context, db *sqlx.DB, task *models.Task) error {
	return nil
}

func (p *Processor) GetTask(ctx context.Context, db *sqlx.DB, id string) (*models.Task, error) {
	return nil, nil
}

func (p *Processor) GetTaskWithMaterials(ctx context.Context, db *sqlx.DB, id string) (*models.TaskWithMaterials, error) {
	return nil, nil
}

func (p *Processor) ListTask(ctx context.Context, db *sqlx.DB) ([]*models.Task, error) {
	return nil, nil
}

func (p *Processor) DeleteTask(ctx context.Context, db *sqlx.DB, id string) error {
	return nil
}
