package build

import (
	"geometricSolver/config"
	errPkg "geometricSolver/internals/myerror"
	"github.com/spf13/viper"
)

const (
	ConfNameMain = "main"
	ConfType     = "yml"
	ConfPath     = "./config/source/"
)

type InstallSetUp struct {
}

// инициализация всех структур для обработчиков
func SetUp() *InstallSetUp {

	var result InstallSetUp

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
