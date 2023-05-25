package controller

import (
	"fmt"
	"forum/models"
	"net/http"
)

func (h *Handler) addComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorMsg(w, http.StatusMethodNotAllowed, errorTemp, "")
		return
	}

	user_id := r.Context().Value(keyUser).(int)
	com, ok := r.Context().Value(keyRequest).(models.Comment)
	if !ok {
		h.errLog.Println("Assertion: context > comment")
		h.errorMsg(w, http.StatusInternalServerError, errorTemp, "")
		return
	}

	com.User_ID = user_id

	if err := h.srv.AddComment(com); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, "error", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

// edit comment
func (h *Handler) editComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		h.errorMsg(w, http.StatusMethodNotAllowed, errorTemp, "")
		return
	}

	user_id := r.Context().Value(keyUser)
	com, ok := r.Context().Value(keyRequest).(models.Comment)
	if !ok {
		h.errLog.Println("Context comment")
		h.errorMsg(w, http.StatusInternalServerError, errorTemp, "")
		return
	}
	fmt.Println(com)

	if err := h.srv.UpdateComment(user_id.(int), com); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// delete comment
func (h *Handler) deleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.errorMsg(w, http.StatusMethodNotAllowed, errorTemp, "")
		return
	}

	user_id := r.Context().Value(keyUser)
	com, ok := r.Context().Value(keyRequest).(models.Comment)
	if !ok {
		h.errLog.Println("Context comment")
		h.errorMsg(w, http.StatusInternalServerError, errorTemp, "")
		return
	}

	com, err := h.srv.GetComment(com.ID)
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	if err = h.srv.DeleteComment(user_id.(int), com.ID); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusNotFound, "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
