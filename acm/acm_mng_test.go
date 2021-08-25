package acm

import (
	"testing"

	"github.com/andrezhz/go-tools/db"
	"github.com/tidwall/gjson"
)

func TestRegisterDB1(t *testing.T) {
	// 1. set acm entry
	RegisterACM(&ACMEntry{
		Endpoint:    "xxxxxx",
		AccessKey:   "xxxxxxxx",
		SecretKey:   "xxxxxxxx",
		NamespaceId: "xxxxxxxxxxxx", // default namespace id
	})
	// 2. fix acm param
	acmParam := NewACMParam("xxxxxxxxxxxxx.db", "db")

	// 4. get acm content
	content := GetString(acmParam)

	// 5. setup
	SetupDB(content)

	// 5. query data
	var count int
	err := db.GORMDB("db_r").Raw("select count(1) from message_template").Count(&count).Error
	if err != nil {
		panic(err)
	}
	t.Log("count ===> ", count)
}

func TestRegisterDB2(t *testing.T) {
	// 1. new acm
	acm := New(&ACMEntry{
		Endpoint:  "xxxxxxxxx",
		AccessKey: "xxxxxxxxxx",
		SecretKey: "xxxxxxxxxxxx",
	})

	// 2. fix acm param & add setup db call func
	acmParam := NewACMParam("xxxxxxxxxxxxx.db", "db").SetNamespaceId("xxxxxxxxxxx")

	// 4. exec handlers
	content := acm.GetString(acmParam)
	SetupDB(content)
	// 5. query data
	var count int
	err := db.GORMDB("db_r").Raw("select count(1) from message_template").Count(&count).Error
	if err != nil {
		panic(err)
	}
	t.Log("count ===> ", count)
}

// setup db
func SetupDB(content string) {
	dbAcmConfig := gjson.Parse(content)
	// register gorm db
	db.RegisterGORMDB(&db.GORMDBEntry{
		InstanceName: "db_r",
		Dialect:      "mysql",
		Host:         dbAcmConfig.Get("xxxxx_r.host").String(),
		Port:         3306,
		User:         dbAcmConfig.Get("xxxxx_r.user").String(),
		Password:     dbAcmConfig.Get("xxxxx_r.password").String(),
		Database:     dbAcmConfig.Get("xxxxx_r.db").String(),
		MaxIdle:      2,
		MaxOpen:      10,
		LogMode:      true,
	})
}
