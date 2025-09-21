package main

import (
	"sync"

	openAPI "github.com/angelina-til/online-diary/shared/pkg/openapi/study_materials/v1"
)

type UserProfileStorage struct {
	mu       sync.RWMutex
	profiles map[int64]*openAPI.Weather
}

func NewUserProfileStorage() *UserProfileStorage {
	return &UserProfileStorage{
		profiles: make(map[int64]*openAPI.Weather),
	}
}
