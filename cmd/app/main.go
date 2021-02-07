package main

import (
	"context"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/Messager"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	v1 "github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Survey documentation API
// @version 1.0
// @description This is an api for creating a app

// @contact.name Agostin Jean-baptiste
// @contact.email Jbagostin@gmail.com

// @license.name MIT

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	gin.ForceConsoleColor()
	router := gin.Default()

	if err := database.Connect(); err != nil {
		log.Panic(err)
	}

	if err := Messager.ConnectToRabbitMQ(); err != nil {
		log.Panic(err)
	}

	domain.InitBusses()

	err := elasticsearch.MakeConnection()
	if err != nil {
		log.Println(err)
		return
	}

	err = elasticsearch.SetUpIndexes()
	if err != nil {
		log.Println(err)
		return
	}

	defer elasticsearch.Close()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1.ApplyRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
