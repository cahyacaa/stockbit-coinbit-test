http-service :
	go run cmd/app/main.go

balance-service :
	go run cmd/app/main.go -balance

deposit-flagger-service :
	go run cmd/app/main.go -depositFlagger
