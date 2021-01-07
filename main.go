package main

import (
	"authservice/controllers"
	"authservice/models"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() (err error) {

	fmt.Println("Read .env file...")
	err = godotenv.Load()
	if err != nil {
		return
	}
	domain := os.Getenv("domain")
	tokenPassword := os.Getenv("token_password")
	urlDB := os.Getenv("url_db")

	fmt.Println("Check connection with DB...")
	err = models.InitConnectionToDB(urlDB, tokenPassword)
	if err != nil {
		return
	}

	fmt.Println("Ready to work")
	session := models.NewSession(models.NewAccount())
	controller := controllers.NewAuthController(&session.Account, session, domain)
	http.HandleFunc("/auth/login", controller.Authenticate)
	http.HandleFunc("/auth/refresh", controller.Refresh)

	http.ListenAndServe(domain+":8080", nil)
	return
}
