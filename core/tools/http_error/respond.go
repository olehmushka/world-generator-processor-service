package httperror

import (
	"context"
	"encoding/json"
	"net/http"

	traceIDTools "world_generator_processor_service/core/tools/trace_id"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

func SendErrorResp(ctx context.Context, w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	wErr := wrapped_error.Cast(err)
	resp := ErrorResp{
		ErrorMessage: wErr.Msg,
		TraceID:      traceIDTools.GetTraceID(ctx),
	}
	if wErr.ExtendedErr != nil {
		resp.Error = wErr.ExtendedErr.ErrorMap()
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(traceIDTools.TraceIDHeader, traceIDTools.GetTraceID(ctx))
	w.WriteHeader(wErr.StatusCode)
	_, _ = w.Write(respJSON)
}
