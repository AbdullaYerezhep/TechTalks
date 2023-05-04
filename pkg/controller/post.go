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
		var categories []string
		categories, _ = h.srv.Post.GetCategories()
		templates["addpost"].Execute(w, categories)

	case http.MethodPost:
		id := r.Context().Value(ctxKey("user_id"))
		user, err := h.srv.GetUserByID(id.(int))
		if err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "error", "")
			return
		}

		if err := r.ParseForm(); err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusBadRequest, "error", "Form modified")
			return
		}

		currentTime := time.Now()
		categories := r.Form["category[]"]
		if len(categories) == 0 {
			h.errorMsg(w, http.StatusBadRequest, "error", "")
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
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "error", "")
			return
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
			h.errorMsg(w, http.StatusBadRequest, "error", "")
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
	id := r.Context().Value(ctxKey("user_id"))
	post_id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, "error", "")
		return
	}

	post, err := h.srv.GetPost(post_id)
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusNotFound, "error", "")
		return
	}

	if id.(int) != post.User_ID {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, "error", "")
		return
	}

	if err = h.srv.DeletePost(post.ID); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusNotFound, "error", "")
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
