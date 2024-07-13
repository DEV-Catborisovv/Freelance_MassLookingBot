package api

import (
	addtask "Freelance_MassLookingBot_Intermediate-server/internal/app/API/handlers/addTask"
	"net/http"
)

type ApiServer struct {
	port string
}

func NewApiServer(port string) *ApiServer {
	return &ApiServer{
		port: port,
	}
}

func (apisrv *ApiServer) Init() error {
	// Adding Handlers For Patterns
	http.HandleFunc("/api/add_task", addtask.HandleAddTask)

	// Initing Server
	err := http.ListenAndServe(apisrv.port, nil)
	if err != nil {
		return err
	}

	return nil
}
