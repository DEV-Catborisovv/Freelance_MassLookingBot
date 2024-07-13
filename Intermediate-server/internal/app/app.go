package app

import (
	api "Freelance_MassLookingBot_Intermediate-server/internal/app/API"
	"Freelance_MassLookingBot_Intermediate-server/internal/app/configs"
)

func InitApp() {
	config := configs.NewConfig()

	api_srv := api.NewApiServer(config.HTTPServer_Port)
	err := api_srv.Init()

	if err != nil {
		panic(err)
	}
}
