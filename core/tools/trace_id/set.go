package traceid

import (
	"context"

	contextTools "github.com/olehmushka/golang-toolkit/context_tools"
)

func SetTraceID(ctx context.Context, traceID string) context.Context {
	return contextTools.SetValue(ctx, TraceIDKey, traceID)
}
