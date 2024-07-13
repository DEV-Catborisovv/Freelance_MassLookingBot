package configs

type Config struct {
	HTTPServer_Port string
}

func NewConfig() *Config {
	return &Config{
		HTTPServer_Port: httpServerPort,
	}
}
