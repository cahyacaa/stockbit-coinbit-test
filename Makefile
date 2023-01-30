http-service :
	go run cmd/app/main.go

balance-consumer :
	go run cmd/app/main.go -balance

deposit-flagger-consumer :
	go run cmd/app/main.go -depositFlagger