package http

import (
	"encoding/json"
	"net/http"
	httpError "world_generator_processor_service/core/tools/http_error"
	traceIDTools "world_generator_processor_service/core/tools/trace_id"

	wrappedError "github.com/olehmushka/golang-toolkit/wrapped_error"
)

// GetHealthCheck godoc
// @Summary Health check end-point
// @Description Health check end-point
// @Tags System
// @Accept json
// @Produce json
// @Param x-trace-id header string false "Custom trace id"
// @Success 200 {object} GetHealthCheckResponse
// @Failure 500 {object} httperror.ErrorResp
// @Router /health-check [get]
func (h *handlers) GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp := GetHealthCheckResponse{Status: "ok"}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		httpError.SendErrorResp(ctx, w, wrappedError.NewInternalServerError(err, "can not marshal response"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(traceIDTools.TraceIDHeader, traceIDTools.GetTraceID(ctx))
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(respJSON)
}
