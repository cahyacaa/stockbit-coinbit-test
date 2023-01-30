# [Coinbit] Golang Engineer Test

## Features :

- Deposit
- Balance info
- Deposit flagger if amount above threshold within certain time

## Tech Stack :
- [Go Language](https://go.dev/)
- [Goka](https://github.com/lovoo/goka)
- [Apache Kafka](https://kafka.apache.org/)


## Setup App :
## There are few steps to run the app

#### 1. download all modules
```sh
go mod tidy
```

#### 2. run docker compose for spin up kafka
```sh
docker compose up -d
```

### In this app there are 3 services : `balance consumer service`, `deposit flagger service`, `http service`. To run all service there must be 3 separated terminal

#### 3. run balance  service :
```sh
make balace-service
```
#### 4. run deposit flagger service :
```sh
make deposit-flagger-service
```
#### 5. run http service :
```sh
make http-service
```

## API List
- ### deposit [POST]
    - URL
    ```http request
     http://localhost:8080/deposits
    ```
    - payload
  ```json
    {
        "walletID":"1",
        "amount":10000
    }
  ```
  - response 
  ```json
    {
        "message": "OK"
    }
   ```
  
- ### balance details [GET]
    - URL
    ```http request
    http://localhost:8080/:wallet_id/details
    ```
    - response
  ```json
    {
        "message": "OK",
        "balance": 50000,
        "isAboveThreshold": true
    }   
   ```