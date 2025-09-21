package main

import (
	"errors"
	"sync"

	openAPI "github.com/angelina-til/online-diary/shared/pkg/openapi/study_materials/v1"
)

type StudyPlanStorage struct {
	mu    sync.RWMutex
	plans map[int64]*openAPI.PlanObj
}

func NewStudyPlanStorage() *StudyPlanStorage {
	return &StudyPlanStorage{
		plans: make(map[int64]*openAPI.PlanObj),
	}
}

func (s *StudyPlanStorage) GetPlan(planID int64) (*openAPI.PlanObj, bool) {
	plan, ok := s.plans[planID]

	return plan, ok
}

func (s *StudyPlanStorage) InsertPlan(p *openAPI.PlanObj) error {
	if p == nil {
		return errors.New("пустой план")
	}

	if !p.ID.IsSet() {
		return errors.New("пустой ID")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.plans[p.ID.Value] = p

	return nil
}
