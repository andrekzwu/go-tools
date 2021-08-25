package config

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
	"time"
)

func TestRegisterConfig(t *testing.T) {
	RegisterConfig(&ConfigEntry{ConfigName: "local_test"})
	t.Log(viper.GetString("name"))
	t.Log(viper.GetUint32("age"))
	t.Log(viper.GetUint32("head.test"))
}

func TestRegisterConfigAndMergeACM(t *testing.T) {
	RegisterConfig(&ConfigEntry{ConfigName: "local_test",ConfigType:"json"})
	for {
		// local config
		fmt.Println(viper.GetString("name"))
		fmt.Println(viper.GetUint32("age"))
		fmt.Println(viper.GetUint32("head.test"))
		time.Sleep(time.Second * 5)
	}
}