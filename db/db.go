package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "sync"
)

type DBEntry struct {
    InstanceName string // Instance name, sql map key
    Dialect      string
    User         string
    Password     string
    Database     string
    Host         string
    Port         string
    MaxIdle      int
    MaxOpen      int
}

var (
    dbMap = new(sync.Map)
)

// RegisterDB
func RegisterDB(dbEntry *DBEntry) {
    // check db entry
    if dbEntry == nil || dbEntry.InstanceName == "" {
        panic("register db err,db entry nil or instance name is empty")
    }
    // url
    url := dbEntry.User + ":" + dbEntry.Password + "@tcp(" + dbEntry.Host + ":" + dbEntry.Port + ")/" + dbEntry.Database + "?charset=utf8&parseTime=True&loc=Local"
    db, err := sql.Open(dbEntry.Dialect, url)
    if err != nil {
        panic("failed to connect database")
    }
    db.SetMaxIdleConns(dbEntry.MaxIdle)
    db.SetMaxOpenConns(dbEntry.MaxOpen)
    if err := db.Ping(); err != nil {
        panic("failed to connect database:" + err.Error())
    }
    dbMap.Store(dbEntry.InstanceName, db)
}

// DB
func DB(instanceName string) *sql.DB {
    db, ok := dbMap.Load(instanceName)
    if !ok {
        panic("db select fail,please register db instance:" + instanceName)
    }
    return db.(*sql.DB)
}
