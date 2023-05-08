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
	mux.HandleFunc("/post/update/", h.getPostID(h.checkAccess(h.updatePost, 1)))
	mux.HandleFunc("/post/delete/", h.getPostID(h.checkAccess(h.deletePost, 1)))
	mux.HandleFunc("/post/rate/", h.getPostID(h.checkAccess(h.ratePost, 1)))
	mux.HandleFunc("/post/comment/", h.getPostID(h.checkAccess(h.addComment, 1)))
	// mux.HandleFunc("/post/comment/update/", h.getCommentID(h.checkAccess(h.updateComment, 1)))
	mux.HandleFunc("/post/comment/delete/", h.getCommentID(h.checkAccess(h.deleteComment, 1)))
	mux.HandleFunc("/post/comment/rate/", h.getCommentID(h.checkAccess(h.rateComment, 1)))

	return mux
}
