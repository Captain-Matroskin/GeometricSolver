package build

import (
	"geometricSolver/config"
	apiGS "geometricSolver/internals/geomSolver/api"
	"geometricSolver/internals/geomSolver/application"
	apiMiddle "geometricSolver/internals/middleware/api"
	errPkg "geometricSolver/internals/myerror"
	"github.com/spf13/viper"
)

const (
	ConfNameMain = "main"
	ConfType     = "yml"
	ConfPath     = "./config/source/"
)

type InstallSetUp struct {
	GeomSolver apiGS.GeomSolverApi
	Middle     apiMiddle.MiddlewareApi
}

// инициализация всех структур для обработчиков
func SetUp(logger errPkg.MultiLoggerInterface) *InstallSetUp {
	geomSolverApp := application.GeomSolverApp{}
	checkErrorApiGS := errPkg.CheckError{Logger: logger}
	geomSolverApi := apiGS.GeomSolverApi{Application: geomSolverApp, CheckErrors: checkErrorApiGS, Logger: logger}
	var _ apiGS.GeomSolverApiInterface = &geomSolverApi

	middleApi := apiMiddle.MiddlewareApi{Logger: logger}
	var _ apiMiddle.MiddlewareApiInterface = &middleApi

	var result InstallSetUp
	result.GeomSolver = geomSolverApi
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
			Text: errRead.Error(),
		}, nil
	}
	mainConfig := config.MainConfig{}
	errUnmarshal := viper.Unmarshal(&mainConfig)
	if errUnmarshal != nil {
		return &errPkg.MyErrors{
			Text: errUnmarshal.Error(),
		}, nil
	}

	var result []interface{}
	result = append(result, mainConfig)

	return nil, result
}
