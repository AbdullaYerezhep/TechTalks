package controller

import (
	"forum/models"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) addPost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates["addpost"].Execute(w, nil)

	case http.MethodPost:
		w.Header().Set("Content-type", "")
		if err := r.ParseForm(); err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusBadRequest, "error", "Form modified")
			return
		}
		id := r.Context().Value(ctxKey("user_id"))

		user, err := h.srv.GetUserByID(id.(int))
		if err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "error", "")
			return
		}
		post := models.Post{
			User_ID: user.ID,
			Title:   r.FormValue("title"),
			Author:  user.Name,
			Content: r.FormValue("content"),
			Created: time.Now(),
			Updated: time.Now(),
		}
		if err = h.srv.CreatePost(post); err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "error", "")
		}
		http.Redirect(w, r, "/", http.StatusFound)

	default:
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
	post_id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.errorMsg(w, http.StatusBadRequest, "error", "")
		return
	}
	switch r.Method {

	case http.MethodGet:
		post, err := h.srv.GetPost(post_id)
		if err != nil {
			h.errorMsg(w, http.StatusNotFound, "error", "")
			return
		}
		if err = templates["post"].Execute(w, post); err != nil {
			h.errorMsg(w, http.StatusInternalServerError, "error", "")
			return
		}

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.errorMsg(w, http.StatusBadRequest, "error", "")
			return
		}
		id := r.Context().Value(ctxKey("user_id"))
		post, err := h.srv.GetPost(post_id)
		if err != nil {
			h.errorMsg(w, http.StatusNotFound, "error", "")
			return
		}
		if post.User_ID != id.(int) {
			return
		}
		post.Title = r.FormValue("title")
		post.Content = r.FormValue("content")
		post.Updated = time.Now()
		h.srv.UpdatePost(post)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}

}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(ctxKey("user_id"))
	post_id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.errorMsg(w, http.StatusBadRequest, "error", "")
		return
	}
	post, err := h.srv.GetPost(post_id)
	if err != nil {
		h.errorMsg(w, http.StatusNotFound, "error", "")
		return
	}
	if id.(int) != post.User_ID {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	if err = h.srv.DeletePost(post_id); err != nil {
		h.errorMsg(w, http.StatusNotFound, "error", "")
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
