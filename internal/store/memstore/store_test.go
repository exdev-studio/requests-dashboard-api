package memstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/exdev-studio/requests-dashboard-api/internal/store/memstore"
)

func TestStore_Request(t *testing.T) {
	s := memstore.New()
	r := s.Request()

	assert.NotNil(t, r)
}
