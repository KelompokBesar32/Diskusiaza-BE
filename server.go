package main

import (
	"Diskusiaza-BE/config"
	"Diskusiaza-BE/constants"
	"Diskusiaza-BE/database"
	"Diskusiaza-BE/routes"
)

func main() {
	// init configuration and running app
	configs := config.InitConfiguration()
	database.InitDB(configs)
	e := routes.New()
	e.Static(constants.StaticFileUsersFoto, constants.DirFileUsersFoto)
	e.Static(constants.StaticFileTherad, constants.DirFileTherad)
	e.Logger.Fatal(e.Start(configs.ServerAddress))
}
