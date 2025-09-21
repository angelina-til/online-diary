package main

import (
	"context"
	"net/http"
	"time"

	userProfile "github.com/angelina-til/online-diary/shared/pkg/openapi/user_profile/v1"
)

type UserProfileHandler struct {
	storage *UserProfileStorage
}

func GetUserProfileHandler(st *UserProfileStorage) *UserProfileHandler {
	return &UserProfileHandler{
		storage: st,
	}
}

func (h *UserProfileHandler) GetWeatherByCity(_ context.Context, _ userProfile.GetWeatherByCityParams) (userProfile.GetWeatherByCityRes, error) {
	return &userProfile.Weather{
		City:        "",
		Temperature: 0,
		UpdatedAt:   time.Time{},
	}, nil
}

func (h *UserProfileHandler) UpdateWeatherByCity(
	_ context.Context,
	_ *userProfile.UpdateWeatherRequest, _ userProfile.UpdateWeatherByCityParams,
) (userProfile.UpdateWeatherByCityRes, error) {
	return &userProfile.Weather{
		City:        "",
		Temperature: 0,
		UpdatedAt:   time.Time{},
	}, nil
}

// NewError создает новую ошибку в формате GenericError
func (h *UserProfileHandler) NewError(_ context.Context, err error) *userProfile.GenericErrorStatusCode {
	return &userProfile.GenericErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: userProfile.GenericError{
			Code:    userProfile.NewOptInt(http.StatusInternalServerError),
			Message: userProfile.NewOptString(err.Error()),
		},
	}
}
