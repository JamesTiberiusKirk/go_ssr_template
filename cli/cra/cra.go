package cra

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func NewProject(options Options) error {
	dest := fmt.Sprintf("%s/%s", options.ProjectParentDir, options.ProjectName)
	log.Printf("Making new project directory %s", dest)
	err := os.Mkdir(dest, 0755)
	if err != nil {
		return err
	}

	mainConfig := ""
	dest += "/"

	escaped := strings.Replace(options.GoProjectModuleName, ".", "\\.", -1)
	escaped = strings.Replace(escaped, "/", "\\/", -1)
	sedCommand := fmt.Sprintf(sedTemplateNameCommand, escaped)

	log.Println("Moving base files")
	for baseFile, t := range baseFiles {
		filePath := fmt.Sprintf("%s/%s", options.TemplateDir, baseFile)
		err := cp(filePath, dest+baseFile, sedCommand, t)
		if err != nil {
			return err
		}
	}

	if options.Selections.SSR {
		log.Println("Moving Site files")
		for ssrFile, t := range ssrFiles {
			filePath := fmt.Sprintf("%s/%s", options.TemplateDir, ssrFile)
			err := cp(filePath, dest+ssrFile, sedCommand, t)
			if err != nil {
				return err
			}
		}

		mainConfig += fmt.Sprintf(configs[siteConfig])
	}

	if options.Selections.API {
		log.Println("Moving API files")
		for apiFile, t := range apiFiles {
			filePath := fmt.Sprintf("%s/%s", options.TemplateDir, apiFile)
			err := cp(filePath, dest+apiFile, sedCommand, t)
			if err != nil {
				return err
			}
		}

		mainConfig += fmt.Sprintf(configs[apiConfig])
	}

	log.Println("Creating main from template")
	err = makeMainFileFromTemplate(options.TemplateDir, dest, mainConfig, options.GoProjectModuleName)
	if err != nil {
		return nil
	}

	if options.Vendoring {
		installCmds = append(installCmds, CMD{
			Key:     "vendoring",
			Command: []string{"go", "mod", "vendor"},
		})

	}

	log.Println("Running go setup commands")
	err = runGoSetupCommands(options.ProjectName, options.GoProjectModuleName, dest)
	if err != nil {
		return nil
	}

	return nil
}

func makeMainFileFromTemplate(templatePath, projectPath, mainConig, moduleName string) error {
	mainConfig := struct {
		Code       string
		ModuleName string
	}{
		Code:       mainConig,
		ModuleName: moduleName,
	}

	mainTPath := fmt.Sprintf("%s/%s", templatePath, mainGoTmpl)
	mainDPath := fmt.Sprintf("%s/%s", projectPath, mainGo)

	t, err := template.ParseFiles(mainTPath)
	if err != nil {
		return err
	}

	file, err := os.Create(mainDPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = t.Execute(file, mainConfig)
	if err != nil {
		return err
	}
	return nil
}
