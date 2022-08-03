package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"animeii.tech/dd"
	"animeii.tech/router"
	"github.com/gin-gonic/gin"
)

func main() {
	dd.Load()
	app := gin.Default()
	router.InitRouter(app)
	gin.SetMode(dd.Cf.App.Mode)
	// handleSignal(app)
	addr := fmt.Sprintf("%s:%s", dd.Cf.App.Host, dd.Cf.App.Port)
	// app.Run(addr)

	// go1.8版本关闭服务
	srv := &http.Server{
		Addr:    addr,
		Handler: app,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}

// func handleSignal(server *gin.Engine) {
// 	c := make(chan os.Signal)
// 	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGALRM)
// 	go func() {
// 		s := <-c
// 		log.Printf("got signal [%s],exiting now", s)
// 		os.Exit(0)
// 	}()
// }
