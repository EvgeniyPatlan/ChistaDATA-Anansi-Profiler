package stucts

import (
	"time"
)

type ExtractedLog struct {
	Timestamp time.Time
	ThreadId  int
	QueryId   string
	LogLevel  string
	Message   string
}
