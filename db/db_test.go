package db

import (
	"fmt"
	"testing"
)

// TestDB
func TestDB(t *testing.T) {
	// 1. regiser db
	RegisterDB(&DBEntry{
		InstanceName: "db_r",
		Dialect:      "mysql",
		Host:         "xxxxxxxxxxx",
		Port:         "3306",
		User:         "xxxxxxx",
		Password:     "xxxxxxx",
		Database:     "xxxxxxxxxx",
		MaxIdle:      2,
		MaxOpen:      10,
	})
	// query
	var total uint32
	sql := `SELECT COUNT(1) FROM xxxxxxxxxxxxxxx`
	if err := DB("db_r").QueryRow(sql).Scan(&total); err != nil {
		t.Errorf("err %v", err)
		return
	}
	fmt.Println(total)

}
