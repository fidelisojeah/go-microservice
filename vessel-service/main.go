// vessel-service/main.go
package main

import (
	"fmt"
	"log"
	"os"

	vesselProto "github.com/fidelisojeah/go-microservice/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
)

const (
	defaultHost = "localhost:27017"
)

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*vesselProto.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	defer session.Close()
	if err != nil {

		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	repo := &VesselRepository{session.Copy()}
	createDummyData(repo)

	srv := k8s.NewService(
		micro.Name("microservice.vessel"),
		micro.Version("latest"),
	)
	srv.Init()

	vesselProto.RegisterVesselServiceHandler(srv.Server(), &service{session})
	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
