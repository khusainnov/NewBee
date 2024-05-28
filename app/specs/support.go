package specs

import "github.com/khusainnov/newbee/app/models"

// Request structs

type AddEducationMaterialReq struct {
	EducationMaterials []*models.EducationMaterial `json:"education_materials"`
}

type UpdateEducationMaterialReq struct {
	EducationMaterial *models.EducationMaterial `json:"education_material"`
}

type GetEducationMaterialReq struct {
	ID string `json:"id"`
}

type DeleteEducationMaterialReq struct {
	ID string `json:"id"`
}

type CreateTaskReq struct {
	Task *models.Task `json:"task"`
}

type UpdateTaskReq struct {
	Task *models.Task `json:"task"`
}

type GetTaskReq struct {
	ID string `json:"id"`
}

type GetTaskWithMaterialsReq struct {
	ID string `json:"id"`
}

type ListTaskReq struct{}

type DeleteTaskReq struct {
	ID string `json:"id"`
}

// Response structs

type GetEducationMaterialResp struct {
	Body *models.EducationMaterial `json:"body"`
}

type GetTaskResp struct {
	Task *models.Task `json:"body"`
}

type GetTaskWithMaterialsResp struct {
	Body struct {
		Task               *models.Task                `json:"task"`
		EducationMaterials []*models.EducationMaterial `json:"education_materials"`
	} `json:"body"`
}

type ListTaskResp struct {
	Body []*models.Task `json:"body"`
}
