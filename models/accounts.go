package models

import (
	"fmt"
	"time"

	"github.com/beevik/guid"
	"github.com/dgrijalva/jwt-go"
)

// Account is interface who implements account struct
type Account interface {
	ValidateGUID() error
	ValidateRefreshToken(string) error
	CreateTokens() (string, string, error)
}

type token struct {
	GUID string
	jwt.StandardClaims
}

type account struct {
	GUID  string `json:"guid" bson:"guid"`
	Token string `json:"token" bson:"token"`
}

// NewAccount is a function who return new account struct
func NewAccount() account {
	return account{}
}

// Valid is a method by Account who decode html body to account, checked GUID format and return error if exist
func (account *account) ValidateGUID() (err error) {
	g := guid.New()
	g, err = guid.ParseString(account.GUID)
	if err != nil {
		return
	}
	account.GUID = g.String() // always to lowercase
	return
}

// ValidateRefreshToken is method who validates the incoming token according to the Secret phrase
func (account *account) ValidateRefreshToken(incomingToken string) (err error) {
	var parsedToken *jwt.Token
	tk := &token{}
	parsedToken, err = jwt.ParseWithClaims(incomingToken, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return
	}

	if !parsedToken.Valid {
		err = fmt.Errorf("token is not valid")
	}
	account.GUID = tk.Id
	return
}

// CreateTokens is a method who prepares and return access and refresh tokens
func (account *account) CreateTokens() (accessToken string, refreshToken string, err error) {
	claims := &jwt.StandardClaims{
		Subject:   "account_token",
		Id:        account.GUID,
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessToken, err = at.SignedString([]byte(Secret))
	if err != nil {
		return
	}

	claims.Subject = "refresh_token"
	claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	refreshToken, err = rt.SignedString([]byte(Secret))
	if err != nil {
		accessToken = ""
		return
	}
	account.Token = refreshToken
	return
}
