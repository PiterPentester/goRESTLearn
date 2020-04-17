package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PiterPentester/goRESTLearn/internal/app/store/test_store"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(test_store.New())

	s.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
