// user-service/handler.go

package main

import (
	"log"

	userProto "github.com/fidelisojeah/go-microservice/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (srv *service) Create(ctx context.Context, req *userProto.User, response *userProto.Resposnse) error {

	// Generates hash of password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.password = hashedPass

	err := srv.repo.Create(req)

	if err != nil {
		return err
	}
	response.User = req
	return nil
}
func (srv *service) Get(ctx context.Context, req *userProto.User, response *userProto.Response) error {

	user, err := srv.repo.Get(req)

	if err != nil {
		return err
	}
	response.User = User

	return nil

}
func (srv *service) GetAll(ctx context.Context, req *userProto.Request, response *userProto.Response) error {

	users, err := srv.repo.GetAll(req)

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
}
func (srv *service) ValidateToken(ctx context.Context, req *userProto.Token, response *userProto.Token) error {

	return nil

}
