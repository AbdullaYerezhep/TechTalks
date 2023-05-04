package controller

import (
	"net/http"
	"strconv"
)

func (h *Handler) likePost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(ctxKey("user_id"))
	post_id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.errorMsg(w, http.StatusBadRequest, errorTemp, "")
		return
	}
	if err = h.srv.Post.LikeDis(id.(int), post_id, 1); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, errorTemp, "")
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) dislikePost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(ctxKey("user_id"))
	post_id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.errorMsg(w, http.StatusBadRequest, errorTemp, "")
		return
	}
	if err = h.srv.Post.LikeDis(id.(int), post_id, 0); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, errorTemp, "")
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
