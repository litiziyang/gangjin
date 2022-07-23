// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: hot.proto

package hot

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for HotService service

func NewHotServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for HotService service

type HotService interface {
	GetSecurityHot(ctx context.Context, in *HotRequest, opts ...client.CallOption) (*HotResponse, error)
}

type hotService struct {
	c    client.Client
	name string
}

func NewHotService(name string, c client.Client) HotService {
	return &hotService{
		c:    c,
		name: name,
	}
}

func (c *hotService) GetSecurityHot(ctx context.Context, in *HotRequest, opts ...client.CallOption) (*HotResponse, error) {
	req := c.c.NewRequest(c.name, "HotService.GetSecurityHot", in)
	out := new(HotResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HotService service

type HotServiceHandler interface {
	GetSecurityHot(context.Context, *HotRequest, *HotResponse) error
}

func RegisterHotServiceHandler(s server.Server, hdlr HotServiceHandler, opts ...server.HandlerOption) error {
	type hotService interface {
		GetSecurityHot(ctx context.Context, in *HotRequest, out *HotResponse) error
	}
	type HotService struct {
		hotService
	}
	h := &hotServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&HotService{h}, opts...))
}

type hotServiceHandler struct {
	HotServiceHandler
}

func (h *hotServiceHandler) GetSecurityHot(ctx context.Context, in *HotRequest, out *HotResponse) error {
	return h.HotServiceHandler.GetSecurityHot(ctx, in, out)
}
