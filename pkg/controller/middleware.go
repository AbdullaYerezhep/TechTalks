package controller

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"
)

const (
	defaultMode  = 0
	redirectMode = 1
)

type ctxKey string

const (
	keyUser    = ctxKey("user_id")
	keyPost    = ctxKey("post_id")
	keyComment = ctxKey("comment_id")
)

// Middleware "checkAccess" works on two modes. If it doesn't validate the token.
// 1. It redirects client to sign-in page
// 2. It still passes to the next HandlerFunc but without any data.
// If it does validate the token passes to the next HandlerFunc with user_id in context.
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
			if !errors.Is(err, sql.ErrNoRows) {
				h.errLog.Println(err.Error())
			}
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
		ctx := context.WithValue(r.Context(), keyUser, session.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// getPostId - sends post_id in context.
func (h *Handler) getPostID(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		post_id_str := r.URL.Query().Get("id")
		postID, err := strconv.Atoi(post_id_str)
		if err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusBadRequest, "error", "")
			return
		}

		ctx := context.WithValue(r.Context(), keyPost, postID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// getPostId - sends comment_id in context.
func (h *Handler) getCommentID(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		comment_id_any := r.URL.Query().Get("id")
		comment_id, err := strconv.Atoi(comment_id_any)
		if err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusBadRequest, "error", "")
			return
		}

		ctx := context.WithValue(r.Context(), keyComment, comment_id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
