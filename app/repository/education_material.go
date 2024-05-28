package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/newbee/app/errors"
	"github.com/khusainnov/newbee/app/models"
	"github.com/khusainnov/newbee/app/repository/builder"
)

const (
	educationMaterialsNumFields = 5
)

type EducationMaterialRepo struct{}

func NewEducationMaterialRepo() *EducationMaterialRepo {
	return &EducationMaterialRepo{}
}

func (r *EducationMaterialRepo) SaveEducationMaterial(
	ctx context.Context,
	db *sqlx.DB,
	data []*models.EducationMaterial,
) error {
	const baseQuery = `
		INSERT INTO education_material (
						name,
		                description,
		                category,
		                material_url,
		                task_id
		) VALUES`

	numEducationMeterials := len(data)
	query := fmt.Sprintf(
		"%s %s;",
		baseQuery,
		builder.BuildMultipleItemInsertQuery(numEducationMeterials, educationMaterialsNumFields),
	)

	args := make([]any, 0, numEducationMeterials*educationMaterialsNumFields)

	for _, em := range data {
		args = append(
			args,
			em.Name,
			em.Description,
			em.Category,
			em.MaterialLink,
			em.TaskID,
		)
	}

	_, err := db.ExecContext(ctx, query, args...)

	return errors.WrapIfErr(err, "failed to save education material")
}
