package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/goinggo/tracelog"

	server "github.com/ashish036/GO-no-code-api/server"
	utils "github.com/ashish036/GO-no-code-api/utils"
)

const (
	SERVER_PORT = "9110"
)

func init() {
	utils.LoadEnvVariables()
}

func main() {
	tracelogMode := tracelog.LevelError
	tracelog.Start(tracelogMode)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		log.Printf("system call: %+v", oscall)
		cancel()
	}()

	_ = server.Serve(ctx, SERVER_PORT)
}
