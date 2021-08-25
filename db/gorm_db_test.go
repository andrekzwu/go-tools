package db

import (
	"fmt"
	"testing"
)

// TestDB
func TestGORMDB(t *testing.T) {
	// 1. regiser db
	RegisterGORMDB(&GORMDBEntry{
		InstanceName: "db_r",
		Dialect:      "mysql",
		Host:         "xxxxxxxxxxxxx",
		Port:         3306,
		User:         "xxxxxxxxx",
		Password:     "xxxxxxxxxx",
		Database:     "xxxxxxxxxxx",
		MaxIdle:      2,
		MaxOpen:      10,
		LogMode:      true,
		Logger:       new(SQLLogger),
	})
	// query
	var total uint32
	sql := `SELECT COUNT(1) FROM xxxxxxxxxxxxxx`
	if err := GORMDB("db_r").Raw(sql).Count(&total).Error; err != nil {
		t.Errorf("err %v", err)
		return
	}
	fmt.Println(total)
}
