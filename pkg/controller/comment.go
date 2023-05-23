package controller

import (
	"forum/models"
	"net/http"
	"forum/models"
)

func (h *Handler) addComment(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(keyUser).(int)
	com, ok := r.Context().Value(keyRequest).(models.Comment)
	if !ok {
		h.errLog.Println("Assertion: context > comment")
		h.errorMsg(w, http.StatusInternalServerError, errorTemp, "")
		return
	}

	com.User_ID = user_id

	if err := h.srv.AddComment(com); err != nil {
		h.errorMsg(w, http.StatusBadRequest, "error", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) editComment(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(keyUser)
	com, ok := r.Context().Value(keyRequest).(models.Comment)
	if !ok {
		h.errLog.Println("Context comment")
		h.errorMsg(w, http.StatusInternalServerError, errorTemp, "")
		return
	}

	if err := h.srv.UpdateComment(user_id.(int), com); err != nil {
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteComment(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(keyUser)
	com, ok := r.Context().Value(keyRequest).(models.Comment)
	if !ok {
		h.errLog.Println("Context comment")
		h.errorMsg(w, http.StatusInternalServerError, errorTemp, "")
		return
	}

	com, err := h.srv.GetComment(com.ID)
	if err != nil {
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	if err = h.srv.DeleteComment(user_id.(int), com.ID); err != nil {
		h.errorMsg(w, http.StatusNotFound, "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
