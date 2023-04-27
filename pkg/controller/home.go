package controller

import (
	"net/http"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	posts, err := h.srv.GetAllPosts()
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "error", "")
		return
	}
	templates["home"].Execute(w, posts)
}
