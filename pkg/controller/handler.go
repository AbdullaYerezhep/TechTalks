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

// func (h *Handler) Router(r *router.Router) {
// 	r.Post("/sign-up", h.signUp)
// 	r.Post("/sign-in", h.signIn)
// 	r.Post("/logout", h.logOut)

// 	r.Get("/", h.home)
// 	r.Post("/", h.home)

// 	r.Get("/post/", h.postDetails)
// 	r.Post("/post/add", h.addPost)
// 	r.Put("/post/edit", h.updatePost)
// 	r.Delete("/post/delete", h.deletePost)
// 	r.Post("/post/rate", h.ratePost)

// 	r.Post("/comment", h.addComment)
// 	r.Put("/comment/edit", h.editComment)
// 	r.Delete("/comment/delete", h.addComment)
// 	r.Post("/comment/rate", h.rateComment)
// }

func (h *Handler) Router() *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./view/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/sign-in", h.signIn)
	mux.HandleFunc("/sign-up", h.signUp)
	mux.HandleFunc("/logout", h.checkAccess(h.logOut, 0))
	mux.HandleFunc("/", h.checkAccess(h.home, 0))
	mux.HandleFunc("/post/", h.checkAccess(h.postDetails, 0))
	mux.HandleFunc("/post/add", h.decodeRequest(h.checkAccess(h.addPost, 1)))
	mux.HandleFunc("/post/edit", h.decodeRequest(h.checkAccess(h.updatePost, 1)))
	mux.HandleFunc("/post/delete", h.decodeRequest(h.checkAccess(h.deletePost, 1)))
	mux.HandleFunc("/post/rate", h.checkAccess(h.ratePost, 1))
	mux.HandleFunc("/comment", h.decodeRequest(h.checkAccess(h.addComment, 1)))
	mux.HandleFunc("/comment/edit", h.decodeRequest(h.checkAccess(h.editComment, 1)))
	mux.HandleFunc("/comment/delete", h.decodeRequest(h.checkAccess(h.deleteComment, 1)))
	mux.HandleFunc("/comment/rate", h.checkAccess(h.rateComment, 1))

	return mux
}
