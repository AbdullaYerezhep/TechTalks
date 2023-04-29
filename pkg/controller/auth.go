package controller

import (
	"forum/models"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
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
		user, err := decodeForm(r.Form)
		if err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusBadRequest, "sign-up", err.Error())
			return
		}

		if err := h.srv.CreateUser(user); err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
			return
		}
		http.Redirect(w, r, "/sign-in", http.StatusFound)

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

		user, err := h.srv.GetUser(name)
		if err != nil || !verifyPass(password, user.Password) {
			h.errorMsg(w, http.StatusBadRequest, "sign-in", "Invalid data")
			return
		}

		s := newSession(user.ID)
		err = h.srv.CreateSession(s)
		if err != nil {
			h.errorMsg(w, http.StatusInternalServerError, "error", "")
			return
		}
		h.infoLog.Println("Session created: ", user.Name)

		c := &http.Cookie{
			Name:  "token",
			Value: s.Token,
		}
		http.SetCookie(w, c)
		http.Redirect(w, r, "/", http.StatusFound)

	default:
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}
}

func (h *Handler) logOut(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(ctxKey("user_id"))
	user, err := h.srv.GetUserByID(id.(int))
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "error", "")
		return
	}
	// get session by user id zhazau kerek
	err = h.srv.DeleteSession(user.ID)
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "error", "")
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
	return
}

func newSession(user_id int) models.Session {
	var s models.Session
	s.UserId = user_id
	token, _ := uuid.NewV4()
	s.Token = token.String()
	s.Expiration_date = time.Now().Add(1 * time.Hour)
	return s
}

func verifyPass(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
