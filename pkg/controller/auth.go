package controller

import (
	"fmt"
	"forum/models"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}

	if err := r.ParseForm(); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, "sign-up", err.Error())
		return
	}

	password := r.FormValue("password")
	connfirmPass := r.FormValue("confirmPass")

	if password != connfirmPass {
		return
	}

	user := models.User{
		Name:     r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	user_id, err := h.srv.CreateUser(user)
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, errorTemp, err.Error())
		return
	}

	s := newSession(user_id)
	err = h.srv.CreateSession(s)
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "error", "")
		return
	}
	h.infoLog.Println("Session created: ", user.Name)

	c := &http.Cookie{
		Name:    "token",
		Value:   s.Token,
		Expires: time.Now().Add(1 * time.Hour),
	}

	http.SetCookie(w, c)

	referer := w.Header().Get("Referer")
	http.Redirect(w, r, referer, http.StatusSeeOther)
	return
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}

	if err := r.ParseForm(); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, "error", "")
		return
	}
	name := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(name, password)
	user, err := h.srv.GetUser(name, password)
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusNotFound, errorTemp, "User not found")
		// w.Header().Set("Error-msg", err.Error())
		return
	}

	s := newSession(user.ID)
	err = h.srv.CreateSession(s)
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "error", "")
		return
	}
	h.infoLog.Println("Session created: ", user.Name)

	c := &http.Cookie{
		Name:    "token",
		Value:   s.Token,
		Expires: time.Now().Add(1 * time.Hour),
	}

	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func (h *Handler) logOut(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(keyUser)
	err := h.srv.DeleteSession(user_id.(int))
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "error", "")
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func newSession(user_id int) models.Session {
	var s models.Session
	s.UserId = user_id
	token, _ := uuid.NewV4()
	s.Token = token.String()
	s.Expiration_date = time.Now().Add(1 * time.Hour)
	return s
}
