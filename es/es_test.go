package es

import (
	"testing"
	"time"
)

func TestRegisterES(t *testing.T) {
	err := RegisterES(&ESEntry{
		IsDevelopment: true,
		URLList:       []string{"http://127.0.0.1:9200"},
		Interval:      time.Second * 10,
		Gzip:          false,
		TraceLogger:   new(DefaultElasticTraceLogger),
		ErrorLogger:   new(DefaultElasticErrorLogger),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ESClient.String())
}
