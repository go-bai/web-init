package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-bai/forward/config"
	"github.com/go-bai/forward/database"
	"github.com/go-bai/forward/task"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	database.InitEngine()
	defer database.Engine.Close()
	database.InitTable()

	task.T.Start()

	route := gin.Default()
	setWeb(route)
	setRoute(route)

	srv := http.Server{
		Addr:    viper.GetString("app.http_addr"),
		Handler: route,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
