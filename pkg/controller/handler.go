package controller

import (
	"forum/pkg/service"
	"log"
	"net/http"
)

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
	mux.HandleFunc("/post/add", h.checkAccess(h.addPost))
	mux.HandleFunc("/post/get/", h.getPost)
	mux.HandleFunc("/post/update/", h.updatePost)
	mux.HandleFunc("/post/delete/", h.deletePost)
	mux.HandleFunc("/post/comment", h.addComment)
	mux.HandleFunc("/post/comment/get/", h.getComment)
	mux.HandleFunc("/post/comment/update/", h.updateComment)
	mux.HandleFunc("/post/comment/delete/", h.deleteComment)
	mux.HandleFunc("/post/comment/like/", h.deleteComment)

	return mux
}
