// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package pbUser

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	UserRegister(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	UserSendEmail(ctx context.Context, in *UserMailRequest, opts ...client.CallOption) (*UserResponse, error)
	UserLogin(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) UserRegister(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "User.UserRegister", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserSendEmail(ctx context.Context, in *UserMailRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "User.UserSendEmail", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserLogin(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "User.UserLogin", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	UserRegister(context.Context, *UserRequest, *UserResponse) error
	UserSendEmail(context.Context, *UserMailRequest, *UserResponse) error
	UserLogin(context.Context, *UserRequest, *UserResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		UserRegister(ctx context.Context, in *UserRequest, out *UserResponse) error
		UserSendEmail(ctx context.Context, in *UserMailRequest, out *UserResponse) error
		UserLogin(ctx context.Context, in *UserRequest, out *UserResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) UserRegister(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserHandler.UserRegister(ctx, in, out)
}

func (h *userHandler) UserSendEmail(ctx context.Context, in *UserMailRequest, out *UserResponse) error {
	return h.UserHandler.UserSendEmail(ctx, in, out)
}

func (h *userHandler) UserLogin(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserHandler.UserLogin(ctx, in, out)
}