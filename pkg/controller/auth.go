package controller

import (
	"fmt"
	"forum/models"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
	switch r.Method {
	case http.MethodGet:
		templates["sign-in"].Execute(w, cookie)
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, err)
			return
		}
		name := r.PostFormValue("username")

		_, err := h.srv.GetByName(name)
		if err != nil {
			h.errLog.Println(err.Error())
		}

		token, err := uuid.NewV4()
		if err != nil {
			log.Println("Error creating token")
		}

		// expiration_date := time.Now().Add(time.Minute * 1)

		c := &http.Cookie{
			Name:  "token",
			Value: token.String(),
		}
		http.SetCookie(w, c)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := templates["sign-up"].Execute(w, nil); err != nil {
			h.errLog.Println(err.Error())
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.errLog.Println(err.Error())
			return
		}
		fmt.Println("Parsed")
		if r.PostFormValue("password") != r.PostFormValue("confirmPass") {
			fmt.Fprint(w, "Failed conformition")
			return
		}

		var u models.User
		u = models.User{
			Name:     r.PostFormValue("username"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}

		h.srv.CreateUser(&u)
		c := &http.Cookie{
			Name:  "username",
			Value: u.Name,
		}
		http.SetCookie(w, c)
		http.Redirect(w, r, "/sign-in", http.StatusFound)
	}
}
