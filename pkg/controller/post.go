package controller

import (
	"forum/models"
	"net/http"
)

func (h *Handler) addPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/add" {
		h.errorMsg(w, http.StatusNotFound, "")
		return
	}

	id := r.Context().Value(keyUser)

	switch r.Method {

	case http.MethodGet:
		var categories []string
		categories, _ = h.srv.Post.GetCategories()
		user, _ := h.srv.GetUserByID(id.(int))
		addPost := models.AddPostPage{
			User:       user,
			Categories: categories,
		}

		if err := templates["addpost"].Execute(w, addPost); err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "")
			return
		}

	case http.MethodPost:

		user, err := h.srv.GetUserByID(id.(int))
		if err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "")
			return
		}

		post, ok := r.Context().Value(keyRequest).(models.Post)
		if !ok {
			h.errLog.Println("Context post")
			h.errorMsg(w, http.StatusInternalServerError, "")
			return
		}

		post.User_ID = user.ID
		post.Author = user.Name

		if err = h.srv.CreatePost(post); err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusBadRequest, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		h.errorMsg(w, http.StatusMethodNotAllowed, "")
		return
	}
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		h.errorMsg(w, http.StatusMethodNotAllowed, "")
		return
	} else if r.URL.Path != "/post/edit" {
		h.errorMsg(w, http.StatusNotFound, "")
		return
	}

	user_id := r.Context().Value(keyUser)
	updatedPost, ok := r.Context().Value(keyRequest).(models.Post)
	if !ok {
		h.errLog.Println("Context post")
		h.errorMsg(w, http.StatusInternalServerError, "")
		return
	}

	if err := h.srv.UpdatePost(user_id.(int), updatedPost); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.errorMsg(w, http.StatusMethodNotAllowed, "")
		return
	} else if r.URL.Path != "/post/delete" {
		h.errorMsg(w, http.StatusNotFound, "")
		return
	}

	user_id := r.Context().Value(keyUser)
	post, ok := r.Context().Value(keyRequest).(models.Post)
	if !ok {
		h.errLog.Println("Context post")
		h.errorMsg(w, http.StatusInternalServerError, "")
		return
	}

	if err := h.srv.DeletePost(user_id.(int), post.ID); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
