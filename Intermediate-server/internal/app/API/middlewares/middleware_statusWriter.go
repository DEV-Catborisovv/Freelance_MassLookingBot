package middlewares

import (
	httpstatuses "Freelance_MassLookingBot_Intermediate-server/pkg/httpStatuses"
	"net/http"
)

type StatusWriter struct {
	Middleware
}

func NewStatusWriterMiddleware() *StatusWriter {
	return &StatusWriter{}
}

func (s *StatusWriter) Write(writer http.ResponseWriter, httpStatus httpstatuses.HTTPStatus) {
	writer.WriteHeader(httpStatus.Code)
	writer.Write([]byte(httpStatus.Message))
}
