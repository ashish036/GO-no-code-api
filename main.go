package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	server "github.com/ashish036/GO-no-code-api/server"
	utils "github.com/ashish036/GO-no-code-api/utils"
)

const (
	SERVER_PORT = "9110"
)

func init() {
	utils.LoadEnvVariables()

	utils.InitializeLogger()
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		log.Printf("system call: %+v", oscall)
		cancel()
	}()

	if err := server.Serve(ctx, SERVER_PORT); err != nil {
		utils.GetLogHelper(ctx, "MAIN", http.StatusInternalServerError).Error(err.Error(), "Failed to serve")
	}
}
