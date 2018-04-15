// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

/*
Package go_micro_srv_consignment is a generated protocol buffer package.

It is generated from these files:
	proto/consignment/consignment.proto

It has these top-level messages:
	Consignment
	Container
	GetRequest
	Response
*/
package go_micro_srv_consignment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Message types
type Consignment struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Container   []*Container `protobuf:"bytes,4,rep,name=container" json:"container,omitempty"`
	VesselId    string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId" json:"vessel_id,omitempty"`
}

func (m *Consignment) Reset()                    { *m = Consignment{} }
func (m *Consignment) String() string            { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()               {}
func (*Consignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainer() []*Container {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin     string `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetRequest struct {
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Response struct {
	Created      bool           `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Consignment  *Consignment   `protobuf:"bytes,2,opt,name=consignment" json:"consignment,omitempty"`
	Consignments []*Consignment `protobuf:"bytes,3,rep,name=consignments" json:"consignments,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.consignment"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	CreateConsignment(context.Context, *Consignment, *Response) error
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ShippingServiceHandler.GetConsignments(ctx, in, out)
}

func init() { proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x46, 0x49, 0xff, 0x33, 0xa9, 0xa8, 0xf0, 0x02, 0x2c, 0x58, 0x10, 0xa5, 0x20, 0x75, 0x15,
	0xa4, 0x72, 0x02, 0xd4, 0x45, 0x95, 0xad, 0xbb, 0x46, 0x50, 0xe2, 0x51, 0x6a, 0x89, 0xd8, 0xc1,
	0x76, 0xcb, 0xd1, 0xe0, 0x1c, 0x9c, 0x08, 0xd5, 0x69, 0xa8, 0x11, 0x2a, 0xea, 0xce, 0xdf, 0x8c,
	0xdf, 0xe4, 0x65, 0x64, 0x18, 0x57, 0x5a, 0x59, 0x75, 0x97, 0x2b, 0x69, 0x44, 0x21, 0x4b, 0x94,
	0xd6, 0x3f, 0xa7, 0xae, 0x4b, 0x68, 0xa1, 0xd2, 0x52, 0xe4, 0x5a, 0xa5, 0x46, 0x6f, 0x52, 0xaf,
	0x9f, 0x7c, 0x04, 0x10, 0xcd, 0xf6, 0x99, 0x9c, 0x42, 0x4b, 0x70, 0x1a, 0xc4, 0xc1, 0x24, 0x64,
	0x2d, 0xc1, 0x49, 0x0c, 0x11, 0x47, 0x93, 0x6b, 0x51, 0x59, 0xa1, 0x24, 0x6d, 0xb9, 0x86, 0x5f,
	0x22, 0xe7, 0xd0, 0x7b, 0x47, 0x51, 0xac, 0x2c, 0x6d, 0xc7, 0xc1, 0xa4, 0xcb, 0x76, 0x89, 0x3c,
	0x40, 0x98, 0x2b, 0x69, 0x97, 0x42, 0xa2, 0xa6, 0x9d, 0xb8, 0x3d, 0x89, 0xa6, 0xe3, 0xf4, 0x90,
	0x47, 0x3a, 0x6b, 0xae, 0xb2, 0x3d, 0x45, 0xae, 0x20, 0xdc, 0xa0, 0x31, 0xf8, 0xfa, 0x24, 0x38,
	0xed, 0xba, 0x4f, 0x0f, 0xea, 0x42, 0xc6, 0x93, 0x12, 0xc2, 0x1f, 0xe8, 0x8f, 0xf6, 0x35, 0x44,
	0xf9, 0xda, 0x58, 0x55, 0xa2, 0xde, 0xb2, 0xb5, 0x36, 0x34, 0xa5, 0x8c, 0x6f, 0xad, 0x95, 0x16,
	0x85, 0x90, 0xce, 0x3a, 0x64, 0xbb, 0x44, 0x2e, 0xa0, 0xbf, 0x36, 0x35, 0xd4, 0xa9, 0x1b, 0xdb,
	0x98, 0xf1, 0x64, 0x08, 0x30, 0x47, 0xcb, 0xf0, 0x6d, 0x8d, 0xc6, 0x26, 0x9f, 0x01, 0x0c, 0x18,
	0x9a, 0x4a, 0x49, 0x83, 0x84, 0x42, 0x3f, 0xd7, 0xb8, 0xb4, 0x58, 0x1b, 0x0c, 0x58, 0x13, 0xc9,
	0x1c, 0x22, 0xef, 0x27, 0x9d, 0x46, 0x34, 0xbd, 0xfd, 0x77, 0x0b, 0xcd, 0x99, 0xf9, 0x24, 0xc9,
	0x60, 0xe8, 0x45, 0x43, 0xdb, 0x6e, 0x9f, 0x47, 0x4e, 0xfa, 0x85, 0x4e, 0xbf, 0x02, 0x18, 0x2d,
	0x56, 0xa2, 0xaa, 0x84, 0x2c, 0x16, 0xa8, 0x37, 0x22, 0x47, 0xf2, 0x0c, 0x67, 0x33, 0xa7, 0xec,
	0x3f, 0x85, 0xe3, 0xa6, 0x5f, 0x26, 0x87, 0xaf, 0x35, 0x1b, 0x4a, 0x4e, 0xc8, 0x23, 0x8c, 0xe6,
	0x68, 0x3d, 0xce, 0x90, 0x9b, 0xc3, 0xe0, 0x7e, 0xd3, 0xc7, 0x8d, 0x7f, 0xe9, 0xb9, 0x77, 0x7e,
	0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x3f, 0x33, 0xdf, 0x0e, 0x03, 0x00, 0x00,
}