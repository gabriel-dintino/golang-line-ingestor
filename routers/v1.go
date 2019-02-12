package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	handlers "github.com/golang-line-ingestor/handlers"
)

func GetRoutersV1(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/ping", handlers.Pong)
	v1.GET("/ingest", handlers.IngestFile)
}
