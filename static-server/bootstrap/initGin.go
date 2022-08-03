package bootstrap

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"i.animeii.tech/router"

	"github.com/gin-gonic/gin"
)

func InitGin() {
	app := gin.Default()
	router.InitRouter(app)
	gin.SetMode(gin.ReleaseMode)
	handleSignal(app)
	// https://av.animeii.tech/images/{imageurl}
	app.Run(fmt.Sprintf("%s:%s", "127.0.0.1", "3060"))
}

func handleSignal(server *gin.Engine) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGALRM)
	go func() {
		s := <-c
		log.Printf("got signal [%s],exiting now", s)
		os.Exit(0)
	}()
}
