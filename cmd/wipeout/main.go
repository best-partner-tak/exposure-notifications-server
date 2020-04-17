package main

import (
	"context"
	"log"
	"net/http"

	"cambio/pkg/api"
	"cambio/pkg/database"
	"cambio/pkg/logging"

	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	logger := logging.FromContext(ctx)

	if err := database.Initialize(); err != nil {
		logger.Fatalf("unable to connect to database: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", api.HandleWipeout())
	logger.Info("starting wipeout server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
