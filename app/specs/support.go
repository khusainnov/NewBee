package specs

import "github.com/khusainnov/newbee/app/models"

type CreateCourseReq struct {
	Courses []*models.EducationMaterial `json:"courses"`
}
