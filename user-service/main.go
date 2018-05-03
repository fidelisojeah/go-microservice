// user-service/main.go

package main

import (
	"fmt"
	"log"
	"os"

	userProto "github.com/fidelisojeah/go-microservice/user-service/proto/auth"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/mdns"
	k8s "github.com/micro/kubernetes/go/micro"
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

	var srv micro.Service
	// Create a new service. Optionally include some options here.
	if os.Getenv("DEV") == "true" {
		srv = micro.NewService(
			// This name must match the package name given in your protobuf definition
			micro.Name("auth"),
		)
	} else {

		srv = k8s.NewService(
			// This name must match the package name given in your protobuf definition
			micro.Name("microservice.auth"),
		)
	}
	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	userProto.RegisterAuthHandler(srv.Server(), &service{repo, tokenService})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
