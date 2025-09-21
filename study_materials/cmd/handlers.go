package main

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	openAPI "github.com/angelina-til/online-diary/shared/pkg/openapi/study_materials/v1"
)

type StudyPlanHandler struct {
	storage *StudyPlanStorage
}

func GetStudyPlanHandler(st *StudyPlanStorage) *StudyPlanHandler {
	return &StudyPlanHandler{
		storage: st,
	}
}

func (h *StudyPlanHandler) GetPlanByID(_ context.Context, p openAPI.GetPlanByIDParams) (openAPI.GetPlanByIDRes, error) {
	plan, ok := h.storage.GetPlan(p.PlanID)
	if !ok {
		return &openAPI.NotFoundError{}, nil
	}

	return &openAPI.PlanObj{
		ID:        plan.ID,
		Title:     plan.Title,
		CreatedAt: plan.CreatedAt,
	}, nil
}

// ToDo:  Исправить роут чтобы не требовал ИД
func (h *StudyPlanHandler) CreatePlan(
	_ context.Context,
	r *openAPI.CreatePlanRequest, _ openAPI.CreatePlanParams,
) (openAPI.CreatePlanRes, error) {
	plan := openAPI.PlanObj{
		ID:        openAPI.NewOptInt64(rand.Int63n(100)),
		Title:     r.Title,
		CreatedAt: time.Now(),
		UpdatedAt: openAPI.NewOptDateTime(time.Now()),
	}

	err := h.storage.InsertPlan(&plan)
	if err != nil {
		return &openAPI.InternalServerError{}, nil
	}

	return &plan, nil
}

// NewError создает новую ошибку в формате GenericError
func (h *StudyPlanHandler) NewError(_ context.Context, err error) *openAPI.GenericErrorStatusCode {
	return &openAPI.GenericErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: openAPI.GenericError{
			Code:    openAPI.NewOptInt(http.StatusInternalServerError),
			Message: openAPI.NewOptString(err.Error()),
		},
	}
}
