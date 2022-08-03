package middleware

import (
	"strconv"
	"time"

	"animeii.tech/dd"
	"animeii.tech/response"
	"animeii.tech/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "null" {
			response.TokenFail(c)
			c.Abort()
			return
		}
		// 截取"bearer "
		tokenStr = tokenStr[len(service.TokenType)+1:]

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &service.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(dd.Cf.Jwt.Secret), nil
		})

		if err != nil {
			response.TokenFail(c)
			c.Abort()
			return
		}

		if err != nil || new(service.JwtService).IsInBlacklist(tokenStr) {
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*service.CustomClaims)
		// Token 发布者校验
		if claims.Issuer != GuardName {
			response.TokenFail(c)
			c.Abort()
			return
		}
		// log.Println(claims.Id)
		// if !new(services.).IsExpires(claims.Id) {
		// 	response.TokenFail(c)
		// 	c.Abort()
		// 	return
		// }
		claims_id, _ := strconv.Atoi(claims.Id)

		// token 续签
		if claims.ExpiresAt-time.Now().Unix() < dd.Cf.Jwt.RefreshGracePeriod {
			lock := dd.Lock("refresh_token_lock", dd.Cf.Jwt.JwtBlacklistGracePeriod)
			if lock.Get() {
				user, err := new(service.JwtService).GetUserInfo(GuardName, uint(claims_id))
				if err != nil {
					dd.Log.Error(err.Error())
					lock.Release()
				} else {
					tokenData, _, _ := new(service.JwtService).CreateToken(GuardName, user)
					c.Header("new-token", tokenData.AccessToken)
					c.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
					_ = new(service.JwtService).JoinBlackList(token)
				}
			}
		}

		// 放在context中，方便登录后知道谁访问服务
		c.Set("token", token)
		c.Set("id", claims.Id)

	}
}
