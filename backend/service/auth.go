package service

import (
	"animeii.tech/dd"
	"animeii.tech/model"
	"animeii.tech/request"
	"animeii.tech/response"
	"animeii.tech/utils"
)

type AuthService struct{}

func (as *AuthService) Register(user *model.User) (res response.UserInfo, err error) {
	err = dd.DB.Create(user).Scan(&res).Error
	return
}

// 账号密码是否正确
func (as *AuthService) CheckAccount(req *request.LoginForm) (res response.UserInfo, err error) {
	var user model.User
	err = dd.DB.First(&user, "account = ? AND password = ?", req.Account, utils.Sha256(req.Password)).Scan(&res).Error
	return
}

// auth login
func (as *AuthService) Repss(form *request.RepassForm) (res response.UserInfo, err error) {
	var user model.User
	err = dd.DB.Model(&user).Where("email = ?", form.Email).Update("password", utils.Sha256(form.Password)).Scan(&res).Error
	return
}

// mark
func (as *AuthService) IsRegisterAccount(account string) (is bool, err error) {
	var res response.Id
	err = dd.DB.Select("id").Where("account = ?", account).First(&model.User{}).Scan(&res).Error
	if res.Id > 0 {
		is = true
	}
	return
}
func (as *AuthService) IsRegisterEmail(email string) (is bool, err error) {
	var res response.Id
	err = dd.DB.Select("id").Where("email = ?", email).First(&model.User{}).Scan(&res).Error
	if res.Id > 0 {
		is = true
	}
	return
}
