package controller

import (
	"fmt"

	"animeii.tech/utils"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

var TC = new(TestController)

func (ts *TestController) IsRegisterByName(c *gin.Context) {

	fmt.Println(utils.Sha256("ssxhyw2515"))
}
