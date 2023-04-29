package controller

import (
	"fmt"
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
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.errLog.Println(err.Error())
			// http.Error(w, err.Error(), http.StatusBadRequest)
			h.errorMsg(w, http.StatusBadRequest, "sign-up", err.Error())
			return
		}
		user, err := decodeForm(r.Form)
		if err != nil {
			h.errLog.Println(err.Error())
			// http.Error(w, err.Error(), http.StatusBadRequest)
			h.errorMsg(w, http.StatusBadRequest, "sign-up", err.Error())
			return
		}

		if err := h.srv.CreateUser(user); err != nil {
			h.errLog.Println(err.Error())
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
			return
		}
		http.Redirect(w, r, "/sign-in", http.StatusFound)

	default:
		// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		templates["sign-in"].Execute(w, nil)

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, err)
			return
		}
		name := r.FormValue("username")
		password := r.FormValue("password")

		user, err := h.srv.GetUser(name)
		if err != nil {
			h.errLog.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !verifyPass(password, user.Password) {
			http.Error(w, "CONFIRM failed", http.StatusBadRequest)
			return
		}

		s := newSession(user.ID)
		err = h.srv.CreateSession(s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
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
