package main

import (
	"encoding/json"
	"fmt"
	"go_web_template/cli/cra"
	"log"
)

var (
	exampleOptions = cra.Options{
		ProjectParentDir:    "/Users/darthvader/Projects",
		TemplateDir:         ".",
		ProjectName:         "test_project",
		GoProjectModuleName: "example.com/me/test_project",
		Vendoring:           true,
		Selections: cra.Selections{
			API:  true,
			SSR:  true,
			Auth: cra.Session,
		},
	}
)

func main() {
	bytes, _ := json.MarshalIndent(exampleOptions, "", "    ")
	fmt.Println("Running with example project configuration")
	fmt.Println(string(bytes))

	err := cra.NewProject(exampleOptions)
	if err != nil {
		log.Fatal(err)
	}
}
