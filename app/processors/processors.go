package processors

import (
	"github.com/khusainnov/newbee/app/processors/education_material"
	"go.uber.org/zap"
)

type Processors struct {
	CourseProcessor education_material.EducationMaterialProcessor
}

func New(log *zap.Logger) *Processors {
	return &Processors{
		CourseProcessor: education_material.New(log),
	}
}
