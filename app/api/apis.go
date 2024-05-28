package api

import (
	"context"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/newbee/app/errors"
	"github.com/khusainnov/newbee/app/processors/education_material"
	"github.com/khusainnov/newbee/app/specs"
)

type Handlers interface {
	Echo(r *http.Request, req *specs.EchoReq, resp *specs.Resp) error
}

type Support interface {
	EducationMaterial
}

type EducationMaterial interface {
	SaveEducationMaterial(r *http.Request, req *specs.AddEducationMaterialReq, resp *specs.Resp) error
	UpdateEducationMaterial(r *http.Request, req *specs.UpdateEducationMaterialReq, resp *specs.Resp) error
}

type Task interface {
}

type API struct {
	db          *sqlx.DB
	emProcessor education_material.EducationMaterialProcessor
}

func NewAPI(db *sqlx.DB, emProcessor education_material.EducationMaterialProcessor) *API {
	return &API{
		db:          db,
		emProcessor: emProcessor,
	}
}

func (a *API) Echo(_ *http.Request, req *specs.EchoReq, resp *specs.Resp) error {
	log.Println(req.Message)
	resp.Message = req.Message
	return nil
}

func (a *API) SaveEducationMaterial(r *http.Request, req *specs.AddEducationMaterialReq, resp *specs.Resp) error {
	if len(req.EducationMaterials) == 0 {
		resp.Message = errors.EmptyRequestBody
		r.Response.Status = "-32001"
		return nil
	}

	return a.emProcessor.InsertEducationMaterial(context.Background(), a.db, req.EducationMaterials)
}

func (a *API) UpdateEducationMaterial(r *http.Request, req *specs.UpdateEducationMaterialReq, resp *specs.Resp) error {
	if req.EducationMaterial.ID == "" {
		resp.Message = errors.EmptyEducationMaterialID
		r.Response.Status = "-32003"
		return nil
	}
	
	return a.emProcessor.UpdateEducationMaterial(context.Background(), a.db, req.EducationMaterial)
}
