package api

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/newbee/app/processors/education_material"
	"github.com/khusainnov/newbee/app/processors/task"
	"github.com/khusainnov/newbee/app/repository"
	"github.com/khusainnov/newbee/app/specs"
	"go.uber.org/zap"
)

type API struct {
	db            *sqlx.DB
	log           *zap.Logger
	emProcessor   education_material.EducationMaterialProcessor
	taskProcessor task.TaskProcessor
}

func NewAPI(
	log *zap.Logger,
	db *sqlx.DB,
) *API {
	return &API{
		db:            db,
		log:           log,
		emProcessor:   education_material.New(log, repository.NewEducationMaterialRepo()),
		taskProcessor: task.New(log, repository.NewTaskRepo()),
	}
}

func (a *API) Echo(_ *http.Request, req *specs.EchoReq, resp *specs.Resp) error {
	log.Println(req.Message)
	resp.Message = req.Message
	return nil
}
