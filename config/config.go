package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

type ConfigEntry struct {
	ConfigName string
	ConfigType string
}

// RegisterConfig
func RegisterConfig(entry *ConfigEntry) {
	viper.SetConfigName(entry.ConfigName)
	viper.SetConfigType(entry.ConfigType)
	// Read from configmap first
	viper.AddConfigPath("/etc/config")
	// Read from current path
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read config fail:%v", err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Info("config.change.event", in.String())
	})
}