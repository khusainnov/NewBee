package api

import (
	"context"
	"net/http"

	"github.com/khusainnov/newbee/app/errors"
	"github.com/khusainnov/newbee/app/specs"
)

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

func (a *API) GetEducationMaterial(
	r *http.Request,
	req *specs.GetEducationMaterialReq,
	resp *specs.GetEducationMaterialResp,
) error {
	data, err := a.emProcessor.GetEducationMaterial(context.Background(), a.db, req.ID)
	if err != nil {
		r.Response.Status = "-32005"
		return err
	}

	resp.Body = data
	return nil
}

func (a *API) DeleteEducationMaterial(r *http.Request, req *specs.DeleteEducationMaterialReq, resp *specs.Resp) error {
	if err := a.emProcessor.DeleteEducationMaterial(context.Background(), a.db, req.ID); err != nil {
		resp.Message = err.Error()
		r.Response.Status = "-32005"
		return nil
	}

	resp.Message = "Success"
	return nil
}
