package traceid

import (
	"net/http"

	randomTools "github.com/olehmushka/golang-toolkit/random_tools"
)

func SetTraceIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := randomTools.UUIDString()
		if v := r.Header.Get(TraceIDHeader); v != "" {
			traceID = v
		}
		ctx = SetTraceID(ctx, traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
