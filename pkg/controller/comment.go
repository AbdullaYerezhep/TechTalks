package controller

import (
	"fmt"
	"forum/models"
	"net/http"
)

func (h *Handler) addComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorMsg(w, http.StatusMethodNotAllowed, "")
		return
	} else if r.URL.Path != "/comment/add" {
		h.errorMsg(w, http.StatusNotFound, "")
		return
	}

	user_id := r.Context().Value(keyUser).(int)
	com, ok := r.Context().Value(keyRequest).(models.Comment)
	if !ok {
		h.errLog.Println("Assertion: context > comment")
		h.errorMsg(w, http.StatusInternalServerError, "")
		return
	}

	com.User_ID = user_id

	if err := h.srv.AddComment(com); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

// edit comment
func (h *Handler) editComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		h.errorMsg(w, http.StatusMethodNotAllowed, "")
		return
	} else if r.URL.Path != "/comment/add" {
		h.errorMsg(w, http.StatusNotFound, "")
		return
	}

	user_id := r.Context().Value(keyUser)
	com, ok := r.Context().Value(keyRequest).(models.Comment)
	if !ok {
		h.errLog.Println("Context comment")
		h.errorMsg(w, http.StatusInternalServerError, "")
		return
	}
	fmt.Println(com)

	if err := h.srv.UpdateComment(user_id.(int), com); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// delete comment
func (h *Handler) deleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.errorMsg(w, http.StatusMethodNotAllowed, "")
		return
	} else if r.URL.Path != "/comment/add" {
		h.errorMsg(w, http.StatusNotFound, "")
		return
	}

	user_id := r.Context().Value(keyUser)
	com, ok := r.Context().Value(keyRequest).(models.Comment)
	if !ok {
		h.errLog.Println("Context comment")
		h.errorMsg(w, http.StatusInternalServerError, "")
		return
	}

	com, err := h.srv.GetComment(com.ID)
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "")
		return
	}

	if err = h.srv.DeleteComment(user_id.(int), com.ID); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusNotFound, "")
		return
	}

	w.WriteHeader(http.StatusOK)
}
