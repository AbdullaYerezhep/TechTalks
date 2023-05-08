package controller

import (
	"forum/models"
	"net/http"
)

func (h *Handler) ratePost(w http.ResponseWriter, r *http.Request) {
	var islike int8
	switch r.URL.Path {
	case "/post/rate/wow":
		islike = 1
	case "/post/rate/boo":
		islike = -1
	default:
		h.errorMsg(w, http.StatusNotFound, "error", "")
	}

	user_id := r.Context().Value(keyUser)
	post_id := r.Context().Value(keyPost)
	rate := models.RatePost{
		User_ID: user_id.(int),
		Post_ID: post_id.(int),
		IsLike:  islike,
	}

	if err := h.srv.Post.RatePost(rate); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, errorTemp, "")
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) rateComment(w http.ResponseWriter, r *http.Request) {
}
