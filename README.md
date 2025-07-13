# Hyperledger Fabric Blockchain Internship Project

This project is a complete solution for a blockchain internship assignment involving the development of a smart contract, deploying it in a Hyperledger Fabric test network, and building a REST API to interact with it.

---

## Project Structure

```
fabric-samples/
├── test-network/                    # Hyperledger Fabric network with Org1 & Org2
├── asset-transfer-accounts/
│   └── chaincode-go/
│       └── main.go                  # Smart contract written in Go
├── rest-api/
│   ├── index.js                     # Express-based REST API
│   ├── enrollAdmin.js              # Script to import admin identity
│   ├── wallet/                     # Wallet stores admin identity
│   ├── connection.json             # Fabric network connection profile
│   ├── Dockerfile                  # Docker image for the REST API
│   └── README.md                   # This file
```

---

## Chaincode Overview (Level 2)

The smart contract (`main.go`) manages telecom-style dealer accounts.

### Fields Stored
- `dealerId`
- `msisdn`
- `mpin`
- `balance`
- `status`
- `transAmount`
- `transType`
- `remarks`

### Chaincode Functions

| Function            | Description                              |
|---------------------|------------------------------------------|
| `CreateAccount`     | Creates a new account                    |
| `ReadAccount`       | Reads an account by dealer ID           |
| `UpdateAccount`     | Updates account balance/status          |
| `GetAccountHistory` | Shows full transaction history          |

---

## Chaincode Lifecycle (Tested via CLI)

```bash
# Step 1: Package
peer lifecycle chaincode package accounts.tar.gz ...

# Step 2: Install
peer lifecycle chaincode install accounts.tar.gz

# Step 3: Approve
peer lifecycle chaincode approveformyorg ...

# Step 4: Commit
peer lifecycle chaincode commit ...

# Step 5: Invoke
peer chaincode invoke ...
```

Transactions were tested and returned status `VALID`.

---

## REST API Overview (Level 3)

Built with Node.js and Express using Fabric Gateway SDK (`fabric-network`).

### Setup

```bash
cd rest-api
npm install
node enrollAdmin.js       # Imports the admin identity into the wallet
node index.js             # Starts the REST API
```

### Available Endpoints

| Method | Route                     | Chaincode Function      |
|--------|---------------------------|--------------------------|
| POST   | `/account`                | CreateAccount            |
| GET    | `/account/:id`            | ReadAccount              |
| PUT    | `/account/:id`            | UpdateAccount            |
| GET    | `/account/:id/history`    | GetAccountHistory        |

All tested and working via Postman and Gitpod browser.

---

## Docker Support

You can also build a container for the REST API.

### Dockerfile

```Dockerfile
FROM node:18
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 4000
CMD ["node", "index.js"]
```

### Build and Run

```bash
docker build -t fabric-rest-api .
docker run -p 4000:4000 fabric-rest-api
```

If using Docker with Fabric, update `localhost` to `host.docker.internal` in `connection.json`.

---

## How to Test

Use Postman or a browser for GET requests.

### Example: Create Account

```http
POST /account
Content-Type: application/json
```

```json
{
  "dealerId": "D001",
  "msisdn": "9876543210",
  "mpin": "1234",
  "balance": 1000,
  "status": "active",
  "transAmount": 0,
  "transType": "credit",
  "remarks": "New account created"
}
```

---

## Author

Arun Allanki  

---

## Project Completion Summary

- [x] Set up and ran Hyperledger Fabric test-network (Level 1)
- [x] Developed custom Go smart contract (Level 2)
- [x] Packaged, installed, and invoked chaincode
- [x] Built REST API with Node.js and Fabric Gateway SDK (Level 3)
- [x] Imported wallet identity
- [x] Dockerized the REST API

