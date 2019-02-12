package main

import (
	"runtime"

	"fmt"

	"github.com/gin-gonic/gin"
	configurations "github.com/golang-line-ingestor/configurations"
	routers "github.com/golang-line-ingestor/routers"
)

var router *gin.Engine

func main() {
	configureRouter()
	//mapUrlsToControllers()
	router.Run(fmt.Sprintf(":%s", configurations.Port))
}

func configureRouter() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	router = gin.Default()
	routers.GetRoutersV1(router)
	//pprof.Register(router, nil)
	//router.NoRoute(noRouteHandler)
	router.RedirectFixedPath = false
	router.RedirectTrailingSlash = false
}

//func noRouteHandler(c *gin.Context) {
//	mlhandlers.RespondError(c, apierrors.NewNotFoundApiError(fmt.Sprintf("Resource %s not found.", c.Request.URL.Path)))
//}
