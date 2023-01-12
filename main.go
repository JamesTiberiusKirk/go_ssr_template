package main

import (
	"encoding/json"
	"fmt"
	"go_web_template/api"
	"go_web_template/session"
	"go_web_template/site"
)

func main() {
	initLogger()

	config := buildConfig()

	db := initDb(config)
	e := initServer(config)

	sessionManager := session.New()

	s := site.NewSite(config.Http.RootPath, db, sessionManager, e)
	s.Serve()

	a := api.NewApi(e.Group(config.Http.RootApiPath), config.Http.RootApiPath, db,
		sessionManager)
	a.Serve()

	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	fmt.Print(string(data))

	e.Logger.Fatal(e.Start(config.Http.Port))
}
