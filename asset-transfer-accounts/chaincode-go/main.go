package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Account struct {
	DealerID    string  `json:"dealerId"`
	MSISDN      string  `json:"msisdn"`
	MPIN        string  `json:"mpin"`
	Balance     float64 `json:"balance"`
	Status      string  `json:"status"`
	TransAmount float64 `json:"transAmount"`
	TransType   string  `json:"transType"`
	Remarks     string  `json:"remarks"`
}

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) CreateAccount(ctx contractapi.TransactionContextInterface, dealerID, msisdn, mpin string, balance float64, status string, transAmount float64, transType, remarks string) error {
	exists, err := s.AccountExists(ctx, dealerID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("account %s already exists", dealerID)
	}

	account := Account{
		DealerID:    dealerID,
		MSISDN:      msisdn,
		MPIN:        mpin,
		Balance:     balance,
		Status:      status,
		TransAmount: transAmount,
		TransType:   transType,
		Remarks:     remarks,
	}

	accountJSON, err := json.Marshal(account)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(dealerID, accountJSON)
}

func (s *SmartContract) ReadAccount(ctx contractapi.TransactionContextInterface, dealerID string) (*Account, error) {
	data, err := ctx.GetStub().GetState(dealerID)
	if err != nil {
		return nil, fmt.Errorf("failed to read account: %v", err)
	}
	if data == nil {
		return nil, fmt.Errorf("account %s does not exist", dealerID)
	}

	var account Account
	if err := json.Unmarshal(data, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *SmartContract) UpdateAccount(ctx contractapi.TransactionContextInterface, dealerID, status string, balance, transAmount float64, transType, remarks string) error {
	account, err := s.ReadAccount(ctx, dealerID)
	if err != nil {
		return err
	}

	account.Status = status
	account.Balance = balance
	account.TransAmount = transAmount
	account.TransType = transType
	account.Remarks = remarks

	accountJSON, err := json.Marshal(account)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(dealerID, accountJSON)
}

func (s *SmartContract) GetAccountHistory(ctx contractapi.TransactionContextInterface, dealerID string) ([]*Account, error) {
	historyIter, err := ctx.GetStub().GetHistoryForKey(dealerID)
	if err != nil {
		return nil, err
	}
	defer historyIter.Close()

	var history []*Account
	for historyIter.HasNext() {
		modification, err := historyIter.Next()
		if err != nil {
			return nil, err
		}
		var account Account
		if err := json.Unmarshal(modification.Value, &account); err == nil {
			history = append(history, &account)
		}
	}

	return history, nil
}

func (s *SmartContract) AccountExists(ctx contractapi.TransactionContextInterface, dealerID string) (bool, error) {
	data, err := ctx.GetStub().GetState(dealerID)
	if err != nil {
		return false, err
	}
	return data != nil, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode: %v", err))
	}
	if err := chaincode.Start(); err != nil {
		panic(fmt.Sprintf("Error starting chaincode: %v", err))
	}
}
