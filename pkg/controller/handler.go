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

	fs := http.FileServer(http.Dir("./view/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/sign-in", h.signIn)
	mux.HandleFunc("/sign-up", h.signUp)
	mux.HandleFunc("/logout", h.checkAccess(h.logOut, 0))
	mux.HandleFunc("/", h.checkAccess(h.home, 0))
	mux.HandleFunc("/post/add", h.checkAccess(h.addPost, 1))
	mux.HandleFunc("/post/update/", h.checkAccess(h.updatePost, 1))
	mux.HandleFunc("/post/delete/", h.checkAccess(h.deletePost, 1))
	mux.HandleFunc("/post/comment", h.checkAccess(h.addComment, 1))
	mux.HandleFunc("/post/comment/get/", h.checkAccess(h.getComment, 1))
	mux.HandleFunc("/post/comment/update/", h.checkAccess(h.updateComment, 1))
	mux.HandleFunc("/post/comment/delete/", h.checkAccess(h.deleteComment, 1))
	mux.HandleFunc("/post/comment/like/", h.checkAccess(h.likeDis, 1))

	return mux
}
