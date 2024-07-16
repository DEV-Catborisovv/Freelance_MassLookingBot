package configs

import "time"

type Config struct {
	HTTPServer_Port string
	Postgres        PostgreSql
	GRPC            GRPCConfig
}

type PostgreSql struct {
	POSTGRESQL_HOST string
	POSTGRESQL_PORT string
	POSTGRESQL_USER string
	POSTGRESQL_PASS string
	POSTGRESQL_DB   string
}

type GRPCConfig struct {
	Port    int
	Timeout time.Duration
}

func NewConfig() *Config {
	return &Config{
		HTTPServer_Port: httpServerPort,
		Postgres: PostgreSql{
			POSTGRESQL_HOST: postgresqlHost,
			POSTGRESQL_PORT: posegresqlPort,
			POSTGRESQL_USER: postgresqlUser,
			POSTGRESQL_PASS: postgresqlPass,
			POSTGRESQL_DB:   postgresqlDb,
		},
		GRPC: GRPCConfig{
			Port:    grpcPort,
			Timeout: grpcTimeout,
		},
	}
}
