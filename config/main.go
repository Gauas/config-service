package config

type AppConfig struct {
	Port      string
	DBUrl     string
	SecretKey string
}

func New() AppConfig {
	appConfig := load()
	validate(appConfig)
	return appConfig
}
