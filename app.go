package authservice

import (
	"net/http"
	"os"

	"github.com/Averianov/authservice/controllers"
	"github.com/Averianov/authservice/models"
	"github.com/joho/godotenv"
)

func Run() (err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}
	domain := os.Getenv("domain")
	tokenPassword := os.Getenv("token_password")
	urlDB := os.Getenv("url_db")

	err = models.InitDB(urlDB, tokenPassword)
	if err != nil {
		return
	}

	session := models.NewSession(models.NewAccount())
	controller := controllers.NewAuthController(&session.Account, session, domain)
	http.HandleFunc("/auth/login", controller.Authenticate)
	http.HandleFunc("/auth/refresh", controller.Refresh)

	http.ListenAndServe(domain+":8080", nil)
	return
}
