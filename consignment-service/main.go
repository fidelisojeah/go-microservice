// consignment-service/main.go
package main

import (
	"fmt"
	"log"

	// Import the generated protobuf code
	pb "github.com/fidelisojeah/go-microservice/consignment-service/proto/consignment"
	vesselProto "github.com/fidelisojeah/go-microservice/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// Repository class
type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// ConsignmentRepository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

// Create method from repo.Create
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// GetAll returns an array of consignments
func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo         Repository
	vesselClient vesselProto.vesselServiceClient
}

// s.CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, response *pb.Response) error {

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}
	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id
	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	response.Created = true
	response.Consignment = consignment
	return nil
}

// s.GetCosignments
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, response *pb.Response) error {
	consignments := s.repo.GetAll()
	response.Consignments = consignments
	return nil
}
func main() {

	repo := &ConsignmentRepository{}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
