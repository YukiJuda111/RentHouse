// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user.proto

package user

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
	SendSms(ctx context.Context, in *SmsRequest, opts ...client.CallOption) (*SmsResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
	GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
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

func (c *userService) SendSms(ctx context.Context, in *SmsRequest, opts ...client.CallOption) (*SmsResponse, error) {
	req := c.c.NewRequest(c.name, "User.SendSms", in)
	out := new(SmsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetUserInfo", in)
	out := new(UserInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	SendSms(context.Context, *SmsRequest, *SmsResponse) error
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
	GetUserInfo(context.Context, *UserInfoRequest, *UserInfoResponse) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		SendSms(ctx context.Context, in *SmsRequest, out *SmsResponse) error
		Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error
		GetUserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
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

func (h *userHandler) SendSms(ctx context.Context, in *SmsRequest, out *SmsResponse) error {
	return h.UserHandler.SendSms(ctx, in, out)
}

func (h *userHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.UserHandler.Register(ctx, in, out)
}

func (h *userHandler) GetUserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error {
	return h.UserHandler.GetUserInfo(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserHandler.Login(ctx, in, out)
}
