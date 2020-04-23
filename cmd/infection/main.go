// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This package is the primary infected keys upload service.
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cambio/pkg/android"
	"cambio/pkg/api"
	"cambio/pkg/database"
	"cambio/pkg/logging"

	"github.com/gorilla/mux"
)

const (
	portEnvVar  = "PORT"
	defaultPort = "8080"
)

func main() {
	ctx := context.Background()
	logger := logging.FromContext(ctx)

	port := os.Getenv(portEnvVar)
	if port == "" {
		port = defaultPort
	}
	logger.Infof("Using port %s (override with $%s)", port, portEnvVar)

	if err := database.Initialize(); err != nil {
		logger.Fatalf("unable to connect to database: %v", err)
	}
	if err := android.InitializeSafetynet(); err != nil {
		logger.Fatalf("android.InitializeSafetynet: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", api.HandlePublish())
	logger.Info("starting infection server")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
