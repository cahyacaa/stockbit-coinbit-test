package main

import (
	"context"
	"flag"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/above_threshold"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	brokers  = []string{"localhost:9092"}
	deposits = flag.Bool("deposits", false, "run collector processor")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())

	// Create topics if they do not already exist
	if *deposits {
		above_threshold.PrepareTopics(brokers)
		go func() {
			err := above_threshold.Run(ctx, brokers)
			if err != nil {
				log.Fatal(err())
			}
		}()
	} else {
		go func() {
			service.Run(brokers, above_threshold.Deposits)
		}()
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		log.Println(sig.String())
		done <- true
		cancel()
	}()
	<-done
	log.Println("app closed")
}
