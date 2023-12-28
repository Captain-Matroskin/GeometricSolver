package main

import (
	"geometricSolver/build"
	"geometricSolver/config"
	errPkg "geometricSolver/internals/myerror"
	"geometricSolver/internals/util"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"os"
)

func main() {
	runServer()
}

func runServer() {
	var logger util.Logger
	logger.Log = util.NewLogger("./logs.txt")

	defer func(loggerErrWarn errPkg.MultiLoggerInterface) {
		errLogger := loggerErrWarn.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerErrWarn the buffer could not be cleared %v", errLogger)
			os.Exit(1)
		}
	}(logger.Log)
	//инициализация конфигов
	errConfig, configRes := build.InitConfig()
	if errConfig != nil {
		logger.Log.Errorf("%s", errConfig.Error())
		os.Exit(2)
	}
	configMain := configRes[0].(config.MainConfig)

	startStructure := build.SetUp(logger.Log)
	lineApi := startStructure.Line
	middlewareApi := startStructure.Middle

	myRouter := router.New()
	apiGroup := myRouter.Group("/api")
	versionGroup := apiGroup.Group("/v1")
	geomSolver := versionGroup.Group("/geomSolver")
	lineSolver := geomSolver.Group("/line")

	lineSolver.POST("/parallelism/", lineApi.GeomSolverHandler)

	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"access-control-allow-origin", "content-type",
			"x-csrf-token", "access-control-expose-headers"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT"},
		ExposedHeaders:   []string{"X-Csrf-Token"},
		AllowCredentials: true,
		AllowMaxAge:      5600,
		Debug:            true,
	})

	addresHttp := ":" + configMain.Main.PortHttp

	logger.Log.Infof("Listen in 127:0.0.1%s", addresHttp)
	errStart := fasthttp.ListenAndServe(addresHttp, withCors.CorsMiddleware(middlewareApi.LogURL(myRouter.Handler)))
	if errStart != nil {
		logger.Log.Errorf("Listen and server http error: %v", errStart)
		os.Exit(3)
	}
}
