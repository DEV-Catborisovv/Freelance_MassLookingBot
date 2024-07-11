package api

import (
	addtask "Freelance_MassLookingBot_Intermediate-server/internal/app/API/handlers/addTask"
	"net/http"
)

type ApiServer struct {
	PORT string
}

func (apisrv *ApiServer) Init() error {
	// Adding Handlers For Patterns
	http.HandleFunc("/api/add_task", addtask.HandleAddTask)

	// Initing Server
	err := http.ListenAndServe(apisrv.PORT, nil)
	if err != nil {
		return err
	}

	return nil
}
