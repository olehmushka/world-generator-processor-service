package http

import (
	"encoding/json"
	httpError "world_generator_processor_service/core/tools/http_error"
	traceIDTools "world_generator_processor_service/core/tools/trace_id"

	"net/http"

	wrappedError "github.com/olehmushka/golang-toolkit/wrapped_error"
)

// GetInfo godoc
// @Summary Get info about app end-point
// @Description Get info about app end-point
// @Tags System
// @Accept json
// @Produce json
// @Param X-Trace-ID header string false "Custom trace id"
// @Success 200 {object} GetInfoResponse
// @Failure 500 {object} httperror.ErrorResp
// @Router /info [get]
func (h *handlers) GetInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// deps := h.pingService.PingServices(ctx)
	resp := GetInfoResponse{Dependencies: DependenciesResponse{
		// MongoDB: deps.MongoDB,
	}}
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
