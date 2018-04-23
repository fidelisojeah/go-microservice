// user-service/handler.go

package main

import (
	"errors"
	"fmt"
	"log"

	userProto "github.com/fidelisojeah/go-microservice/user-service/proto/auth"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (srv *service) Create(ctx context.Context, req *userProto.User, response *userProto.Response) error {
	log.Println("Creating user: ", req)
	// Generates hash of password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}
	req.Password = string(hashedPass)

	createErr := srv.repo.Create(req)

	if createErr != nil {
		return fmt.Errorf("error creating User: %v", err)
	}
	response.User = req

	return nil
}
func (srv *service) Get(ctx context.Context, req *userProto.User, response *userProto.Response) error {

	user, err := srv.repo.Get(req.Id)

	if err != nil {
		return err
	}
	response.User = user

	return nil

}
func (srv *service) GetAll(ctx context.Context, req *userProto.Request, response *userProto.Response) error {

	users, err := srv.repo.GetAll()

	if err != nil {
		return err
	}
	response.Users = users

	return nil

}
func (srv *service) Auth(ctx context.Context, req *userProto.User, response *userProto.Token) error {
	log.Println("Logging in with:", req.Email)

	user, err := srv.repo.GetByEmail(req.Email)

	log.Println(user)
	if err != nil {
		return err
	}
	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}
	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	response.Token = token
	return nil
}
func (srv *service) ValidateToken(ctx context.Context, req *userProto.Token, res *userProto.Token) error {

	// Decode token
	claims, err := srv.tokenService.Decode(req.Token)

	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true

	return nil
}
