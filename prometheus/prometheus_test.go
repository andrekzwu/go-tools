package prometheus

import (
	"testing"
)

func TestPrometheusRecord(t *testing.T) {
	// init
	config := Config{
		NameSpace: "xxxxxxx",
		System:    "xxxxxxxx",
	}
	New(config)

	// count request
	CountRequest("api/xxx/xxxxx/v1")

	// record
	Record("api/xxxxx/xxxx/v1", "0", 300)

}
