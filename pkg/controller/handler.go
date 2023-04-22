package controller

import (
	"fmt"
	"forum/pkg/service"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
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
			log.Fatal("IDK what happened")
		}
		fmt.Println(name)
		templates[name] = tmpl
	}
}

type Handler struct {
	infoLog *log.Logger
	errLog  *log.Logger
	srv     *service.Service
}

func New(info, err *log.Logger, srv *service.Service) *Handler {
	return &Handler{
		infoLog: info,
		errLog:  err,
		srv:     srv,
	}
}

func (h *Handler) Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/sign-in", h.signIn)
	mux.HandleFunc("/sign-up", h.signUp)
	mux.HandleFunc("/", h.home)
	return mux
}

func verifyPass(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
