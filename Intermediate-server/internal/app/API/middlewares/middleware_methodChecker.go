package middlewares

import (
	"fmt"
	"net/http"
)

type MethodCheckerMiddleware struct {
	Middleware
}

func NewMethodCheckerMiddleware() *MethodCheckerMiddleware {
	return &MethodCheckerMiddleware{}
}

func (s *MethodCheckerMiddleware) CheckMethod(request *http.Request, needMethod string) error {
	if request.Method != needMethod {
		return fmt.Errorf("Method is not supported")
	}
	return nil
}
