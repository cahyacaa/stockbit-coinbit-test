package main

import (
	"context"
	"flag"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/above_threshold"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/balance"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/service"
	"log"
	"net/http"
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

	//init http service and goka view
	router, viewFlagger, viewBalance, emitter := service.Init(brokers, above_threshold.Deposits)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Create topics if they do not already exist
	above_threshold.PrepareTopics(brokers)

	if *flaggerConsumer {
		go func() {
			err := above_threshold.Run(ctx, brokers)
			if err != nil {
				log.Fatal(err())
			}
		}()
	} else if *balanceConsumer {
		go func() {
			err := balance.Run(ctx, brokers)
			if err != nil {
				log.Fatal(err())
			}
		}()
	} else {
		//running goka view topic service
		go func() {
			err := viewFlagger.Run(context.Background())
			if err != nil {
				log.Fatal(err)
			}
		}()

		go func() {
			err := viewBalance.Run(context.Background())
			if err != nil {
				log.Fatal(err)
			}
		}()

		//running http server
		go func() {
			log.Println("Http server is running")
			err := srv.ListenAndServe()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		log.Println(sig.String())
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}
		err = emitter.Finish()
		if err != nil {
			log.Fatal(err)
		}
		done <- true
		cancel()
	}()
	<-done
	log.Println("app closed")
}
