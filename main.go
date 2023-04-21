package main

import (
	"log"

	"github.com/hyperedger/task/chaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	shipmentChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("error creating shipment chaincode: %v ", err)
	}
	if err = shipmentChaincode.Start(); err != nil {
		log.Panicf("error starting shipment chaincode: %v", err)
	}
}
