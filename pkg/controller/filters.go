package controller

import (
	"forum/models"
	"net/http"
)

func (h *Handler) homeFilteredByLikes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorMsg(w, http.StatusMethodNotAllowed, "")
		return
	} else if r.URL.Path != "/popular" {
		h.errorMsg(w, http.StatusNotFound, "")
		return
	}

	var data models.HomePage
	id := r.Context().Value(keyUser)
	if id != nil {
		user, err := h.srv.GetUserByID(id.(int))
		if err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "")
			return
		}
		data.User = user
	}

	posts, err := h.srv.GetTopPostsByLikes()
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "")
		return
	}

	data.Posts = posts

	categories, err := h.srv.GetCategories()
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "")
		return
	}
	data.Categories = categories

	if err = templates["home"].Execute(w, data); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "")
		return
	}
}

func (h *Handler) MyPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorMsg(w, http.StatusMethodNotAllowed, "")
		return
	} else if r.URL.Path != "/myposts" {
		h.errorMsg(w, http.StatusNotFound, "")
		return
	}

	var data models.HomePage

	id, ok := r.Context().Value(keyUser).(int)
	if ok {
		user, err := h.srv.GetUserByID(id)
		if err != nil {
			h.errLog.Println(err.Error())
			h.errorMsg(w, http.StatusInternalServerError, "")
			return
		}
		data.User = user
	}

	posts, err := h.srv.GetMyPosts(id)
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "")
		return
	}

	data.Posts = posts

	categories, err := h.srv.GetCategories()
	if err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "")
		return
	}
	data.Categories = categories

	if err = templates["home"].Execute(w, data); err != nil {
		h.errLog.Println(err.Error())
		h.errorMsg(w, http.StatusInternalServerError, "")
		return
	}
}
