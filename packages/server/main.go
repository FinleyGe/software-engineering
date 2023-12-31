package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	. "se/config"
	"se/grpc"
	_ "se/log"
	. "se/router"
	"syscall"
	"time"
)

func main() {
	var server *http.Server
	var port string = ":" + Config.Server.Port
	var grpcPort string = ":" + Config.Grpc.Port
	log.Println("Running Server at", port)
	log.Println("Running GRPC Server at", grpcPort)

	server = &http.Server{
		Addr:    port,
		Handler: Router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	go func() {
		if err := grpc.GrpcServer.Serve(grpc.Lis); err != nil {
			log.Fatalf("[GRPC] Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer grpc.GrpcServer.Stop()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server Exiting")
}
