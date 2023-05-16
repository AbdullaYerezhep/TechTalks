package controller

import (
	"forum/models"
	"net/http"
	"strconv"
)

func (h *Handler) ratePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}

	if err := r.ParseForm(); err != nil {
		h.errorMsg(w, http.StatusBadRequest, "sign-up", err.Error())
		return
	}

	var islike int8
	post_id_str := r.FormValue("post_id")
	action := r.FormValue("action")
	switch action {
	case "wow":
		islike = 1
	case "boo":
		islike = -1
	}

	user_id := r.Context().Value(keyUser)
	post_id, _ := strconv.Atoi(post_id_str)

	rate := models.RatePost{
		User_ID: user_id.(int),
		Post_ID: post_id,
		IsLike:  islike,
	}

	if err := h.srv.Post.RatePost(rate); err != nil {
		h.errorMsg(w, http.StatusBadRequest, errorTemp, err.Error())
		return
	}
	referer := r.Header.Get("Referer")
	http.Redirect(w, r, referer, http.StatusSeeOther)
}

func (h *Handler) rateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}

	if err := r.ParseForm(); err != nil {
		h.errorMsg(w, http.StatusBadRequest, "sign-up", err.Error())
		return
	}

	var islike int8
	comment_id_str := r.FormValue("comment_id")
	action := r.FormValue("action")
	switch action {
	case "wow":
		islike = 1
	case "boo":
		islike = -1
	}

	user_id := r.Context().Value(keyUser)
	comment_id, _ := strconv.Atoi(comment_id_str)

	rate := models.RateComment{
		User_ID:    user_id.(int),
		Comment_ID: comment_id,
		IsLike:     islike,
	}

	if err := h.srv.Comment.RateComment(rate); err != nil {
		h.errorMsg(w, http.StatusBadRequest, errorTemp, err.Error())
		return
	}

	referer := r.Header.Get("Referer")
	http.Redirect(w, r, referer, http.StatusSeeOther)
}
