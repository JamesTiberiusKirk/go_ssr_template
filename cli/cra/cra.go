package cra

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func (wt *WebTemplate) NewProject(options Options) error {
	dest := fmt.Sprintf("%s/%s", options.ProjectParentDir, options.ProjectName)
	log.Printf("Making new project directory %s", dest)
	err := os.Mkdir(dest, 0755)
	if err != nil {
		return err
	}

	// mainConfig := ""
	dest += "/"

	escaped := strings.Replace(options.GoProjectModuleName, ".", "\\.", -1)
	escaped = strings.Replace(escaped, "/", "\\/", -1)
	sedCommand := fmt.Sprintf(wt.sedTemplateNameCommand, escaped)

	log.Println("Moving base files")
	for baseFile, t := range wt.baseFiles {
		filePath := fmt.Sprintf("%s/%s", options.TemplateDir, baseFile)
		err := customCopy(filePath, dest+baseFile, sedCommand, t)
		if err != nil {
			return err
		}
	}

	log.Println("Moving Site files")
	for ssrFile, t := range wt.siteFiles {
		filePath := fmt.Sprintf("%s/%s", options.TemplateDir, ssrFile)
		err := customCopy(filePath, dest+ssrFile, sedCommand, t)
		if err != nil {
			return err
		}
	}

	log.Println("Moving API files")
	for apiFile, t := range wt.apiFiles {
		filePath := fmt.Sprintf("%s/%s", options.TemplateDir, apiFile)
		err := customCopy(filePath, dest+apiFile, sedCommand, t)
		if err != nil {
			return err
		}
	}

	err = os.Chmod(dest+"dev_run.sh", 0755)
	if err != nil {
		return err
	}

	// log.Println("Creating main from template")
	// err = wt.makeMainFileFromTemplate(mainConfig)
	// if err != nil {
	// 	return err
	// }

	if options.Vendoring {
		wt.commands = append(wt.commands, CMD{
			Key:     "vendoring",
			Command: []string{"go", "mod", "vendor"},
		})
	}

	if options.Selections.SiteConfig.SPA {
		wt.commands = append(wt.commands, CMD{
			Key:     "spa:install",
			Command: []string{"npm", "run", "spa:install"},
		})
	}

	log.Println("Running go setup commands")
	err = wt.runGoSetupCommands(dest)
	if err != nil {
		return err
	}

	return nil
}

func (wt *WebTemplate) makeMainFileFromTemplate(mainConfigCode string) error {
	mainConfig := struct {
		Code       string
		ModuleName string
	}{
		Code:       mainConfigCode,
		ModuleName: wt.options.GoProjectModuleName,
	}

	mainTPath := fmt.Sprintf("%s/%s", wt.options.TemplateDir, wt.mainGoTemplate)
	mainDPath := fmt.Sprintf("%s/%s/%s", wt.options.ProjectParentDir, wt.options.ProjectName, wt.mainGo)

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

func (wt *WebTemplate) runGoSetupCommands(projectPath string) error {
	for _, cmd := range wt.commands {
		var err error
		switch cmd.Key {
		case "init":
			err = runCMD(projectPath, append(cmd.Command, wt.options.GoProjectModuleName), wt.verbose)
		case "imports":
			err = runCMD(projectPath, append(cmd.Command, wt.mainGo), wt.verbose)
		default:
			err = runCMD(projectPath, cmd.Command, wt.verbose)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
