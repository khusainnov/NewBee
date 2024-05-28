package api

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/newbee/app/processors/education_material"
	"github.com/khusainnov/newbee/app/specs"
)

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
