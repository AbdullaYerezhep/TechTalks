package controller

import (
	"fmt"
	"html/template"
	"log"
	"path/filepath"
)

var templates = make(map[string]*template.Template)

func init() {
	files, err := filepath.Glob("view/templates/*page.html")
	if err != nil {
		log.Fatal("Error parsing tempaltes")
		return
	}
	for _, file := range files {
		name := filepath.Base(file)
		name = name[:len(name)-10]

		tmpl, err := template.ParseFiles(file)
		if err != nil {
			log.Fatal(err.Error())
		}

		tmpl, err = tmpl.ParseGlob("view/templates/common.html")

		if err != nil {
			log.Fatal(err.Error())
		}

		tmpl, err = tmpl.ParseGlob("view/templates/base.html")

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("template parsed: " + name)
		templates[name] = tmpl
	}
}
