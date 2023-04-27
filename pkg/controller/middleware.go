package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (h *Handler) checkAccess(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			h.infoLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "error", "")
			return
		}
		session, err := h.srv.Session.GetSession(token.Value)
		if err != nil {
			h.errLog.Println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
		}
		if session.Expiration_date.Before(time.Now()) {
			fmt.Println(session.Expiration_date)
			err = h.srv.Session.DeleteSession(session.ID)
			if err != nil {
				h.errLog.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
		}
		ctx := context.WithValue(r.Context(), "user_id", session.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
