package controller

import (
	"forum/models"
	"net/http"
	"net/url"
	"strconv"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	var data models.HomePage
	id := r.Context().Value(keyUser)
	if id != nil {
		user, err := h.srv.GetUserByID(id.(int))
		if err != nil {
			h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
			return
		}
		data.User = user
	}

	posts, err := h.srv.GetAllPosts()
	if err != nil {
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	data.Posts = posts

	categories, err := h.srv.GetCategories()
	if err != nil {
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}
	data.Categories = categories

	if err = templates["home"].Execute(w, data); err != nil {
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}
}

func (h *Handler) postDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorMsg(w, http.StatusMethodNotAllowed, "error", "Method error")
		return
	}
	var data models.PostPageData
	user_id := r.Context().Value(keyUser)

	if user_id != nil {
		user, err := h.srv.GetUserByID(user_id.(int))
		if err == nil {
			data.User = &user
		}
	}
	queryParams, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, errorTemp, "")
		return
	}

	post_id_str := queryParams.Get("id")
	post_id, err := strconv.Atoi(post_id_str)
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusBadRequest, errorTemp, "")
		return
	}

	post, err := h.srv.GetPost(post_id)
	if err != nil {
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}
	data.Post = post
	data.Categories, _ = h.srv.GetCategories()
	comments, err := h.srv.GetPostComments(post_id)
	if err != nil {
		h.errorMsg(w, http.StatusBadRequest, "error", err.Error())
		return
	}
	data.Comments = comments
	if err := templates["post"].Execute(w, data); err != nil {
		h.errorMsg(w, http.StatusInternalServerError, "error", err.Error())
		return
	}
}
