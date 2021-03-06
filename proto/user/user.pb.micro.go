// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	//注册
	Register(ctx context.Context, in *UserRegisterRequest, opts ...client.CallOption) (*UserRegisterResponse, error)
	//登录
	Login(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error)
	//查询用户信息
	GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error)
	//通过id查询用户信息
	GetUserInfoById(ctx context.Context, in *UserId, opts ...client.CallOption) (*UserInfoResponse, error)
	// 更改用户信息
	UpdateUserInfo(ctx context.Context, in *UserRegisterRequest, opts ...client.CallOption) (*UserInfoResponse, error)
	// 删除用户
	DeleteUser(ctx context.Context, in *UserId, opts ...client.CallOption) (*UserRegisterResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Register(ctx context.Context, in *UserRegisterRequest, opts ...client.CallOption) (*UserRegisterResponse, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(UserRegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(UserLoginResponse)
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

func (c *userService) GetUserInfoById(ctx context.Context, in *UserId, opts ...client.CallOption) (*UserInfoResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetUserInfoById", in)
	out := new(UserInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUserInfo(ctx context.Context, in *UserRegisterRequest, opts ...client.CallOption) (*UserInfoResponse, error) {
	req := c.c.NewRequest(c.name, "User.UpdateUserInfo", in)
	out := new(UserInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteUser(ctx context.Context, in *UserId, opts ...client.CallOption) (*UserRegisterResponse, error) {
	req := c.c.NewRequest(c.name, "User.DeleteUser", in)
	out := new(UserRegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	//注册
	Register(context.Context, *UserRegisterRequest, *UserRegisterResponse) error
	//登录
	Login(context.Context, *UserLoginRequest, *UserLoginResponse) error
	//查询用户信息
	GetUserInfo(context.Context, *UserInfoRequest, *UserInfoResponse) error
	//通过id查询用户信息
	GetUserInfoById(context.Context, *UserId, *UserInfoResponse) error
	// 更改用户信息
	UpdateUserInfo(context.Context, *UserRegisterRequest, *UserInfoResponse) error
	// 删除用户
	DeleteUser(context.Context, *UserId, *UserRegisterResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Register(ctx context.Context, in *UserRegisterRequest, out *UserRegisterResponse) error
		Login(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error
		GetUserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error
		GetUserInfoById(ctx context.Context, in *UserId, out *UserInfoResponse) error
		UpdateUserInfo(ctx context.Context, in *UserRegisterRequest, out *UserInfoResponse) error
		DeleteUser(ctx context.Context, in *UserId, out *UserRegisterResponse) error
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

func (h *userHandler) Register(ctx context.Context, in *UserRegisterRequest, out *UserRegisterResponse) error {
	return h.UserHandler.Register(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) GetUserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error {
	return h.UserHandler.GetUserInfo(ctx, in, out)
}

func (h *userHandler) GetUserInfoById(ctx context.Context, in *UserId, out *UserInfoResponse) error {
	return h.UserHandler.GetUserInfoById(ctx, in, out)
}

func (h *userHandler) UpdateUserInfo(ctx context.Context, in *UserRegisterRequest, out *UserInfoResponse) error {
	return h.UserHandler.UpdateUserInfo(ctx, in, out)
}

func (h *userHandler) DeleteUser(ctx context.Context, in *UserId, out *UserRegisterResponse) error {
	return h.UserHandler.DeleteUser(ctx, in, out)
}
