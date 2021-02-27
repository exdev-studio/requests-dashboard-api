package store

import (
	"github.com/exdev-studio/requests-dashboard-api/internal/model"
)

type RequestRepository interface {
	Get(id int) (*model.Request, error)
	List() ([]*model.Request, error)
	Collect(*model.Request) error
}
