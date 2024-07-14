package storage

import (
	"Freelance_MassLookingBot_Intermediate-server/internal/app/models"
	"context"
	"sync"
	"time"

	"github.com/samber/lo"
)

type TasksPostgresStorage struct {
	Storage
}

// singleton variables and methods

var (
	instanceOfTasks     *TasksPostgresStorage
	onceInstanceOfTasks sync.Once
)

func getSingleTasksInstance() *TasksPostgresStorage {
	onceInstanceOfTasks.Do(func() {
		dbConn, err := getSingleDatabaseInstance()
		if err != nil {
			panic(err)
		}

		instanceOfTasks = &TasksPostgresStorage{}
		instanceOfTasks.db = dbConn
	})
	return instanceOfTasks
}

// methods of struct

func (s *TasksPostgresStorage) GetAll(ctx context.Context) ([]models.Task, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var tasks []dbTask
	if err := conn.SelectContext(ctx, tasks, `SELECT * FROM tasks;`); err != nil {
		return nil, err
	}

	return lo.Map(tasks, func(task dbTask, _ int) models.Task { return models.Task(task) }), nil
}

func (s *TasksPostgresStorage) GetById(ctx context.Context, id int) (*models.Task, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var task dbTask
	if err := conn.GetContext(ctx, task, `SELCT * FROM tasks WHERE id = $1`, id); err != nil {
		return nil, err
	}

	return (*models.Task)(&task), nil
}

func (s *TasksPostgresStorage) Add(ctx context.Context, task models.Task) (int, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var returnedId int
	row := conn.QueryRowxContext(ctx, "INSERT INTO tasks (status) VALUES ($1) RETURNING id", task.Status)

	if err := row.Err(); err != nil {
		return 0, err
	}

	if err := row.Scan(&returnedId); err != nil {
		return 0, err
	}

	return returnedId, nil
}

func (s *TasksPostgresStorage) UpdateStatus(ctx context.Context, id int, status string) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.QueryxContext(ctx, `UPDATE telegram_api_configs SET status = $1 WHERE id = $2`, status, id); err != nil {
		return err
	}
	return nil
}

type dbTask struct {
	ID        int       `db:"id"`
	Status    string    `db:"status"`
	CreaednAt time.Time `db:"creation_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
