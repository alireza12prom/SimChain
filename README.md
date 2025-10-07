# SimChain

A fun, simple and educational blockchain project, implemented by Golang, just to get familiar with blockchain stuff.

## ⚒ Tech Stack

- Golang `v1.25.0`
- Gin `v4`
- BadgerDB `v4`

## 📁 Directory Structure

```
- internal
  - core
    - block.go
    - blockchain.go
    - transaction.go
  - main.go (start point)
```

## 🚀 REST APIs

This project exposes a small HTTP API for interacting with the simulated blockchain. Run the server from the repository root:

```powershell
go run internal/main.go
```

Base URL: http://localhost:8080 (default)

### Endpoints:
- POST `/transaction/new`
  - Description: submit a new transaction into the pending pool
  - Request JSON:
    {
      "from": "<sender>",
      "to": "<recipient>",
      "amount": 1.23
    }
  - Success response: 200 OK with a JSON confirmation

- GET `/transaction/pending`
  - Description: returns the current pending transactions
  - Success response: 200 OK
    {
      "data": [ /* transactions */ ],
      "count": 3
    }

- POST `/block/mine`
  - Description: trigger mining of a new block using pending transactions
  - Success response: 200 OK with a short message describing the mined block

- GET `/block/history`
  - Description: return the blockchain history (all mined blocks)
  - Success response: 200 OK with block list
