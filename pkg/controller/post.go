package controller

import (
	"forum/models"
	"net/http"
	"time"
)

func (h *Handler) addPost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates["post"].Execute(w, nil)

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusBadRequest, "error", "Form modified")
			return
		}
		id := r.Context().Value("user_id")
		user, err := h.srv.GetUserByID(id.(int))
		if err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "error", "")
			return
		}
		post := models.Post{
			User_ID: user.ID,
			Title:   r.FormValue("title"),
			Author:  user.Name,
			Content: r.FormValue("content"),
			Created: time.Now(),
		}
		if err = h.srv.CreatePost(post); err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "error", "")
		}
		http.Redirect(w, r, "/", http.StatusAccepted)

	default:
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "Posting error")
		return
	}
}

func (h *Handler) getPost(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
}
