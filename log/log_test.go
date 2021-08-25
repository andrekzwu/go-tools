package log

import (
	"testing"
	"time"
)

// TestLogPrint
func TestLogPrint(t *testing.T) {
	// new logger
	RegisterLogger(&LoggerEntry{
		IsDevelopment: false,
		SystemName:    "go_tools",
		FilePath:      "./xxxxxxx.log",
		NameSpace:     "xxxxxxxx",
	})
	ti := time.Now()
	var err error
	// print log
	PrintLog("test log print", ti, &err, "in", "out")
}
