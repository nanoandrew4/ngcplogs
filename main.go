// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package main

import (
	"github.com/docker/go-plugins-helpers/sdk"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("/tmp/ngcplogs.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	nGCPDriver := createDriver()

	sdkHandler := sdk.NewHandler(`{"Implements": ["LoggingDriver"]}`)
	registerHandlers(&sdkHandler, nGCPDriver)
	err = sdkHandler.ServeUnix("ngcplogs", 0)
	if err != nil {
		log.Fatalf("Error in socket handler: %s", err)
	}
}
