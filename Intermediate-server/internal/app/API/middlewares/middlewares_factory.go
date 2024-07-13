package middlewares

import (
	"fmt"
)

func GetMiddleware(middlewareType string) (IMiddleware, error) {
	switch middlewareType {
	default:
		return nil, fmt.Errorf("Wrong middleware type passed")
	case MiddlewareLogger:
		return NewLoggerMiddleware(), nil
	case MiddlewareMethodChecker:
		return NewMethodCheckerMiddleware(), nil
	case MiddlewareErrorWriter:
		return NewErrorWriterMiddleware(), nil
	}
}
