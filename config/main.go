package config

type Config struct {
	Port      string
	DBUrl     string
	SecretKey string
}

func New() Config {
	cfg := load()
	validate(cfg)
	return cfg
}
