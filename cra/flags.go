package main

import (
	"flag"
)

type UserOptions struct {
	ProjectDir   string
	ProjectName  string
	GoModuleName string
	Vendoring    bool
	Verbose      bool
	Debug        bool
}

func buildFlags() UserOptions {

	userOptions := UserOptions{}

	flag.StringVar(&userOptions.ProjectDir, "project-dir", "", "Where to create the new project folder (full path)")
	flag.StringVar(&userOptions.ProjectName, "name", "", "Name of project")
	flag.StringVar(&userOptions.GoModuleName, "gomod", "", "Go module name")
	flag.BoolVar(&userOptions.Vendoring, "V", false, "Enable vendoring")
	flag.BoolVar(&userOptions.Verbose, "vvv", false, "Verbose output")
	flag.BoolVar(&userOptions.Debug, "d", false, "Debug mode usually used for development")

	flag.Parse()
	// if flag.NFlag() == 0 {
	// 	flag.PrintDefaults()
	// }

	return userOptions
}
