package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ved2pj/Duanlink/config"
	"github.com/ved2pj/Duanlink/internal/datastore"
	"github.com/ved2pj/Duanlink/internal/handlers"
	"github.com/ved2pj/Duanlink/internal/repos"
	"github.com/ved2pj/Duanlink/internal/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if err = datastore.NewDatastore(cfg); err != nil {
		log.Fatal(err)
	}

	shortLinkRepo := repos.NewShortLinkRepo(datastore.Get().MySQL)
	shortLinkService := services.NewShortLinkService(shortLinkRepo)
	shortLinkHandler := handlers.NewShortLinkHandler(shortLinkService)

	router := gin.Default()
	router.POST("/shortlinks", shortLinkHandler.Create)
	router.GET("/shortlinks/:short_code", shortLinkHandler.Lookup)
	router.GET("/:short_code", shortLinkHandler.Redirect)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.API.Host, cfg.API.Port),
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
