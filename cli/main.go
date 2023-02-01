package main

import (
	"encoding/json"
	"fmt"
	"go_web_template/cli/cra"
	"log"
)

var (
	exampleOptions = cra.Options{
		ProjectParentDir: "/home/darthvader/Projects",
		// ProjectParentDir:    "/Users/darthvader/Projects",
		TemplateDir:         ".",
		ProjectName:         "test_project",
		GoProjectModuleName: "example.com/me/test_project",
		Vendoring:           false,
		Selections: cra.Selections{
			API:  true,
			Site: true,
			SiteConfig: &cra.Site{
				Templating: true,
				SSR:        true,
				Static:     true,
				SPA:        true,
			},
			Auth: cra.Session,
		},
	}
)

func main() {
	bytes, _ := json.MarshalIndent(exampleOptions, "", "    ")
	fmt.Println("Running with example project configuration")
	fmt.Println(string(bytes))

	tp := cra.NewWebTemplate(exampleOptions)
	err := tp.NewProject(exampleOptions)
	if err != nil {
		log.Fatal(err)
	}
}
