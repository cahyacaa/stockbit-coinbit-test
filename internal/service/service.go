package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/above_threshold"
	wallet "github.com/cahyacaa/stockbit-coinbit-test/model"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lovoo/goka"
)

func Run(brokers []string, stream goka.Stream) {
	view, err := goka.NewView(brokers, above_threshold.Table, new(above_threshold.AboveThresholdCodec))
	if err != nil {
		log.Fatal(err)
	}
	go view.Run(context.Background())

	emitter, err := goka.NewEmitter(brokers, stream, new(above_threshold.AboveThresholdCodec))
	if err != nil {
		log.Fatal(err)
	}
	defer emitter.Finish()

	router := mux.NewRouter()
	router.HandleFunc("/deposits", send(emitter, stream)).Methods("POST")
	router.HandleFunc("/{wallet_id}/details", feed(view)).Methods("GET")
	router.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		out, err := json.Marshal(map[string]string{
			"message": "OK",
		})

		if err != nil {
			log.Println(err)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		if _, err := writer.Write(out); err != nil {
			log.Println(err)
			return
		}

	}).Methods("GET")

	log.Printf("Listen port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func send(emitter *goka.Emitter, stream goka.Stream) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var d wallet.Wallet

		b, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		err = json.Unmarshal(b, &d)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}
		err = emitter.EmitSync(d.WalletId, d)

		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}
		log.Printf("Sent message: %v \n", d.String())

		out, err := json.Marshal(map[string]string{
			"message": "OK",
		})

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(out)
		if err != nil {
			return
		}
	}
}

func feed(view *goka.View) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := mux.Vars(r)["wallet_id"]
		val, _ := view.Get(user)
		fmt.Println(val)

		if val == nil {
			fmt.Fprintf(w, "%s not found!", user)
			return
		}
		messages := val.(wallet.Wallet)
		fmt.Printf("Latest balance for %s, is : %v\n", user, messages.Amount)

		out, err := json.Marshal(messages)

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(out)
		if err != nil {
			return
		}

	}
}
