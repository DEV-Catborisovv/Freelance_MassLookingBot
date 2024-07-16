package addtask

import (
	"Freelance_MassLookingBot_Intermediate-server/internal/app/API/middlewares"
	"Freelance_MassLookingBot_Intermediate-server/internal/app/models"
	"Freelance_MassLookingBot_Intermediate-server/internal/app/pyrunner"
	"Freelance_MassLookingBot_Intermediate-server/internal/app/storage"
	httpstatuses "Freelance_MassLookingBot_Intermediate-server/pkg/httpStatuses"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandleAddTask(w http.ResponseWriter, r *http.Request) {
	MethodChecker, err := middlewares.GetMiddleware(middlewares.MiddlewareMethodChecker)

	// Here i'm disable getting error because if i can't get middleware for error-writer i really have problems :)
	StatusWriter, _ := middlewares.GetMiddleware(middlewares.MiddlewareStatusWriter)
	Logger, _ := middlewares.GetMiddleware(middlewares.MiddlewareLogger)

	// casting to types
	statusWriter, _ := StatusWriter.(*middlewares.StatusWriter)

	logger, ok := Logger.(*middlewares.LoggerMiddleWare)
	if !ok {
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	methodChecker, ok := MethodChecker.(*middlewares.MethodCheckerMiddleware)
	if !ok {
		logger.Message("Cant cast method checker interface to (MethodCheckerMiddleware) type")
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	// logic of method

	err = methodChecker.CheckMethod(r, http.MethodPut)
	if err != nil {
		statusWriter.Write(w, httpstatuses.MethodNotAllowed)
		return
	}

	// de-serialization

	body, err := io.ReadAll(r.Body)
	if err != nil {
		statusWriter.Write(w, httpstatuses.BadRequest)
		return
	}

	req := request{}
	if err := json.Unmarshal(body, &req); err != nil {
		statusWriter.Write(w, httpstatuses.BadRequest)
		return
	}

	// Work with data

	storageInterface, err := storage.GetStorage(storage.StorageTasks)
	if err != nil {
		logger.Message(fmt.Sprintf("Can't get storage interface from storage-fabric: %s", err))
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	tasksStorage, ok := storageInterface.(*storage.TasksPostgresStorage)
	if !ok {
		logger.Message(fmt.Sprintf("Can't cast tasks storage interface to TasksPostgresStorage object: %s", err))
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	ctx := context.Background()
	newTaskId, err := tasksStorage.Add(ctx, models.Task{
		Status: "Created",
	})

	if err != nil {
		logger.Message(fmt.Sprintf("Can't create new task: %s", err))
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	// Creating and casting to telegramapiconfigs storage

	TelegramApiConfigsStorage, err := storage.GetStorage(storage.StorageTelegramApiConfigs)
	if err != nil {
		logger.Message(fmt.Sprintf("Can't get storage interface from storage-fabric: %s", err))
		statusWriter.Write(w, httpstatuses.InternalServerError)
	}

	telegramApiConfigsStorage, ok := TelegramApiConfigsStorage.(*storage.TelegramApiConfigsStorage)
	if !ok {
		logger.Message(fmt.Sprintf("Can't cast telegramApiConfigs storage interface to telegramApiConfigsStorage object: %s", err))
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	ctx = context.Background()
	_, err = telegramApiConfigsStorage.Add(ctx, models.TelegramApiConfig{
		TaskId:   newTaskId,
		API_ID:   req.API_ID,
		API_HASH: req.API_HASH,
	})

	if err != nil {
		logger.Message(fmt.Sprintf("Error with add telegramApiConfig to storage: %s", err))
		statusWriter.Write(w, httpstatuses.Processing)
		return
	}

	// Create task

	go func() {
		pyRunner := pyrunner.NewPyRunner(req.API_ID, req.API_HASH, req.PhoneNumber, req.Chats)
		pyRunner.Run()
	}()

	// Write resp
	statusWriter.Write(w, httpstatuses.Processing)
}
