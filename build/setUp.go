package build

import (
	"geometricSolver/config"
	apiLine "geometricSolver/internals/line/api"
	appLine "geometricSolver/internals/line/application"
	apiMiddle "geometricSolver/internals/middleware/api"
	errPkg "geometricSolver/internals/myerror"
	apiPoint "geometricSolver/internals/point/api"
	appPoint "geometricSolver/internals/point/application"
	"github.com/spf13/viper"
)

const (
	ConfNameMain = "main"
	ConfType     = "yml"
	ConfPath     = "./config/source/"
)

type InstallSetUp struct {
	Point  apiPoint.PointApi
	Line   apiLine.LineApi
	Middle apiMiddle.MiddlewareApi
}

// инициализация всех структур для обработчиков
func SetUp(logger errPkg.MultiLoggerInterface) *InstallSetUp {
	pointApp := appPoint.PointApp{}
	checkErrorApiPoint := errPkg.CheckError{Logger: logger}
	pointApi := apiPoint.PointApi{Application: &pointApp, CheckErrors: &checkErrorApiPoint, Logger: logger}
	var _ apiPoint.PointApiInterface = &pointApi

	lineApp := appLine.LineApp{}
	checkErrorApiLine := errPkg.CheckError{Logger: logger}
	lineApi := apiLine.LineApi{Application: &lineApp, CheckErrors: &checkErrorApiLine, Logger: logger}
	var _ apiLine.LineApiInterface = &lineApi

	middleApi := apiMiddle.MiddlewareApi{Logger: logger}
	var _ apiMiddle.MiddlewareApiInterface = &middleApi

	var result InstallSetUp
	result.Point = pointApi
	result.Line = lineApi
	result.Middle = middleApi

	return &result
}

// инициализация конфигов
func InitConfig() (error, []interface{}) {
	viper.AddConfigPath(ConfPath)
	viper.SetConfigType(ConfType)

	viper.SetConfigName(ConfNameMain)
	errRead := viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.MyErrors{
			ProjectTypeText: errRead.Error(),
		}, nil
	}
	mainConfig := config.MainConfig{}
	errUnmarshal := viper.Unmarshal(&mainConfig)
	if errUnmarshal != nil {
		return &errPkg.MyErrors{
			ProjectTypeText: errUnmarshal.Error(),
		}, nil
	}

	var result []interface{}
	result = append(result, mainConfig)

	return nil, result
}
