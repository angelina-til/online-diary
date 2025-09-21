package main

import (
	"sync"

	userProfile "github.com/angelina-til/online-diary/shared/pkg/openapi/user_profile/v1"
)

type UserProfileStorage struct {
	mu       sync.RWMutex
	profiles map[int64]*userProfile.Weather
}

func NewUserProfileStorage() *UserProfileStorage {
	return &UserProfileStorage{
		profiles: make(map[int64]*userProfile.Weather),
	}
}
