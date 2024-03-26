package main

import (
	"github.com/docker/docker/daemon/logger"
	"io"
)

type logPair struct {
	jsonl   logger.Logger
	gLogger logger.Logger
	logFile io.ReadCloser
	info    logger.Info
}

func (lp *logPair) Close() {
	lp.logFile.Close()
	lp.gLogger.Close()
	lp.jsonl.Close()
}
