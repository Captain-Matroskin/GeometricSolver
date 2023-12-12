package main

import (
	"geometricSolver/build"
	"geometricSolver/config"
	errPkg "geometricSolver/internals/myerror"
	"geometricSolver/internals/util"
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

	pointApi := startStructure.Point
	lineApi := startStructure.Line
	middlewareApi := startStructure.Middle

	myRouter := router.New()
	apiGroup := myRouter.Group("/api")
	versionGroup := apiGroup.Group("/v1")
	geomSolver := versionGroup.Group("/geomSolver")
	lineSolver := geomSolver.Group("/line")
	pointSolver := geomSolver.Group("/point")

	pointSolver.POST("/equal/", pointApi.EqualTwoPointsHandler)
	pointSolver.POST("/distance/", pointApi.DistanceBetweenTwoPointsHandler)
	pointSolver.POST("/fixation/", pointApi.FixationPointHandler)
	pointSolver.POST("/belongOfLine/", pointApi.BelongingPointOfLineHandler)
	lineSolver.POST("/parallelism/", lineApi.ParallelismTwoLinesHandler)
	lineSolver.POST("/perpendicular/", lineApi.PerpendicularTwoLinesHandler)
	lineSolver.POST("/corner/", lineApi.CornerTwoLinesHandler)
	lineSolver.POST("/vertical/", lineApi.VerticalLineHandler)
	lineSolver.POST("/horizontal/", lineApi.HorizontalLineHandler)

	addresHttp := ":" + configMain.Main.PortHttp

	logger.Log.Infof("Listen in 127:0.0.1%s", addresHttp)
	errStart := fasthttp.ListenAndServe(addresHttp, middlewareApi.LogURL(myRouter.Handler))
	if errStart != nil {
		logger.Log.Errorf("Listen and server http error: %v", errStart)
		os.Exit(3)
	}
}
