package main

import (
	"github.com/JamesTiberiusKirk/go_web_template/api"
	"github.com/JamesTiberiusKirk/go_web_template/session"
	"github.com/JamesTiberiusKirk/go_web_template/site"
)

func main() {
	initLogger()

	config := buildConfig()

	db := initDB(config)
	e := initServer()

	sessionManager := session.New()

	apiServer := api.NewAPI(e.Group(config.HTTP.RootAPIPath), config.HTTP.RootAPIPath, db,
		sessionManager)
	siteServer := site.NewSite(e, config.HTTP.RootSitePath, db, sessionManager, config.Debug)

	apiServer.Serve()
	apiRoutes := apiServer.GetRoutes()

	siteServer.SetRoutes("api", apiRoutes)
	siteServer.Serve()

	e.Logger.Fatal(e.Start(config.HTTP.Port))
}
