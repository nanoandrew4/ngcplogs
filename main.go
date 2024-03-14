// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package main

import (
	"github.com/docker/go-plugins-helpers/sdk"
	"log"
	"os"
	"strconv"
)

//const socketAddress = "/run/docker/plugins/ngcplogs.sock"

var (
	localLoggingEnabled = false
)

func main() {
	nGCPDriver := createDriver()

	parsedLocalLogging, err := strconv.ParseBool(os.Getenv("local-logging"))
	if err == nil {
		localLoggingEnabled = parsedLocalLogging
	}

	sdkHandler := sdk.NewHandler(`{"Implements": ["LoggingDriver"]}`)
	registerHandlers(&sdkHandler, nGCPDriver)
	err = sdkHandler.ServeUnix("ngcplogs", 0)
	if err != nil {
		log.Fatalf("Error in socket handler: %s", err)
	}
}
