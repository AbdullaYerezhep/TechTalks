package controller

import (
	"net/http"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		h.errLog.Println(err.Error())
		http.Error(w, "ISE", http.StatusInternalServerError)
		return
	}
	user, err := h.srv.GetUserByToken(cookie.Value)
	if err != nil {
		h.errLog.Println(err.Error())
		return
	}
	templates["home"].Execute(w, user)
}
