package main

import (
	"context"
	"flag"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/above_threshold"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/balance"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	brokers         = []string{"localhost:9092"}
	flaggerConsumer = flag.Bool("depositFlagger", false, "run above_threshold flagger processor")
	balanceConsumer = flag.Bool("balance", false, "run balance processor")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())

	// Create topics if they do not already exist
	above_threshold.PrepareTopics(brokers)

	if *flaggerConsumer {
		go func() {
			err := above_threshold.Run(ctx, brokers)
			if err != nil {
				log.Println(err())
			}
		}()
	} else if *balanceConsumer {
		go func() {
			err := balance.Run(ctx, brokers)
			if err != nil {
				log.Println(err())
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
