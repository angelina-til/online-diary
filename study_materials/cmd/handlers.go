package main

import (
	"context"
	"net/http"
	"time"

	openAPI "github.com/angelina-til/online-diary/shared/pkg/openapi/study_materials/v1"
)

type UserProfileHandler struct {
	storage *UserProfileStorage
}

func GetUserProfileHandler(st *UserProfileStorage) *UserProfileHandler {
	return &UserProfileHandler{
		storage: st,
	}
}

func (h *UserProfileHandler) GetWeatherByCity(_ context.Context, _ openAPI.GetWeatherByCityParams) (openAPI.GetWeatherByCityRes, error) {
	return &openAPI.Weather{
		City:        "",
		Temperature: 0,
		UpdatedAt:   time.Time{},
	}, nil
}

func (h *UserProfileHandler) UpdateWeatherByCity(
	_ context.Context,
	_ *openAPI.UpdateWeatherRequest, _ openAPI.UpdateWeatherByCityParams,
) (openAPI.UpdateWeatherByCityRes, error) {
	return &openAPI.Weather{
		City:        "",
		Temperature: 0,
		UpdatedAt:   time.Time{},
	}, nil
}

// NewError создает новую ошибку в формате GenericError
func (h *UserProfileHandler) NewError(_ context.Context, err error) *openAPI.GenericErrorStatusCode {
	return &openAPI.GenericErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: openAPI.GenericError{
			Code:    openAPI.NewOptInt(http.StatusInternalServerError),
			Message: openAPI.NewOptString(err.Error()),
		},
	}
}
