// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

/*
Package go_micro_srv_vessel is a generated protocol buffer package.

It is generated from these files:
	proto/vessel/vessel.proto

It has these top-level messages:
	Vessel
	Specification
	Response
*/
package go_micro_srv_vessel

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

type Vessel struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Capacity  int32  `protobuf:"varint,2,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight int32  `protobuf:"varint,3,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Available bool   `protobuf:"varint,5,opt,name=available" json:"available,omitempty"`
	OwnerId   string `protobuf:"bytes,6,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
}

func (m *Vessel) Reset()                    { *m = Vessel{} }
func (m *Vessel) String() string            { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()               {}
func (*Vessel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity  int32 `protobuf:"varint,1,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight int32 `protobuf:"varint,2,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
}

func (m *Specification) Reset()                    { *m = Specification{} }
func (m *Specification) String() string            { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()               {}
func (*Specification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Created bool      `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Vessel  *Vessel   `protobuf:"bytes,1,opt,name=vessel" json:"vessel,omitempty"`
	Vessels []*Vessel `protobuf:"bytes,2,rep,name=vessels" json:"vessels,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "go.micro.srv.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "go.micro.srv.vessel.Specification")
	proto.RegisterType((*Response)(nil), "go.micro.srv.vessel.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x75, 0xd3, 0x36, 0x4d, 0x47, 0xe2, 0x61, 0xbd, 0x6c, 0xab, 0x85, 0x90, 0x53, 0x4e, 0x2b,
	0xb4, 0xf8, 0x01, 0x5e, 0x04, 0x3d, 0xa6, 0xa2, 0xc7, 0xb2, 0xdd, 0x8c, 0x75, 0x20, 0xc9, 0x86,
	0x6c, 0x48, 0xeb, 0xdf, 0xf8, 0xa9, 0xc2, 0xa6, 0xad, 0x54, 0x82, 0x9e, 0xf6, 0xcd, 0xcc, 0x9b,
	0xc7, 0x9b, 0xb7, 0x30, 0xad, 0x6a, 0xd3, 0x98, 0xbb, 0x16, 0xad, 0xc5, 0xfc, 0xf0, 0x48, 0xd7,
	0xe3, 0xd7, 0x5b, 0x23, 0x0b, 0xd2, 0xb5, 0x91, 0xb6, 0x6e, 0x65, 0x37, 0x8a, 0xbf, 0x18, 0xf8,
	0xaf, 0x0e, 0xf2, 0x2b, 0xf0, 0x28, 0x13, 0x2c, 0x62, 0xc9, 0x24, 0xf5, 0x28, 0xe3, 0x33, 0x08,
	0xb4, 0xaa, 0x94, 0xa6, 0xe6, 0x53, 0x78, 0x11, 0x4b, 0x46, 0xe9, 0xa9, 0xe6, 0x73, 0x80, 0x42,
	0xed, 0xd7, 0x3b, 0xa4, 0xed, 0x47, 0x23, 0x06, 0x6e, 0x3a, 0x29, 0xd4, 0xfe, 0xcd, 0x35, 0x38,
	0x87, 0x61, 0xa9, 0x0a, 0x14, 0x43, 0x27, 0xe6, 0x30, 0xbf, 0x85, 0x89, 0x6a, 0x15, 0xe5, 0x6a,
	0x93, 0xa3, 0x18, 0x45, 0x2c, 0x09, 0xd2, 0x9f, 0x06, 0x9f, 0x42, 0x60, 0x76, 0x25, 0xd6, 0x6b,
	0xca, 0x84, 0xef, 0xb6, 0xc6, 0xae, 0x7e, 0xca, 0xe2, 0x67, 0x08, 0x57, 0x15, 0x6a, 0x7a, 0x27,
	0xad, 0x1a, 0x32, 0xe5, 0x99, 0x31, 0xf6, 0xa7, 0x31, 0xef, 0x97, 0xb1, 0xb8, 0x85, 0x20, 0x45,
	0x5b, 0x99, 0xd2, 0x22, 0x5f, 0x82, 0xdf, 0x85, 0xe0, 0x44, 0x2e, 0x17, 0x37, 0xb2, 0x27, 0x20,
	0xd9, 0x85, 0x93, 0x1e, 0xa8, 0xfc, 0x1e, 0xc6, 0x1d, 0xb2, 0xc2, 0x8b, 0x06, 0xff, 0x6d, 0x1d,
	0xb9, 0x0b, 0x84, 0xb0, 0x83, 0x2b, 0xac, 0x5b, 0xd2, 0xc8, 0x5f, 0x20, 0x7c, 0xa4, 0x32, 0x7b,
	0x38, 0x05, 0x10, 0xf7, 0xea, 0x9c, 0x1d, 0x3e, 0x9b, 0xf7, 0x72, 0x8e, 0x07, 0xc5, 0x17, 0x1b,
	0xdf, 0xfd, 0xf4, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x7f, 0xd6, 0xf5, 0x06, 0x02, 0x00,
	0x00,
}
