package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHealthCheck(t *testing.T) {
	h := &handlers{}

	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/health-check", nil)
	rec := httptest.NewRecorder()

	h.GetHealthCheck(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var resp GetHealthCheckResponse
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, "ok", resp.Status)
}

func TestGetInfo(t *testing.T) {
	h := &handlers{}

	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/info", nil)
	rec := httptest.NewRecorder()

	h.GetInfo(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var resp GetInfoResponse
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
}
