package main

import (
	"fmt"
	"time"
)

type nGCPError struct {
	Msg  string `json:"msg"`
	File string `json:"file"`
	Line int    `json:"line"`
	ts   time.Time
}

func (e *nGCPError) Error() string {
	return fmt.Sprintf("%s/%d - %s - %s", e.File, e.Line, e.ts, e.Msg)
}

type driverError struct {
	err error
}

func (e *driverError) Set(err error) {
	if e.err == nil {
		e.err = err
	}
}

func (e *driverError) Get() error {
	return e.err
}
