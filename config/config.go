package config

type MainConfig struct {
	Main Main `mapstructure:"main"`
}

type Main struct {
	PortHttp string
}
