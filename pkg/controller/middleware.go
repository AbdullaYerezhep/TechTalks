package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"forum/models"
	"net/http"
	"strings"
	"time"
)

const (
	defaultMode  = 0
	redirectMode = 1
)

type ctxKey string

const (
	keyUser    = ctxKey("user_id")
	keyRequest = ctxKey("requestData")
)

// Middleware "checkAccess" works on two modes. If it doesn't validate the token.
// 1. It redirects client to sign-in page
// 2. It still passes to the next HandlerFunc but without any data.
// If it does validate the token passes to the next HandlerFunc with user_id in context.
func (h *Handler) checkAccess(next http.HandlerFunc, mode int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err == http.ErrNoCookie {
			if mode == defaultMode {
				next.ServeHTTP(w, r)
				return
			} else {
				h.errLog.Println(err.Error())
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		session, err := h.srv.Session.GetSession(token.Value)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				h.errLog.Println(err.Error())
			}
			if mode == defaultMode {
				next.ServeHTTP(w, r)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if session.Expiration_date.Before(time.Now()) {
			err = h.srv.Session.DeleteSession(session.UserId)
			if err != nil {
				h.errLog.Println(err.Error())
				h.errorMsg(w, http.StatusInternalServerError, "")
				return
			}
			if mode == defaultMode {
				next.ServeHTTP(w, r)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), keyUser, session.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
		return
	})
}

func (h *Handler) decodeRequest(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var post models.Post
		var com models.Comment
		ctx := context.WithValue(r.Context(), keyRequest, nil)

		if strings.HasPrefix(r.URL.Path, "/post") {
			if err := decoder.Decode(&post); err != nil {
				h.errLog.Println(err)
				h.errorMsg(w, http.StatusBadRequest, "")
				return
			}
			ctx = context.WithValue(r.Context(), keyRequest, post)

		} else if strings.HasPrefix(r.URL.Path, "/comment") {
			if err := decoder.Decode(&com); err != nil {
				h.errLog.Println(err)
				h.errorMsg(w, http.StatusBadRequest, "")
				return
			}
			ctx = context.WithValue(r.Context(), keyRequest, com)

		} else {
			next.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
