package traceid

import (
	"context"

	contextTools "github.com/olehmushka/golang-toolkit/context_tools"
)

func GetTraceID(ctx context.Context) string {
	return contextTools.GetValue(ctx, TraceIDKey)
}
