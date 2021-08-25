package main

import (
	"fmt"

	"github.com/andrezhz/go-tools/acm"
)

func main() {
	// init acm
	SetupACM()
	// init db & redis
	Setup()
}

func SetupACM() {
	acm.RegisterACM(&acm.ACMEntry{
		Endpoint:    "",
		AccessKey:   "",
		SecretKey:   "",
		NamespaceId: "", // default namespace id
	})
}

func Setup() {
	// use default namespace id
	dbContent := acm.GetString(acm.NewACMParam("xxxxx.db", "db"))
	fmt.Println("db => ", dbContent)
	// setup db ...
	redisContent := acm.GetString(acm.NewACMParam("xxxxxxx.redis", "xxxxx").SetNamespaceId(""))
	fmt.Println("redis => ", redisContent)
	// setup redis
}
