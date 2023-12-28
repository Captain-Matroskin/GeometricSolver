package build

import (
	"geometricSolver/config"
	apiMiddle "geometricSolver/internals/middleware/api"
	errPkg "geometricSolver/internals/myerror"
	apiLine "geometricSolver/internals/solver/api"
	appLine "geometricSolver/internals/solver/application"
	"github.com/spf13/viper"
)

const (
	ConfNameMain = "main"
	ConfType     = "yml"
	ConfPath     = "./config/source/"
)

type InstallSetUp struct {
	Line   apiLine.SolverApi
	Middle apiMiddle.MiddlewareApi
}

// инициализация всех структур для обработчиков
func SetUp(logger errPkg.MultiLoggerInterface) *InstallSetUp {
	lineApp := appLine.SolverApp{}
	checkErrorApiLine := errPkg.CheckError{Logger: logger}
	lineApi := apiLine.SolverApi{Application: &lineApp, CheckErrors: &checkErrorApiLine, Logger: logger}
	var _ apiLine.SolverApiInterface = &lineApi

	middleApi := apiMiddle.MiddlewareApi{Logger: logger}
	var _ apiMiddle.MiddlewareApiInterface = &middleApi

	var result InstallSetUp
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
