// user-cli/cli.go
package main

import (
	"context"
	"log"
	"os"

	userProto "github.com/fidelisojeah/go-microservice/user-service/proto/user"
	micro "github.com/micro/go-micro"
)

func main() {

	srv := micro.NewService(

		micro.Name("go.micro.srv.user-cli"),
		micro.Version("latest"),
	)

	srv.Init()

	// Create new greeter client
	client := userProto.NewUserServiceClient("auth", microclient.DefaultClient)

	name := "Fidelis Ojeah"
	password := "Some password"
	email := "fidelis.ojeah@gmail.com"
	company := "Andela"

	// Call our user service
	r, err := client.Create(context.TODO(), &userProto.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &userProto.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	} else {
		log.Println("Users below=====>")
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}
	authResponse, err := client.Auth(context.TODO(), &userProto.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)
	os.Exit(0)

}
