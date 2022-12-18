package config

import (
	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	Driver  string
	Dbname  string
	LogMode bool
}

func DbConfiguration() string {
	name := "demo.sqlite"
	viper.SetDefault("DB_NAME", name)
	return name
}
