package education_material

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/newbee/app/models"
	"go.uber.org/zap"
)

type EducationMaterialProcessor interface {
	InsertEducationMaterial(ctx context.Context, db *sqlx.DB, educationMaterial []*models.EducationMaterial) error
	UpdateEducationMaterial(ctx context.Context, db *sqlx.DB, educationMaterial *models.EducationMaterial) error
	GetEducationMaterial(ctx context.Context, db *sqlx.DB, id string) (*models.EducationMaterial, error)
	DeleteEducationMaterial(ctx context.Context, db *sqlx.DB, id string) error
}

type EducationMaterialRepo interface {
	SaveEducationMaterial(ctx context.Context, db *sqlx.DB, educationMaterial []*models.EducationMaterial) error
	UpdateEducationMaterial(ctx context.Context, db *sqlx.DB, educationMaterial *models.EducationMaterial) error
	GetEducationMaterial(ctx context.Context, db *sqlx.DB, id string) (*models.EducationMaterial, error)
	DeleteEducationMaterial(ctx context.Context, db *sqlx.DB, id string) error
}

type Processor struct {
	log    *zap.Logger
	emRepo EducationMaterialRepo
}

func New(log *zap.Logger, emRepo EducationMaterialRepo) *Processor {
	return &Processor{
		log:    log,
		emRepo: emRepo,
	}
}

func (p *Processor) InsertEducationMaterial(ctx context.Context, db *sqlx.DB, data []*models.EducationMaterial) error {
	if err := p.emRepo.SaveEducationMaterial(ctx, db, data); err != nil {
		return err
	}

	return nil
}

func (p *Processor) UpdateEducationMaterial(
	ctx context.Context,
	db *sqlx.DB,
	educationMaterial *models.EducationMaterial,
) error {
	if err := p.emRepo.UpdateEducationMaterial(ctx, db, educationMaterial); err != nil {
		return err
	}

	return nil
}

func (p *Processor) GetEducationMaterial(
	ctx context.Context,
	db *sqlx.DB,
	id string,
) (*models.EducationMaterial, error) {
	if id == "" {
		return nil, fmt.Errorf("unable to get material due empty id")
	}

	resp, err := p.emRepo.GetEducationMaterial(ctx, db, id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *Processor) DeleteEducationMaterial(
	ctx context.Context,
	db *sqlx.DB,
	id string,
) error {
	if id == "" {
		return fmt.Errorf("unable to delete material due empty id")
	}

	if err := p.emRepo.DeleteEducationMaterial(ctx, db, id); err != nil {
		return err
	}

	return nil
}
