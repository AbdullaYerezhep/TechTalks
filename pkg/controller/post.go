package controller

import (
	"forum/models"
	"net/http"
	"time"
)

func (h *Handler) addPost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(keyUser)
	switch r.Method {
	case http.MethodGet:
		var categories []string
		categories, _ = h.srv.Post.GetCategories()
		user, _ := h.srv.GetUserByID(id.(int))
		addPost := models.AddPostPage{
			User: user,
			Categories: categories,
		}
		if err := templates["addpost"].Execute(w, addPost); err != nil {
			h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
			return
		}

	case http.MethodPost:
		user, err := h.srv.GetUserByID(id.(int))
		if err != nil {
			h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
			return
		}

		if err := r.ParseForm(); err != nil {
			h.errorMsg(w, http.StatusBadRequest, "error", err.Error())
			return
		}

		currentTime := time.Now()
		categories := r.Form["category[]"]
		if len(categories) == 0 {
			h.errorMsg(w, http.StatusBadRequest, "error", err.Error())
			return
		}

		post := models.Post{
			User_ID:  user.ID,
			Title:    r.FormValue("title"),
			Author:   user.Name,
			Category: categories,
			Content:  r.FormValue("content"),
			Created:  currentTime,
			Updated:  currentTime,
		}

		if err = h.srv.CreatePost(post); err != nil {
			h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)

	default:
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
	post_id := r.Context().Value(keyPost)

	switch r.Method {

	case http.MethodGet:
		post, err := h.srv.GetPost(post_id.(int))
		if err != nil {
			h.errorMsg(w, http.StatusNotFound, "error", err.Error())
			return
		}
		if err = templates["update-post"].Execute(w, post); err != nil {
			h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
			return
		}

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.errorMsg(w, http.StatusBadRequest, "error", err.Error())
			return
		}
		user_id := r.Context().Value(keyUser)
		post, err := h.srv.GetPost(post_id.(int))
		if err != nil {
			h.errorMsg(w, http.StatusNotFound, "error", err.Error())
			return
		}
		if post.User_ID != user_id.(int) {
			h.errorMsg(w, http.StatusBadRequest, "error", err.Error())
			return
		}
		post.Title = r.FormValue("title")
		post.Content = r.FormValue("content")
		post.Updated = time.Now()
		err = h.srv.UpdatePost(post)
		if err != nil {
			h.errorMsg(w, http.StatusBadRequest, "post", "")
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(keyUser)
	post_id := r.Context().Value(keyPost)

	if err := h.srv.DeletePost(user_id.(int), post_id.(int)); err != nil {
		h.errorMsg(w, http.StatusNotFound, "error", err.Error())
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
