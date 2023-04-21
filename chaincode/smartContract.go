package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

const (
	INITIATE_SHIPMENT         = "initiate_shipment"
	SHIPPED_FROM_MANUFACTURER = "shipped_from_manufacturer"
	RECEIVED_BY_DEALER        = "received_by_dealer"
	RECEIVED_BY_WHOLESALER    = "received_by_wholesaler"
	RECEIVED_BY_RETAILER      = "received_by_retailer"
)

type Asset struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	Size   string `json:"size"`
	Seller string `json:"seller"`
}

type Order struct {
	Asset   Asset  `json:"asset"`
	Buyer   string `json:"buyer"`
	Address string `json:"address"`
	Status  string `json:"status"`
}

func (s *SmartContract) CreateShipment(ctx contractapi.TransactionContextInterface, id, name, color, size, seller, buyer, address string) error {

	existed, err := s.AlreadyInitiatedShipment(ctx, id)
	if existed {
		return fmt.Errorf("shipment is already initiated with id %s", id)
	}
	order := Order{Asset: Asset{
		ID:     id,
		Name:   name,
		Color:  color,
		Size:   size,
		Seller: seller,
	},
		Buyer:   buyer,
		Address: address,
		Status:  INITIATE_SHIPMENT,
	}

	orderJson, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(order.Asset.ID, orderJson)

}

func (s *SmartContract) UpdateShipment(ctx contractapi.TransactionContextInterface, id, status string) error {
	bz, _ := ctx.GetStub().GetState(id)
	if bz == nil {
		return fmt.Errorf("shipment is not initiated with id %s", id)
	}

	var order Order
	err := json.Unmarshal(bz, &order)
	if err != nil {
		return err
	}

	if status == SHIPPED_FROM_MANUFACTURER && order.Status != INITIATE_SHIPMENT {
		return fmt.Errorf("cannot update the order status from %s to %s", order.Status, SHIPPED_FROM_MANUFACTURER)
	} else if status == RECEIVED_BY_WHOLESALER && order.Status != SHIPPED_FROM_MANUFACTURER {
		return fmt.Errorf("cannot update the order status from %s to %s", order.Status, RECEIVED_BY_WHOLESALER)
	} else if status == RECEIVED_BY_DEALER && order.Status != RECEIVED_BY_WHOLESALER {
		return fmt.Errorf("cannot update the order status from %s to %s", order.Status, RECEIVED_BY_DEALER)
	} else if status == RECEIVED_BY_RETAILER && order.Status != RECEIVED_BY_DEALER {
		return fmt.Errorf("cannot update the order status from %s to %s", order.Status, RECEIVED_BY_RETAILER)
	}
	order.Status = status
	orderBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, orderBytes)
}

func (s *SmartContract) GetShipment(ctx contractapi.TransactionContextInterface, id string) (*Order, error) {
	bz, _ := ctx.GetStub().GetState(id)
	if bz == nil {
		return nil, fmt.Errorf("shipment is not present with the id %s", id)
	}

	var order Order
	err := json.Unmarshal(bz, &order)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling order shipment %s", err.Error())
	}

	return &order, nil
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	asset := Asset{ID: "asset1", Name: "kurthi", Color: "white", Size: "small", Seller: "ravali"}
	orders := []Order{
		{Asset: asset,
			Buyer:   "gangasani",
			Address: "hyderabad",
			Status:  INITIATE_SHIPMENT,
		},
	}
	for _, order := range orders {
		orderJSON, err := json.Marshal(order)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(order.Asset.ID, orderJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

func (s *SmartContract) AlreadyInitiatedShipment(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	bz, err := ctx.GetStub().GetState(id)
	if bz == nil {
		return false, err
	}

	var order *Order
	err = json.Unmarshal(bz, &order)
	if err != nil {
		return false, err
	}

	if order.Status == INITIATE_SHIPMENT {
		return true, nil
	}
	return false, nil
}

func (s *SmartContract) GetAllShipments(ctx contractapi.TransactionContextInterface) ([]*Order, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var orders []*Order
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var order Order
		err = json.Unmarshal(queryResponse.Value, &order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}
