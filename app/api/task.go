package api

import (
	"context"
	"net/http"

	"github.com/khusainnov/newbee/app/specs"
	"go.uber.org/zap"
)

func (a *API) CreateTask(r *http.Request, req *specs.CreateTaskReq, resp *specs.Resp) error {
	return a.taskProcessor.InsertTask(context.Background(), a.db, req.Task)
}

func (a *API) UpdateTask(r *http.Request, req *specs.UpdateTaskReq, resp *specs.Resp) error {
	if err := a.taskProcessor.UpdateTask(context.Background(), a.db, req.Task); err != nil {
		return err
	}

	return nil
}

func (a *API) GetTask(r *http.Request, req *specs.GetTaskReq, resp *specs.GetTaskResp) error {
	respTask, err := a.taskProcessor.GetTask(context.Background(), a.db, req.ID)
	if err != nil {
		a.log.Error("failed to get task", zap.Error(err))

		return err
	}

	resp.Task = respTask
	return nil
}

func (a *API) GetTaskWithMaterials(
	r *http.Request,
	req *specs.GetTaskWithMaterialsReq,
	resp *specs.GetTaskWithMaterialsResp,
) error {
	respData, err := a.taskProcessor.GetTaskWithMaterials(context.Background(), a.db, req.ID)
	if err != nil {
		a.log.Error("failed to list task with materials", zap.Error(err))

		return err
	}

	resp.Body.Task = respData.Task
	resp.Body.EducationMaterials = respData.EducationMaterials

	return nil
}

func (a *API) ListTask(r *http.Request, req *specs.ListTaskReq, resp *specs.ListTaskResp) error {
	tasks, err := a.taskProcessor.ListTask(context.Background(), a.db)
	if err != nil {
		a.log.Error("failed to list tasks", zap.Error(err))

		return err
	}

	resp.Body = tasks

	return nil
}

func (a *API) DeleteTask(r *http.Request, req *specs.DeleteTaskReq, resp *specs.Resp) error {
	if err := a.taskProcessor.DeleteTask(context.Background(), a.db, req.ID); err != nil {
		a.log.Error("failed to delete task", zap.String("task_id", req.ID), zap.Error(err))

		return err
	}

	return nil
}
