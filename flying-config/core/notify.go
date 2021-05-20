package core

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"os/signal"
	"time"
)

func RunStopNotify(srv *grpc.Server) {
	processed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		_, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		srv.GracefulStop()
		log.Println("server gracefully shutdown")

		close(processed)
	}()
}

//func waitForGracefulShutdown(srv *grpc.Server) {
//	fmt.Println("Grpc messaging server started ...")
//	interruptChan := make(chan os.Signal, 1)
//	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
//
//	// Block until we receive our signal.
//	<-interruptChan
//
//	// Create a deadline to wait for.
//	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
//	defer cancel()
//	srv.GracefulStop()
//	publisher, messagingError := messaging.GetPublisherInstance()
//	if messagingError.Error == nil {
//		publisher.Close()
//	}
//
//	log.Println("Shutting down grpc messaging server.")
//	os.Exit(0)
