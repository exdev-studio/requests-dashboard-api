package memstore_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/exdev-studio/requests-dashboard-api/internal/model"
	"github.com/exdev-studio/requests-dashboard-api/internal/store"
	"github.com/exdev-studio/requests-dashboard-api/internal/store/memstore"
)

func TestRequestRepository_Collect(t *testing.T) {
	s := memstore.New()
	r := model.TestRequest()

	assert.NoError(t, s.Request().Collect(r))
	assert.NotNil(t, r)
}

func TestRequestRepository_List(t *testing.T) {
	s := memstore.New()
	reqs, err := s.Request().List()

	assert.NoError(t, err)
	assert.Equal(t, []*model.Request{}, reqs)

	r := model.TestRequest()
	err = s.Request().Collect(r)
	assert.NoError(t, err)

	reqs, err = s.Request().List()
	assert.NoError(t, err)
	assert.Equal(t, []*model.Request{r}, reqs)
}

func TestRequestRepository_Get(t *testing.T) {
	id := 1
	s := memstore.New()
	req, err := s.Request().Get(id)

	assert.EqualError(t, err, errors.New(store.ErrNotFound).Error())
	assert.Nil(t, req)

	tr := model.TestRequest()
	err = s.Request().Collect(tr)
	assert.NoError(t, err)

	r, err := s.Request().Get(id)
	assert.NoError(t, err)
	assert.Equal(t, r, tr)
}
