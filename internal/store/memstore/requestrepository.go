package memstore

import (
	"errors"

	"github.com/exdev-studio/requests-dashboard-api/internal/model"
	"github.com/exdev-studio/requests-dashboard-api/internal/store"
)

type RequestRepository struct {
	store    *Store
	requests map[int]*model.Request
}

func (r *RequestRepository) Get(id int) (*model.Request, error) {
	req, ok := r.requests[id]
	if !ok {
		return nil, errors.New(store.ErrNotFound)
	}

	return req, nil
}

func (r *RequestRepository) List() ([]*model.Request, error) {
	requests := make([]*model.Request, 0)

	for _, req := range r.requests {
		requests = append(requests, req)
	}

	return requests, nil
}
