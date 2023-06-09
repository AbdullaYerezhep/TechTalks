package controller

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Status     int
	StatusText string
	CustomText string
}

func (h *Handler) errorMsg(w http.ResponseWriter, status int, tmpl, msg string) {
	e := CustomError{
		Status:     status,
		StatusText: http.StatusText(status),
		CustomText: msg,
	}
	w.WriteHeader(status)
	if err := templates["error"].Execute(w, e); err != nil {
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}
}
