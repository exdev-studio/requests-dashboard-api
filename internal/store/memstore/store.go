package memstore

import (
	"github.com/exdev-studio/requests-dashboard-api/internal/model"
	"github.com/exdev-studio/requests-dashboard-api/internal/store"
)

type Store struct {
	requestRepository *RequestRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Request() store.RequestRepository {
	if s.requestRepository != nil {
		return s.requestRepository
	}

	s.requestRepository = &RequestRepository{
		store:    s,
		requests: make(map[int]*model.Request),
	}

	return s.requestRepository
}
