package storage

import (
	"Freelance_MassLookingBot_Intermediate-server/internal/app/configs"
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	singleDatabaseInstance *sqlx.DB
	onceDatabaseInstance   sync.Once
)

func getSingleDatabaseInstance() (*sqlx.DB, error) {
	var err error
	onceDatabaseInstance.Do(func() {
		config := configs.NewConfig()

		singleDatabaseInstance, err = sqlx.Connect("postgres",
			fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Postgres.POSTGRESQL_HOST, config.Postgres.POSTGRESQL_PORT, config.Postgres.POSTGRESQL_USER, config.Postgres.POSTGRESQL_PASS, config.Postgres.POSTGRESQL_DB))
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	})

	if singleDatabaseInstance == nil {
		return nil, fmt.Errorf("failed to create database instance")
	}

	return singleDatabaseInstance, nil
}
