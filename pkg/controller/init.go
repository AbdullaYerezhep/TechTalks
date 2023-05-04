package controller

import (
	"fmt"
	"html/template"
	"log"
	"path/filepath"
)

const (
	errorTemp   = "error"
	signInTemp  = "sign-in"
	signUpTemp  = "sign-up"
	homeTemp    = "home"
	postTemp    = "post"
	addPostTemp = "addPost"
)

var templates = make(map[string]*template.Template)

func init() {
	files, err := filepath.Glob("view/templates/*.html")
	if err != nil {
		log.Fatal("Error parsing tempaltes")
		return
	}
	for _, file := range files {
		name := filepath.Base(file)
		name = name[:len(name)-5]
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("template parsed: " + name)
		templates[name] = tmpl
	}
}
