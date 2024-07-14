package storage

import (
	"Freelance_MassLookingBot_Intermediate-server/internal/app/models"
	"context"
	"sync"

	"github.com/samber/lo"
)

type TelegramApiConfigsStorage struct {
	Storage
}

// singleton variables and methods

var (
	instanceOfTelegramApiConfigs *TelegramApiConfigsStorage
	onceTelegramApiConfigs       sync.Once
)

func getSingleTelegramApiConfigsInstance() *TelegramApiConfigsStorage {
	onceTelegramApiConfigs.Do(func() {
		dbConn, err := getSingleDatabaseInstance()
		if err != nil {
			panic(err)
		}

		instanceOfTelegramApiConfigs = &TelegramApiConfigsStorage{}
		instanceOfTelegramApiConfigs.db = dbConn
	})
	return instanceOfTelegramApiConfigs
}

// methods of struct

func (s *TelegramApiConfigsStorage) GetAll(ctx context.Context) ([]models.TelegramApiConfig, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var telegramApiConfig []dbTelegramApiConfig
	if err := conn.SelectContext(ctx, telegramApiConfig, `SELECT * FROM telegram_api_configs;`); err != nil {
		return nil, err
	}

	return lo.Map(telegramApiConfig, func(tgapiconfig dbTelegramApiConfig, _ int) models.TelegramApiConfig {
		return models.TelegramApiConfig(tgapiconfig)
	}), nil
}

func (s *TelegramApiConfigsStorage) GetById(ctx context.Context, id int) (*models.TelegramApiConfig, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var tgbotapiconfig dbTelegramApiConfig
	if err := conn.GetContext(ctx, tgbotapiconfig, `SELECT * FROM tasks WHERE id = $1`, id); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *TelegramApiConfigsStorage) Add(ctx context.Context, telegramApiConfig models.TelegramApiConfig) (int, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var returnedId int
	row := conn.QueryRowContext(ctx, "INSERT INTO telegram_api_configs (task_id, API_ID, API_HASH) VALUES ($1, $2, $3) RETURNING id", telegramApiConfig.TaskId, telegramApiConfig.API_ID, telegramApiConfig.API_HASH)

	if err := row.Err(); err != nil {
		return 0, err
	}

	if err := row.Scan(&returnedId); err != nil {
		return 0, err
	}

	return 0, nil
}

type dbTelegramApiConfig struct {
	ID       int    `db:"id"`
	TaskId   int    `db:"task_id"`
	API_ID   string `db:"API_ID"`
	API_HASH string `db:"API_HASH"`
}
