// user-service/token_service.go
package main

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	userProto "github.com/fidelisojeah/go-microservice/user-service/proto/auth"
)

var (

	// Define a secure key string used
	// as a salt when hashing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something.
	key = []byte(os.Getenv("JSONSecret"))
)

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *userProto.User
	jwt.StandardClaims
}

// Authable methods
type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *userProto.User) (string, error)
}

// TokenService - class
type TokenService struct {
	repo Repository
}

// Decode - Decode jwt token
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	// Validate the token and return the custom claims
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, err

}

// Encode - Encode jwt token
func (srv *TokenService) Encode(user *userProto.User) (string, error) {

	expireToken := time.Now().Add(time.Hour * 720).Unix()
	// remove password from token
	user.Password = ""
	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "auth",
		},
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)

}
