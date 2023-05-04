package controller

import (
	"context"
	"net/http"
	"time"
)

const (
	defaultMode  = 0
	redirectMode = 1
)

type ctxKey string

func (h *Handler) checkAccess(next http.HandlerFunc, mode int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			if mode == defaultMode {
				next.ServeHTTP(w, r)
				return
			}
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}
		session, err := h.srv.Session.GetSession(token.Value)
		if err != nil {
			if mode == defaultMode {
				next.ServeHTTP(w, r)
				return
			}
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}
		if session.Expiration_date.Before(time.Now()) {
			err = h.srv.Session.DeleteSession(session.UserId)
			if err != nil {
				h.errLog.Println(err.Error())
				h.errorMsg(w, http.StatusInternalServerError, "error", "")
				return
			}
			if mode == defaultMode {
				next.ServeHTTP(w, r)
				return
			}
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ctxKey("user_id"), session.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
