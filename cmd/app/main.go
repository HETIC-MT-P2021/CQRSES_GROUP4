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

	"github.com/gin-gonic/gin"
	_ "github.com/jibe0123/CQRSES_GROUP4/api"
	v1 "github.com/jibe0123/CQRSES_GROUP4/cmd/v1"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/database"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/database/elasticsearch"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	tryConnectToElasticSearch()

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

func tryConnectToElasticSearch() {
	for index := 0; index <= 5; index++ {
		es, err := elasticsearch.NewElastic(fmt.Sprintf("http://%s", "elasticsearch:9200"))
		if err != nil {
			log.Println(err)
			time.Sleep(3)
		} else {
			fmt.Println("You're connected to elastic search...")
			elasticsearch.SetRepository(es)
		}
	}
	defer elasticsearch.Close()
}
