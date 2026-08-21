package traceid

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAndGetTraceID(t *testing.T) {
	tCases := map[string]struct {
		traceID string
	}{
		"should round-trip a non-empty trace id": {
			traceID: "trace-abc-123",
		},
		"should round-trip an empty trace id": {
			traceID: "",
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			ctx := SetTraceID(context.Background(), tc.traceID)
			assert.Equal(tt, tc.traceID, GetTraceID(ctx))
		})
	}
}

func TestGetTraceID_NotSet(t *testing.T) {
	assert.Empty(t, GetTraceID(context.Background()))
}

func TestSetTraceIDMiddleware(t *testing.T) {
	tCases := map[string]struct {
		headerValue     string
		expectGenerated bool
	}{
		"should propagate the trace id supplied via the request header": {
			headerValue: "incoming-trace-id",
		},
		"should generate a trace id when the request has none": {
			headerValue:     "",
			expectGenerated: true,
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			var gotTraceID string
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotTraceID = GetTraceID(r.Context())
			})

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tc.headerValue != "" {
				req.Header.Set(TraceIDHeader, tc.headerValue)
			}
			rec := httptest.NewRecorder()

			SetTraceIDMiddleware(next).ServeHTTP(rec, req)

			if tc.expectGenerated {
				assert.NotEmpty(tt, gotTraceID)
			} else {
				assert.Equal(tt, tc.headerValue, gotTraceID)
			}
		})
	}
}
