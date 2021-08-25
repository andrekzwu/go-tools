package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
)

type GORMDBEntry struct {
	InstanceName string // Instance name, sql map key
	Dialect      string
	User         string
	Password     string
	Database     string
	Host         string
	Port         uint32
	MaxIdle      int
	MaxOpen      int
	LogMode      bool
	Logger       gorm.LogWriter
	Params       Param
}

type Param string

const (
	DEFAULT_UTF8_PARAM    Param = "charset=utf8&parseTime=True&loc=Local"
	DEFAULT_UTF8MB4_PARAM Param = "charset=utf8mb4&parseTime=True&loc=Local"
)

func (param Param) handleParam() Param {
	if param == "" {
		return DEFAULT_UTF8_PARAM
	}
	return param
}

func (entry *GORMDBEntry) DNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", entry.User, entry.Password, entry.Host, entry.Port, entry.Database, entry.Params.handleParam())
}

var (
	gormDBMap = new(sync.Map)
)

// RegisterDB
func RegisterGORMDB(entry *GORMDBEntry) {
	// check db entry
	if entry == nil || entry.InstanceName == "" {
		panic("register db err,db entry nil or instance name is empty")
	}
	db, err := gorm.Open(entry.Dialect, entry.DNS())
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(entry.LogMode)
	if entry.Logger != nil {
		db.SetLogger(gorm.Logger{LogWriter: entry.Logger})
	}
	db.DB().SetMaxIdleConns(entry.MaxIdle)
	db.DB().SetMaxOpenConns(entry.MaxOpen)
	if err := db.DB().Ping(); err != nil {
		panic("failed to connect database:" + err.Error())
	}
	gormDBMap.Store(entry.InstanceName, db)
}

// DB
func GORMDB(instanceName string) *gorm.DB {
	item, ok := gormDBMap.Load(instanceName)
	if !ok {
		panic("db select fail,please register db instance:" + instanceName)
	}
	return item.(*gorm.DB)
}
