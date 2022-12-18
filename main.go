package main

import (
	"github.com/leslie-qiwa/react-admin-demo/config"
	"github.com/leslie-qiwa/react-admin-demo/infra/database"
	"github.com/leslie-qiwa/react-admin-demo/infra/logger"
	"github.com/leslie-qiwa/react-admin-demo/routers"
	"github.com/spf13/viper"
	"time"
)

func main() {
	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "America/Los_Angeles")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	dbName := config.DbConfiguration()

	if err := database.DBConnection(dbName); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	router := routers.Routes()

	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
