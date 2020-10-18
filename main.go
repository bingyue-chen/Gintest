package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Gintest/src/middlewares"
	"Gintest/src/routes"
	"Gintest/src/utilities"
	"github.com/gin-gonic/gin"
)

// @title Gintest
// @version 0.1
// @description This is a gin practice project.

// @contact.name Bing Yue Chen
// @contact.url https://github.com/bingyue-chen
// @contact.email snow.shanalike@gmail.com

// @host gintest.snowcookie.moe
// @BasePath /api/v1

// @securityDefinitions.apikey JWTAuth
// @in header
// @name Authorization

func main() {

	router := gin.Default()

	router.Use(middlewares.ErrorsHandler(utilities.Getenv("GIN_DEBUG")))

	api_router := routes.ApiRouter{Router: router}

	api_router.Setup()

	//router.Run(":" + utilities.Getenv("GIN_PORT"))

	/*
		|--------------------------------------------------------------------------
		| Graceful shutdown or restart
		|--------------------------------------------------------------------------
		| Follow gin document:
		| https://github.com/gin-gonic/gin#graceful-shutdown-or-restart
	*/

	srv := &http.Server{
		Addr:    ":" + utilities.Getenv("GIN_PORT"),
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
