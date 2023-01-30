package service

import (
	"encoding/json"
	"fmt"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/above_threshold"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/balance"
	proto_codec "github.com/cahyacaa/stockbit-coinbit-test/internal/proto"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/proto/proto_models"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lovoo/goka"
)

func Init(brokers []string, stream goka.Stream) (*mux.Router, *goka.View, *goka.View, *goka.Emitter) {
	viewBalance, err := goka.NewView(brokers, balance.Table, new(balance.BalanceCodec))
	if err != nil {
		log.Fatal(err)
	}

	viewFlagger, errFlagger := goka.NewView(brokers, above_threshold.Table, new(above_threshold.AboveThresholdCodec))
	if errFlagger != nil {
		log.Fatal(err)
	}

	emitter, err := goka.NewEmitter(brokers, stream, new(proto_codec.ProtoCodec))
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/deposits", send(emitter)).Methods("POST")
	router.HandleFunc("/{wallet_id}/details", feed(viewBalance, viewFlagger)).Methods("GET")
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
	return router, viewFlagger, viewBalance, emitter
}

func send(emitter *goka.Emitter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var d deposit.Deposit

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
		err = emitter.EmitSync(d.WalletID, d)

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

func feed(viewBalance, viewFlagger *goka.View) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := mux.Vars(r)["wallet_id"]
		val, _ := viewBalance.Get(user)

		flaggerData, _ := viewFlagger.Get(user)

		if val == nil {
			fmt.Fprintf(w, "%s not found!", user)
			return
		}
		balanceData := val.(deposit.Balance)
		log.Printf("Latest balance for wallet ID %s, is : %v\n", user, balanceData.Balance)

		balanceData.IsAboveThreshold = flaggerData.(deposit.DepositFlagger).IsAboveThreshold

		out, err := json.Marshal(struct {
			Message          string  `json:"message"`
			Balance          float32 `json:"balance"`
			IsAboveThreshold bool    `json:"isAboveThreshold"`
		}{
			Message:          "OK",
			Balance:          balanceData.Balance,
			IsAboveThreshold: balanceData.IsAboveThreshold,
		})

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(out)
		if err != nil {
			return
		}

	}
}
