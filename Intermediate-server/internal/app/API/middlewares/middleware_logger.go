package middlewares

import (
	"log"
)

type LoggerMiddleWare struct {
	middleware Middleware
}

func NewLoggerMiddleware() *LoggerMiddleWare {
	return &LoggerMiddleWare{}
}

func (s *LoggerMiddleWare) Message(message string) {
	log.Println(message)
}
