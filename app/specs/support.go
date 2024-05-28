package specs

import "github.com/khusainnov/newbee/app/models"

type AddEducationMaterialReq struct {
	EducationMaterials []*models.EducationMaterial `json:"education_materials"`
}

type UpdateEducationMaterialReq struct {
	EducationMaterial *models.EducationMaterial `json:"education_material"`
}
