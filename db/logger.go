package db

import (
	"fmt"

	"github.com/andrezhz/go-tools/log"
	"github.com/andrezhz/go-tools/util"
	"github.com/jinzhu/gorm"
)

type SQLLogger struct {
	gorm.LogWriter
}

func (logger *SQLLogger) Println(v ...interface{}) {
	log.LOGSQL("sql(%s).point(%s).use(%s)", util.CompressStr(fmt.Sprintf("%v", v[3]), " "), v[1], v[2])
}
