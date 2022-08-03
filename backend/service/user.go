package service

import (
	"animeii.tech/dd"
	"animeii.tech/model"
	"animeii.tech/response"
)

type UserService struct{}

func (us *UserService) GetInfo(id uint) (res response.UserInfo, err error) {
	var user model.User
	err = dd.DB.First(&user, id).Scan(&res).Error
	return res, err
}
