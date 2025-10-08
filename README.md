# SimChain

A fun, simple and educational blockchain project, implemented by Golang, just to get familiar with blockchain stuff.

## ⚒ Tech Stack

- Golang `v1.25.0`
- Gin `v4`
- BadgerDB `v4`

## 📁 Directory Structure

```
simchain/
├── cmd/
│   └── simchain/
│       └── main.go          # Wires deps
├── internal/
│   ├── domain/              # Pure core
│   │   ├── block.go
│   │   ├── blockchain.go
│   │   └── transaction.go
│   ├── blockchain/          # Implements domain interfaces
│   │   ├── blockchain.go
│   │   ├── block_validator.go
│   │   └── tx_pool.go       
│   ├── p2p/                 # Network layer: P2P gossip, discovery
│   │   ├── node.go
│   │   ├── protocol.go
│   │   └── broadcaster.go
│   ├── api/                 # HTTP/REST layer
│   │   ├── server.go
│   │   ├── handlers/ 
│   │   │   ├── block.go
│   │   │   └── transaction.go
│   └── utils/
│       └── hash.go
├── go.mod
├── go.sum
└── README.md
```

## 🚀 REST APIs

This project exposes a small HTTP API for interacting with the simulated blockchain. Run the server from the repository root (or build a binary):

```powershell
go run internal/main.go
# or build then run (Windows PowerShell)
go build -o simchain ./internal; .\simchain.exe
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
