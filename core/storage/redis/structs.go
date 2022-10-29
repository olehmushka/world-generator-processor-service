package redis

type Message struct {
	Data      []byte
	TraceID   string
	Timestamp string
}
