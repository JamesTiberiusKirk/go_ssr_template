package main

import (
	"encoding/json"
	"fmt"
	"go_web_template/api"
	"go_web_template/server"
	"go_web_template/session"
	"go_web_template/site"
)

func main() {
	initLogger()

	config := buildConfig()

	db := initDb(config)
	e := initServer(config)

	sessionManager := session.New()

	servers := []server.Server{
		api.NewApi(e.Group(config.Http.RootApiPath), config.Http.RootApiPath, db,
			sessionManager),
		site.NewSite(e, config.Http.RootSitePath, db, sessionManager),
	}

	for _, s := range servers {
		s.Serve()
	}

	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	fmt.Print(string(data))

	e.Logger.Fatal(e.Start(config.Http.Port))
}
