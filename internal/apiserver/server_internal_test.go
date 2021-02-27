package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_handleRequestsList(t *testing.T) {
	c := NewConfig()
	c.LogLevel = "info"
	s := newServer(c)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/requests", nil)
	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestServer_HandleRequestsCollect(t *testing.T) {
	c := NewConfig()
	c.LogLevel = "info"
	s := newServer(c)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"type": "VK_CREDENTIALS",
				"fields": []map[string]string{
					{
						"type":  "string",
						"name":  "login",
						"value": "value",
					},
				},
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid type",
			payload: map[string]interface{}{
				"type": "invalid",
				"fields": []map[string]string{
					{
						"type":  "string",
						"name":  "login",
						"value": "value",
					},
				},
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "empty type",
			payload: map[string]interface{}{
				"type": "",
				"fields": []map[string]string{
					{
						"type":  "string",
						"name":  "login",
						"value": "value",
					},
				},
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "type are not set",
			payload: map[string]interface{}{
				"fields": []map[string]string{
					{
						"type":  "string",
						"name":  "login",
						"value": "value",
					},
				},
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "empty fields",
			payload: map[string]interface{}{
				"type":   "VK_CREDENTIALS",
				"fields": []map[string]string{},
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "fields are not set",
			payload: map[string]interface{}{
				"type": "VK_CREDENTIALS",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid field type",
			payload: map[string]interface{}{
				"type": "VK_CREDENTIALS",
				"fields": []map[string]string{
					{
						"type":  "invalid",
						"name":  "login",
						"value": "value",
					},
				},
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "field type is not set",
			payload: map[string]interface{}{
				"type": "VK_CREDENTIALS",
				"fields": []map[string]string{
					{
						"name":  "login",
						"value": "value",
					},
				},
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "name not set",
			payload: map[string]interface{}{
				"type": "VK_CREDENTIALS",
				"fields": []map[string]string{
					{
						"type":  "string",
						"value": "value",
					},
				},
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "name is empty",
			payload: map[string]interface{}{
				"type": "VK_CREDENTIALS",
				"fields": []map[string]string{
					{
						"type":  "string",
						"name":  "",
						"value": "value",
					},
				},
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "name not unique",
			payload: map[string]interface{}{
				"type": "VK_CREDENTIALS",
				"fields": []map[string]string{
					{
						"type":  "string",
						"name":  "login",
						"value": "value",
					},
					{
						"type":  "string",
						"name":  "login",
						"value": "value",
					},
				},
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			if err := json.NewEncoder(b).Encode(tc.payload); err != nil {
				t.Fatal(err)
			}
			req, _ := http.NewRequest(http.MethodPost, "/requests/collect", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
