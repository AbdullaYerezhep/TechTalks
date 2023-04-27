package controller

import "net/http"

type CustomError struct {
	Status     int
	StatusText string
	CustomText string
}

func (h *Handler) errorMsg(w http.ResponseWriter, status int, tmpl, msg string) {
	w.WriteHeader(status)
	e := CustomError{
		Status:     status,
		StatusText: http.StatusText(status),
		CustomText: msg,
	}
	templates[tmpl].Execute(w, e)
}
