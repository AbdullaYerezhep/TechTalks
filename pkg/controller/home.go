package controller

import (
	"forum/models"
	"net/http"
)

type HomeData struct {
	User  models.User
	Posts []models.Post
}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	var data HomeData
	id := r.Context().Value(ctxKey("user_id"))
	if id != nil {
		user, err := h.srv.GetUserByID(id.(int))
		if err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "error", "")
		}
		data.User = user
	}
	posts, err := h.srv.GetAllPosts()
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "error", "")
		return
	}
	data.Posts = posts
	if err = templates["home"].Execute(w, data); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "error", "")
		return
	}
}
