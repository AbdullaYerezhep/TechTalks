package controller

import (
	"encoding/json"
	"forum/models"
	"net/http"
)

func (h *Handler) ratePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}

	user_id := r.Context().Value(keyUser)
	decoder := json.NewDecoder(r.Body)

	var rate models.RatePost
	
	if err := decoder.Decode(&rate); err != nil {
		h.errorMsg(w, http.StatusBadRequest, "error", "Bad Request Body")
		return
	}
	rate.User_ID = user_id.(int)

	if err := h.srv.Post.RatePost(rate); err != nil {
		h.errorMsg(w, http.StatusBadRequest, errorTemp, err.Error())
		return
	}
	
}

func (h *Handler) rateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "")
		return
	}

	user_id := r.Context().Value(keyUser)
	decoder := json.NewDecoder(r.Body)

	var rate models.RateComment
	
	if err := decoder.Decode(&rate); err != nil {
		h.errorMsg(w, http.StatusBadRequest, "error", "Bad Request Body")
		return
	}
	rate.User_ID = user_id.(int)

	if err := h.srv.Comment.RateComment(rate); err != nil {
		h.errorMsg(w, http.StatusBadRequest, errorTemp, err.Error())
		return
	}

}