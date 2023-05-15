package controller

import (
	"forum/models"
	"net/http"
	"time"
)

func (h *Handler) addComment(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.errorMsg(w, http.StatusBadRequest, "error", err.Error())
		return
	}

	user_id := r.Context().Value(keyUser)
	post_id := r.Context().Value(keyPost)
	current_time := time.Now().Format("02-01-2006 15:04")
	comment := models.Comment{
		User_ID: user_id.(int),
		Post_ID: post_id.(int),
		Content: r.FormValue("comment-content"),
		Created: current_time,
	}

	if err := h.srv.AddComment(comment); err != nil {
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

}

func PostComments(w http.ResponseWriter, r *http.Request) {
}

// func (h *Handler) updateComment(w http.ResponseWriter, r *http.Request) {
// 	user_id := r.Context().Value(keyUser)
// 	comment_id := r.Context().Value(keyComment)

// 	com, err := h.srv.GetComment(comment_id.(int))
// 	if err != nil {
// 		h.errLog.Println(err.Error())
// 		h.errorMsg(w, http.StatusInternalServerError, "error", "")
// 		return
// 	}

// 	if com.User_ID != user_id.(int) {
// 		h.errorMsg(w, http.StatusBadRequest, "error", "")
// 		return
// 	}

// 	com.Content = r.FormValue("comment-content")

// 	err = h.srv.UpdateComment(com)
// 	if err != nil {
// 		h.errorMsg(w, http.StatusInternalServerError, "error", "")
// 		return
// 	}

// 	referer := r.Header.Get("Referer")

// 	// Redirect the user to the previous URL
// 	http.Redirect(w, r, referer, http.StatusSeeOther)
// }

func (h *Handler) deleteComment(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(keyUser)
	comment_id := r.Context().Value(keyComment)

	com, err := h.srv.GetComment(comment_id.(int))
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "error", "")
		return
	}

	if com.User_ID != user_id.(int) {
		h.errorMsg(w, http.StatusBadRequest, "error", "")
		return
	}

	if err = h.srv.DeleteComment(com.ID); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusNotFound, "error", "")
		return
	}


	// Redirect the user to the previous URL
}
