package controller

import (
	"forum/models"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		if err := templates["sign-up"].Execute(w, nil); err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "error", "")
			return
		}

	case http.MethodPost:
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

		if err := h.srv.CreateUser(user); err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
			return
		}

	default:
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		if err := templates["sign-in"].Execute(w, nil); err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "error", "")
			return
		}

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.errorMsg(w, http.StatusBadRequest, "error", "")
			return
		}
		name := r.FormValue("username")
		password := r.FormValue("password")

		user, err := h.srv.GetUser(name, password)
		if err != nil {
			h.errorMsg(w, http.StatusBadRequest, "sign-in", "User not found!")
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
			Name:  "token",
			Value: s.Token,
		}

		http.SetCookie(w, c)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	default:
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}
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
