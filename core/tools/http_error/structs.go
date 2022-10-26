package httperror

type ErrorResp struct {
	Error        map[string]interface{} `json:"error"`
	ErrorMessage string                 `json:"error_message"`
	ErrorCode    string                 `json:"error_code"`
	TraceID      string                 `json:"trace_id"`
}
