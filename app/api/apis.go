package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/newbee/app/processors/education_material"
	"github.com/khusainnov/newbee/app/specs"
)

var (
	EmptyRequestBody = "request has empty data"
)

type Handlers interface {
	Echo(r *http.Request, req *specs.EchoReq, resp *specs.Resp) error
}

type Support interface {
	SaveEducationMaterial(r *http.Request, req *specs.CreateCourseReq, resp *specs.Resp) error
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

func (a *API) SaveEducationMaterial(r *http.Request, req *specs.CreateCourseReq, resp *specs.Resp) error {
	if len(req.Courses) == 0 {
		resp.Message = EmptyRequestBody
		r.Response.Status = "-32001"
		return nil
	}

	fmt.Printf("\n%v\n", *req.Courses[0])
	ctx := context.Background()

	return a.emProcessor.InsertEducationMaterial(ctx, a.db, req.Courses)
}
