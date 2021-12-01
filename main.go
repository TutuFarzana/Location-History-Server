package main

import (
	"flink/internal/app"
	"flink/internal/data/model"
	"os"
)

func main() {

	port := os.Getenv(model.PortEnv)
	if len(port) == 0 {
		port = model.PortEnvDefault
	}

	h := app.NewHandler(port)
	h.StartServer(h.InitRouter())
}
