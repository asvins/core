package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asvins/common_db/postgres"
	"github.com/asvins/common_io"
	"github.com/asvins/utils/config"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
)

// var DatabaseConfig *postgres.Config
var (
	ServerConfig *Config        = new(Config)
	rend         *render.Render = render.New()
	db           *gorm.DB
	producer     *common_io.Producer
	consumer     *common_io.Consumer
)

func init() {
	fmt.Println("[INFO] Initializing server")

	/*
	*	Server config
	 */
	err := config.Load("core_config.gcfg", ServerConfig)
	if err != nil {
		log.Fatal(err)
	}

	/*
	*	Database
	 */
	DatabaseConfig := postgres.NewConfig(ServerConfig.Database.User, ServerConfig.Database.DbName, ServerConfig.Database.SSLMode)
	db = postgres.GetDatabase(DatabaseConfig)
	fmt.Println("[INFO] Initialization Done!")

	/*
	*	Common io
	 */
	//	setupCommonIo()
}

func main() {
	r := DefRoutes()

	fmt.Println("[INFO] Server running on port:", ServerConfig.Server.Port)
	http.ListenAndServe(":"+ServerConfig.Server.Port, r)
}
