// vessel-service/handler.go
package main

import (
	"golang.org/x/net/context"

	vesselProto "github.com/fidelisojeah/go-microservice/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) FindAvailable(ctx context.Context, req *vesselProto.Specification, response *vesselProto.Response) error {
	defer s.GetRepo().Close()

	// Find the next available vessel
	vessel, err := s.GetRepo().FindAvailable(req)

	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	response.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *vesselProto.Vessel, response *vesselProto.Response) error {
	defer s.GetRepo().Close()

	// Find the next available vessel
	err := s.GetRepo().Create(req)

	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	response.Vessel = req
	response.Created = true
	return nil
}
