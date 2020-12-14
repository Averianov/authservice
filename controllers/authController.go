package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Averianov/authservice/models"
	u "github.com/Averianov/authservice/utils"
)

// AuthController struct has two model interfaces - Account and Token
type AuthController struct {
	account models.Account
	session models.Session
	domain  string
}

// NewAuthController is a function who return new AuthController struct (account models.Account, token models.Token)
func NewAuthController(account models.Account, session models.Session, domain string) *AuthController {
	return &AuthController{account: account, session: session, domain: domain}
}

// Authenticate is a Handler by AuthController who check auth data and return two tokens
func (controller *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
	var err error

	err = json.NewDecoder(r.Body).Decode(controller.account)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	err = controller.account.ValidateGUID()
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	controller.prepareResponse(w)
}

// Refresh is a Handler by AuthController who refresh access and refresh tokens
func (controller *AuthController) Refresh(w http.ResponseWriter, r *http.Request) {
	var err error

	tokenHeader := r.Header.Get("Authorization")

	if tokenHeader == "" {
		return
	}
	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		return
	}
	incomingToken := splitted[1]

	err = controller.account.ValidateRefreshToken(incomingToken)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	err = controller.session.CompareWithExisting(incomingToken)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	controller.prepareResponse(w)
}

func (controller *AuthController) prepareResponse(w http.ResponseWriter) {
	var err error
	var accessToken, refreshToken string

	accessToken, refreshToken, err = controller.account.CreateTokens()
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	err = controller.session.Save()
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	cookie := &http.Cookie{}
	cookie.Name = "refresh_Token"
	cookie.Value = refreshToken
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.Domain = controller.domain
	cookie.HttpOnly = true
	cookie.Path = "/auth"
	// cookie.Secure = true 									  // for tls connections (port: 443)
	cookie.SameSite = 1
	http.SetCookie(w, cookie)

	response := u.Message(true, "Tokens has been created")
	response["access_token"] = accessToken
	// response["refresh_token"] = refreshToken 				  // If not using cookie, then using refreshToken as JWT
	// response["X-XSS-Protection"] = 1                           // Turns on XSS filtering in browsers by default
	// response["X-Frame-Options"] = "SAMEORIGIN"                 // The page can only be displayed in a frame on the same origin as the page itself
	// response["Content-Security-Policy"] = "default-src 'self'" // Limit content sources to source server only
	u.Respond(w, response)
}
