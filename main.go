package main

import (
	"net/http"

	"authservice/controllers"
	"authservice/models"
)

func main() {
	const DOMAIN = "localhost"
	const SECRETKEY = "So$0meP3r[hektK&y"
	const URLDB = "localhost:27017"

	models.Init(URLDB, SECRETKEY)

	session := models.NewSession(models.NewAccount())
	controller := controllers.NewAuthController(&session.Account, session, DOMAIN)
	http.HandleFunc("/auth/login", controller.Authenticate)
	http.HandleFunc("/auth/refresh", controller.Refresh)

	http.ListenAndServe(DOMAIN+":8080", nil)
}
