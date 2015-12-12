package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "start_ios_project"
	app.Version = "1.0.0"
	app.Usage = "generate Gemfile/Podfile/.gitignore"
	app.Action = func(c *cli.Context) {
		copyFiles()
	}
	app.Run(os.Args)
}

func getSourceDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}

func copyFile(path string, name string) {
	content, error := ioutil.ReadFile(path)
	if error != nil {
		panic(error)
	}

	error = ioutil.WriteFile(name, content, 0644)
	if error != nil {
		panic(error)
	}

	fmt.Printf("generate %s.\n", name)
}

func copyFiles() {
	files := map[string]string{
		"Gemfile":   "Gemfile",
		"Podfile":   "Podfile",
		"gitignore": ".gitignore",
	}

	templatesDir := filepath.Join(getSourceDir(), "templates")

	for key, value := range files {
		copyFile(filepath.Join(templatesDir, key), value)
	}
}
