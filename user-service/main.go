// user-service/main.go

package main

import (
	"fmt"
	"log"

	userProto "github.com/fidelisojeah/go-microservice/user-service/proto/auth"
	"github.com/micro/go-micro"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&userProto.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("auth"),
		micro.Version("latest"),
	)
	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	userProto.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
