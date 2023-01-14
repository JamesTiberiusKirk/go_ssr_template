package cra

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/rwtodd/Go.Sed/sed"
)

func customCopy(source, dest, sedCmd string, t ItemType) error {
	switch t {
	case file:
		err := customCopyFile(source, dest, sedCmd)
		if err != nil {
			return err
		}
	case templateFile:
		err := customCopyTemplateFile(source, dest, nil)
		if err != nil {
			return err
		}
	case dir:
		err := customCopyDir(source, dest, sedCmd)
		if err != nil {
			return err
		}
	default:
		log.Printf("Type: %s | Unsupported file type", t)
		return nil
	}

	return nil
}

func customCopyTemplateFile(source, dest string, template any) error {
	return nil
}

func customCopyFile(source, dest, sedCmd string) error {
	sourceFileStat, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", source)
	}

	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if sedCmd != "" {
		engine, err := sed.New(strings.NewReader(sedCmd))
		if err != nil {
			return err
		}

		// Can't for some reason properly typecast io.Reader to a os.File for just assigning
		//	engine.Wrap back to srcFile for code re-usability
		_, err = io.Copy(destFile, engine.Wrap(srcFile))
		if err != nil {
			return err
		}

		// log.Printf("Copied %s to %s", source, dest)

		return nil
	}

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	// log.Printf("Copied %s to %s", source, dest)

	return nil
}

func customCopyDir(source, dest, sedCmd string) error {
	sourceDirStat, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !sourceDirStat.IsDir() {
		return fmt.Errorf("%s is not a directory", source)
	}

	err = os.Mkdir(dest, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %s", err.Error())
	}

	files, err := ioutil.ReadDir(source)
	if err != nil {
		return fmt.Errorf("error reading directory: %e", err)
	}

	for _, f := range files {
		switch f.IsDir() {
		case true:
			sourceDir := path.Join(source, f.Name())
			destDir := path.Join(dest, f.Name())

			err = customCopyDir(sourceDir, destDir, sedCmd)
			if err != nil {
				return err
			}
		case false:
			sourceFile := path.Join(source, f.Name())
			destFile := path.Join(dest, f.Name())
			err = customCopyFile(sourceFile, destFile, sedCmd)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
