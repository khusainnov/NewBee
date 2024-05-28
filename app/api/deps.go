package api

import (
	"net/http"

	"github.com/khusainnov/newbee/app/specs"
)

type Handlers interface {
	Echo(r *http.Request, req *specs.EchoReq, resp *specs.Resp) error
}

type Support interface {
	EducationMaterial
	Task
}

type EducationMaterial interface {
	SaveEducationMaterial(r *http.Request, req *specs.AddEducationMaterialReq, resp *specs.Resp) error
	UpdateEducationMaterial(r *http.Request, req *specs.UpdateEducationMaterialReq, resp *specs.Resp) error
	GetEducationMaterial(r *http.Request, req *specs.GetEducationMaterialReq, resp *specs.GetEducationMaterialResp) error
	DeleteEducationMaterial(r *http.Request, req *specs.DeleteEducationMaterialReq, resp *specs.Resp) error
}

type Task interface {
	CreateTask(r *http.Request, resp *specs.Resp) error
	UpdateTask(r *http.Request, resp *specs.Resp) error
	GetTask(r *http.Request, resp *specs.GetTaskResp) error
	GetTaskWithMaterials(r *http.Request, resp *specs.GetTaskWithMaterialsResp) error
	ListTask(r *http.Request, resp *specs.ListTaskResp) error
	DeleteTask(r *http.Request, resp *specs.Resp) error
}
