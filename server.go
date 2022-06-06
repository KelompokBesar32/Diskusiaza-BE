package main

import (
	"Diskusiaza-BE/config"
	"Diskusiaza-BE/database"
	"Diskusiaza-BE/routes"
)

func main() {
	// init configuration and running app
	configs := config.InitConfiguration()
	database.InitDB(configs)
	e := routes.New()
	e.Logger.Fatal(e.Start(configs.ServerAddress))
}
