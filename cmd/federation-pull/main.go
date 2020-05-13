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

// This package is the service that pulls federation results from federation partners; it is intended to be invoked over HTTP by Cloud Scheduler.
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/google/exposure-notifications-server/internal/api/federation"
	"github.com/google/exposure-notifications-server/internal/logging"
	"github.com/google/exposure-notifications-server/internal/setup"
)

func main() {
	ctx := context.Background()
	logger := logging.FromContext(ctx)

	var config federation.PullConfig
	env, closer, err := setup.Setup(ctx, &config)
	if err != nil {
		logger.Fatal("setup.Setup: %v", err)
	}
	defer closer()

	http.Handle("/", federation.NewPullHandler(env.Database(), config))
	logger.Infof("starting federation puller on :%s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
