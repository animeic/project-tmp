package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"animeii.tech/dd"
	"animeii.tech/response"
	"animeii.tech/utils"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenType    = "bearer"
	AppGuardName = "app"
)

type JwtService struct{}

type JwtUser interface {
	GetUid() string
}

type CustomClaims struct {
	jwt.StandardClaims
}

func (js *JwtService) CreateToken(GuardName string, user JwtUser) (tokenData response.Token, token *jwt.Token, err error) {
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + dd.Cf.Jwt.JwtTtl,
				Id:        user.GetUid(),
				Issuer:    GuardName, // 用于在中间件中区分不同客户端颁发的 token，避免 token 跨端使用
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)

	tokenStr, err := token.SignedString([]byte(dd.Cf.Jwt.Secret))
	tokenData = response.Token{
		AccessToken: tokenStr,
		ExpiresIn:   int(dd.Cf.Jwt.JwtTtl),
		TokenType:   TokenType,
	}
	return
}

// 获取黑名单
func (js *JwtService) getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + utils.MD5([]byte(tokenStr))
}

func (js *JwtService) JoinBlackList(token *jwt.Token) (err error) {
	nowUnix := time.Now().Unix()
	timer := time.Duration(token.Claims.(*CustomClaims).ExpiresAt-nowUnix) * time.Second
	err = dd.Redis.SetNX(context.Background(), js.getBlackListKey(token.Raw), nowUnix, timer).Err()
	return
}

func (jwtService *JwtService) IsInBlacklist(tokenStr string) bool {
	joinUnixStr, _ := dd.Redis.Get(context.Background(), jwtService.getBlackListKey(tokenStr)).Result()
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	// JwtBlacklistGracePeriod 为黑名单宽限时间，避免并发请求失效
	if time.Now().Unix()-joinUnix < dd.Cf.Jwt.JwtBlacklistGracePeriod {
		return false
	}
	return true

}

func (js *JwtService) GetUserInfo(GuardName string, id uint) (user JwtUser, err error) {
	switch GuardName {
	case AppGuardName:
		return new(UserService).GetInfo(id)
	default:
		err = errors.New("guard " + GuardName + " does not exist")
	}
	return
}
