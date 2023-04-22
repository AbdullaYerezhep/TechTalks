package controller

import (
	"net/http"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	templates["home"].Execute(w, nil)
}
