package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asvins/common_db/postgres"
	"github.com/asvins/utils/config"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
)

var (
	ServerConfig *Config        = new(Config)
	rend         *render.Render = render.New()
	db           *gorm.DB
)

// function that will run before main
func init() {
	fmt.Println("[INFO] Initializing server")
	err := config.Load("core_config.gcfg", ServerConfig)
	if err != nil {
		log.Fatal(err)
	}

	DatabaseConfig := postgres.NewConfig(ServerConfig.Database.User, ServerConfig.Database.DbName, ServerConfig.Database.SSLMode)
	db = postgres.GetDatabase(DatabaseConfig)
	fmt.Println("[INFO] Initialization Done!")
}

func main() {
	r := DefRoutes()

	fmt.Println("[INFO] Server running on port:", ServerConfig.Server.Port)
	http.ListenAndServe(":"+ServerConfig.Server.Port, r)
}
