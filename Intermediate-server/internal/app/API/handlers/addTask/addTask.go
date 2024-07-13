package addtask

import (
	"Freelance_MassLookingBot_Intermediate-server/internal/app/API/middlewares"
	"Freelance_MassLookingBot_Intermediate-server/pkg/httperrors"
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

}
