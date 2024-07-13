package middlewares

import (
	"Freelance_MassLookingBot_Intermediate-server/pkg/httperrors"
	"net/http"
)

type ErrorWriter struct {
	Middleware
}

func NewErrorWriterMiddleware() *ErrorWriter {
	return &ErrorWriter{}
}

func (s *ErrorWriter) WriteError(writer http.ResponseWriter, httpError httperrors.HTTPError) {
	writer.WriteHeader(httpError.Code)
	writer.Write([]byte(httpError.Message))
}
