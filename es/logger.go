package es

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/andrezhz/go-tools/log"
)

type ConsoleLogger struct {
}

func (*ConsoleLogger) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

type DefaultElasticTraceLogger struct {
}

func (*DefaultElasticTraceLogger) Printf(format string, v ...interface{}) {
	strArray := strings.Split(v[0].(string), "\r\n")
	t := time.Now()
	var err error
	if len(strArray) > 4 {
		// request
		first := strArray[0]
		array := strings.Split(strings.TrimSpace(first), " ")
		var buff bytes.Buffer
		buff.WriteString(strings.Join(array[1:3], " ") + " ")
		buff.WriteString(strings.Join(array[len(array)-3:len(array)-1], " ") + " ")
		log.PrintLog("es.request.log", t, &err, buff.String(), strArray[9])
	} else {
		// response
		log.PrintLog("es.response.log", t, &err, "", strArray[3])
	}
}

type DefaultElasticErrorLogger struct {
}

func (*DefaultElasticErrorLogger) Printf(format string, v ...interface{}) {
	err := fmt.Errorf(format, v...)
	log.PrintLog("es.error.log", time.Now(), &err, "", "")
}
