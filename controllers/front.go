package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	// These two routes are different as far as go is concerned
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
