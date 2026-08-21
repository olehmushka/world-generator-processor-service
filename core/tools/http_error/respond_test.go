package httperror

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	traceIDTools "world_generator_processor_service/core/tools/trace_id"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendErrorResp(t *testing.T) {
	tCases := map[string]struct {
		ctx                context.Context
		err                error
		expectedStatusCode int
		expectedErrMessage string
		expectHasError     bool
	}{
		"should not write anything for a nil error": {
			ctx:                context.Background(),
			err:                nil,
			expectedStatusCode: 0, // no WriteHeader call
		},
		"should respond with the wrapped error status code and message": {
			ctx:                context.Background(),
			err:                wrapped_error.NewNotFoundError(errors.New("not found"), "resource not found"),
			expectedStatusCode: http.StatusNotFound,
			expectedErrMessage: "resource not found",
			expectHasError:     true,
		},
		"should default to 500 for a plain, non-wrapped error": {
			ctx:                context.Background(),
			err:                errors.New("boom"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedErrMessage: "",
		},
		"should include the trace id from the context in the response and headers": {
			ctx:                traceIDTools.SetTraceID(context.Background(), "trace-123"),
			err:                wrapped_error.NewBadRequestError(errors.New("bad"), "bad request"),
			expectedStatusCode: http.StatusBadRequest,
			expectedErrMessage: "bad request",
			expectHasError:     true,
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			rec := httptest.NewRecorder()

			SendErrorResp(tc.ctx, rec, tc.err)

			if tc.err == nil {
				assert.Equal(tt, 200, rec.Code) // httptest defaults Code to 200 when WriteHeader is never called
				assert.Empty(tt, rec.Body.String())
				return
			}

			assert.Equal(tt, tc.expectedStatusCode, rec.Code)
			assert.Equal(tt, "application/json", rec.Header().Get("Content-Type"))
			assert.Equal(tt, traceIDTools.GetTraceID(tc.ctx), rec.Header().Get(traceIDTools.TraceIDHeader))

			var resp ErrorResp
			require.NoError(tt, json.Unmarshal(rec.Body.Bytes(), &resp))
			assert.Equal(tt, tc.expectedErrMessage, resp.ErrorMessage)
			assert.Equal(tt, traceIDTools.GetTraceID(tc.ctx), resp.TraceID)
			if tc.expectHasError {
				assert.NotEmpty(tt, resp.Error)
			} else {
				assert.Empty(tt, resp.Error)
			}
		})
	}
}
