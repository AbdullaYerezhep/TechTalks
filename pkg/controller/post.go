package controller

import (
	"forum/models"
	"net/http"
)

func (h *Handler) addPost(w http.ResponseWriter, r *http.Request) {
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
			h.errorMsg(w, http.StatusInternalServerError, errorTemp, err.Error())
			return
		}

	case http.MethodPost:
		user, err := h.srv.GetUserByID(id.(int))
		if err != nil {
			h.errorMsg(w, http.StatusInternalServerError, errorTemp, err.Error())
			return
		}

		post, ok := r.Context().Value(keyRequest).(models.Post)
		if !ok {
			h.errLog.Println("Context post")
			h.errorMsg(w, http.StatusInternalServerError, errorTemp, "")
			return
		}

		post.User_ID = user.ID
		post.Author = user.Name

		if err = h.srv.CreatePost(post); err != nil {
			h.errorMsg(w, http.StatusInternalServerError, errorTemp, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		h.errorMsg(w, http.StatusMethodNotAllowed, errorTemp, "")
		return
	}
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		h.errorMsg(w, http.StatusMethodNotAllowed, errorTemp, "")
		return
	}

	user_id := r.Context().Value(keyUser)
	updatedPost, ok := r.Context().Value(keyRequest).(models.Post)
	if !ok {
		h.errLog.Println("Context post")
		h.errorMsg(w, http.StatusInternalServerError, errorTemp, "")
		return
	}

	if err := h.srv.UpdatePost(user_id.(int), updatedPost); err != nil {
		h.errorMsg(w, http.StatusBadRequest, errorTemp, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.errorMsg(w, http.StatusMethodNotAllowed, errorTemp, "")
		return
	}

	user_id := r.Context().Value(keyUser)
	post, ok := r.Context().Value(keyRequest).(models.Post)
	if !ok {
		h.errLog.Println("Context post")
		h.errorMsg(w, http.StatusInternalServerError, errorTemp, "")
		return
	}

	if err := h.srv.DeletePost(user_id.(int), post.ID); err != nil {
		h.errorMsg(w, http.StatusNotFound, errorTemp, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
