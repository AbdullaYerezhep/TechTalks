package controller

import (
	"encoding/json"
	"forum/models"
	"net/http"
	"time"
)

func (h *Handler) addComment(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(keyUser)
	decoder := json.NewDecoder(r.Body)

	var com models.Comment

	if err := decoder.Decode(&com); err != nil {
		h.errorMsg(w, http.StatusBadRequest, "error", "Bad Request Body")
		return
	}

	post_id := r.Context().Value(keyPost)

	current_time := time.Now().Format("02-01-2006 15:04")
	com.Post_ID = post_id.(int)
	com.User_ID = user_id.(int)
	com.Created = current_time

	if err := h.srv.AddComment(com); err != nil {
		h.errorMsg(w, http.StatusBadRequest, "error", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) editComment(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(keyUser)
	decoder := json.NewDecoder(r.Body)

	var req models.Comment

	if err := decoder.Decode(&req); err != nil {
		h.errorMsg(w, http.StatusBadRequest, "error", "Bad Request Body")
		return
	}

	comment_id := req.ID

	com, err := h.srv.GetComment(comment_id)
	if err != nil {
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	if com.User_ID != user_id.(int) {
		h.errorMsg(w, http.StatusBadRequest, "error", "")
		return
	}

	current_time := time.Now().Format("02-01-2006 15:04")
	com.Content = req.Content
	com.Updated = &current_time

	if err = h.srv.UpdateComment(com); err != nil {
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteComment(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(keyUser)
	decoder := json.NewDecoder(r.Body)

	var req models.Comment

	if err := decoder.Decode(&req); err != nil {
		h.errorMsg(w, http.StatusBadRequest, "error", "Bad Request Body")
		return
	}

	comment_id := req.ID

	com, err := h.srv.GetComment(comment_id)
	if err != nil {
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	if com.User_ID != user_id.(int) {
		h.errorMsg(w, http.StatusBadRequest, "error", "")
		return
	}

	if err = h.srv.DeleteComment(com.ID); err != nil {
		h.errorMsg(w, http.StatusNotFound, "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
