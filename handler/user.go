package handler

import (
	"context"
	"github.com/2559065/common"
	"github.com/2559065/user/domain/model"
	"github.com/2559065/user/domain/service"
	user "github.com/2559065/user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

//注册
func (u *User) Register(ctx context.Context, userRegisterRequest *user.UserRegisterRequest, userRegisterResponse *user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     userRegisterRequest.UserName,
		FirstName:    userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	userRegisterResponse.Message = "添加成功"
	return nil
}

//登录
func (u *User) Login(ctx context.Context, userLogin *user.UserLoginRequest, loginResponse *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(userLogin.UserName, userLogin.Pwd)
	if err != nil {
		return err
	}
	loginResponse.IsSuccess = isOk
	return nil
}

//查询用户信息
func (u *User) GetUserInfo(ctx context.Context, userInfoRequest *user.UserInfoRequest, userInfoResponse *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return err
	}
	userInfoResponse.UserName = userInfo.UserName
	userInfoResponse.FirstName = userInfo.FirstName
	userInfoResponse.UserId = userInfo.ID
	return nil
}

// 更改用户信息
func (u *User) UpdateUserInfo(ctx context.Context, req *user.UserRegisterRequest, res *user.UserInfoResponse) error {
	user := &model.User{}
	common.SwapTo(req, user)
	err := u.UserDataService.UpdateUser(user, req.Pwd != "")
	if err != nil {
		return err
	}
	res.UserName = user.UserName
	res.UserId = user.ID
	res.FirstName = user.FirstName
	return nil
}

// 删除用户
func (u *User) DeleteUser(ctx context.Context, req *user.UserId, res *user.UserRegisterResponse) error {
	err := u.UserDataService.DeleteUser(req.Id)
	if err != nil {
		return err
	}
	res.Message = "删除成功"
	return nil
}
