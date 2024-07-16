package addverificationcode

import (
	"Freelance_MassLookingBot_Intermediate-server/internal/app/API/middlewares"
	memorystorage "Freelance_MassLookingBot_Intermediate-server/internal/app/memoryStorage"
	httpstatuses "Freelance_MassLookingBot_Intermediate-server/pkg/httpStatuses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandleAddingVerificationCode(w http.ResponseWriter, r *http.Request) {
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

	memStorage := memorystorage.GetInstance()
	gettedChan, err := memStorage.GetChannel(req.Phone_number)
	if err != nil {
		logger.Message(fmt.Sprintf("Can't get chan from storage: %s\n", err))
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	go func() {
		gettedChan <- req.Code
	}()
}
