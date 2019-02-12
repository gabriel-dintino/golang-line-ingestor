package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	configurations "github.com/golang-line-ingestor/configurations"
	file "github.com/golang-line-ingestor/utilities"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, fmt.Sprintf("Pong!"))
}

func IngestFile(c *gin.Context) {
	fmt.Println("Doing...")
	if exists, err := file.Exists(configurations.LineFileFullpath); err == nil {
		if exists {
			//procesar File
		}
	}
}
