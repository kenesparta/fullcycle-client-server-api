# FullCycle client-server api

Fullcycle Pos Go Expert - Client Server API

> We are using golang workspaces on this project!

# 🚀 How to start?

1. If you haven't been created a golang workspace, you can run this command to create it: `make init`.
2. Then, open a terminal and run the server: `make run-server`.
3. Open another terminal window and run the client `make run-client`.

# Postman

## GET /cotacao

- Gets the information from `economia.awesomeapi.com`
- No input needed.
- Output:

```json
{
  "id": "f585ec8d-991f-4e14-8c6c-1c2cb5f635a4",
  "code": "BRL",
  "codein": "BRL",
  "name": "Dólar Americano/Real Brasileiro",
  "high": 4.9552,
  "low": 4.9008,
  "varBid": -0.0069,
  "pctChange": -0.14,
  "bid": 4.946,
  "ask": 4.949,
  "timestamp": "1693601998",
  "create_date": "2023-09-01 17:59:58"
}
```

## GET /read

- Get the list of saved requests.
- No input needed.
- Output:

```json
[
  {
    "id": "f585ec8d-991f-4e14-8c6c-1c2cb5f635a4",
    "code": "BRL",
    "codein": "BRL",
    "name": "Dólar Americano/Real Brasileiro",
    "high": 4.9552,
    "low": 4.9008,
    "varBid": -0.0069,
    "pctChange": -0.14,
    "bid": 4.946,
    "ask": 4.949,
    "timestamp": "1693601998",
    "create_date": "2023-09-01 17:59:58"
  },
  {
    "id": "410b0e24-7016-454c-8733-156f06093962",
    "code": "BRL",
    "codein": "BRL",
    "name": "Dólar Americano/Real Brasileiro",
    "high": 4.9552,
    "low": 4.9008,
    "varBid": -0.0069,
    "pctChange": -0.14,
    "bid": 4.946,
    "ask": 4.949,
    "timestamp": "1693601998",
    "create_date": "2023-09-01 17:59:58"
  },
  {
    "id": "a0bda865-fb49-4bb6-bf92-3bb8d47ba91a",
    "code": "BRL",
    "codein": "BRL",
    "name": "Dólar Americano/Real Brasileiro",
    "high": 4.9552,
    "low": 4.9008,
    "varBid": -0.0069,
    "pctChange": -0.14,
    "bid": 4.946,
    "ask": 4.949,
    "timestamp": "1693601998",
    "create_date": "2023-09-01 17:59:58"
  }
]
```