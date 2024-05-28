package education_material

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/newbee/app/models"
	"go.uber.org/zap"
)

type EducationMaterialProcessor interface {
	InsertEducationMaterial(ctx context.Context, db *sqlx.DB, courses []*models.EducationMaterial) error
}

type EducationMaterialRepo interface {
	SaveEducationMaterial(ctx context.Context, db *sqlx.DB, data []*models.EducationMaterial) error
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
