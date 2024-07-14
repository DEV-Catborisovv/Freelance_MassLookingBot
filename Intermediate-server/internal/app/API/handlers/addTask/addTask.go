package addtask

import (
	"Freelance_MassLookingBot_Intermediate-server/internal/app/API/middlewares"
	"Freelance_MassLookingBot_Intermediate-server/internal/app/models"
	"Freelance_MassLookingBot_Intermediate-server/internal/app/storage"
	"Freelance_MassLookingBot_Intermediate-server/pkg/httperrors"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandleAddTask(w http.ResponseWriter, r *http.Request) {
	MethodChecker, err := middlewares.GetMiddleware(middlewares.MiddlewareMethodChecker)

	// Here i'm disable getting error because if i can't get middleware for error-writer i really have problems :)
	ErrorWriter, _ := middlewares.GetMiddleware(middlewares.MiddlewareErrorWriter)
	Logger, _ := middlewares.GetMiddleware(middlewares.MiddlewareLogger)

	// casting to types
	errorWriter, _ := ErrorWriter.(*middlewares.ErrorWriter)

	logger, ok := Logger.(*middlewares.LoggerMiddleWare)
	if !ok {
		errorWriter.WriteError(w, httperrors.InternalServerError)
		return
	}

	methodChecker, ok := MethodChecker.(*middlewares.MethodCheckerMiddleware)
	if !ok {
		logger.Message("Cant cast method checker interface to (MethodCheckerMiddleware) type")
		errorWriter.WriteError(w, httperrors.InternalServerError)
		return
	}

	// logic of method

	err = methodChecker.CheckMethod(r, http.MethodPut)
	if err != nil {
		errorWriter.WriteError(w, httperrors.MethodNotAllowed)
		return
	}

	// de-serialization

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorWriter.WriteError(w, httperrors.BadRequest)
		return
	}

	req := request{}
	if err := json.Unmarshal(body, &req); err != nil {
		errorWriter.WriteError(w, httperrors.BadRequest)
		return
	}

	// Work with data

	storageInterface, err := storage.GetStorage(storage.StorageTasks)
	if err != nil {
		logger.Message(fmt.Sprintf("Can't get storage interface from storage-fabric: %s", err))
		errorWriter.WriteError(w, httperrors.InternalServerError)
		return
	}

	tasksStorage, ok := storageInterface.(*storage.TasksPostgresStorage)
	if !ok {
		logger.Message(fmt.Sprintf("Can't cast task's storage interface to TasksPostgresStorage object: %s", err))
		errorWriter.WriteError(w, httperrors.InternalServerError)
		return
	}

	ctx := context.Background()
	modelTask := models.Task{}

	newTaskId, err := tasksStorage.Add(ctx, modelTask)
	if err != nil {
		logger.Message(fmt.Sprintf("Can't create new task: %s", err))
		errorWriter.WriteError(w, httperrors.InternalServerError)
		return
	}

	// Returning response

	resp := response{
		task_id: newTaskId,
	}
	respJson, err := json.Marshal(resp)

	if err != nil {
		logger.Message(fmt.Sprintf("Error with marshal response: %s", err))
		errorWriter.WriteError(w, httperrors.InternalServerError)
		return
	}

	w.Write(respJson)
}
