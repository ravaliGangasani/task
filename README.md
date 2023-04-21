# task

# Running the test-network

[use the fabric test network](https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html)

# Deploy a chaincode on the channel

# Run REST client
- cd into task/client directory
- Run `go run main.go` to run the REST server

# Request Body
``` sh
type RequestBody struct {
	ChaincodeID string   `json:"chain_code_id"`
	ChannelID   string   `json:"channel_id"`
	Function    string   `json:"function"`
	Args        []string `json:"args"`
}
```

# Request Example
- Initiate a shipment
```url: http://localhost:3000/invoke
   sample request body: 
   {
    "chain_code_id":"basic",
    "channel_id":"channel1",
    "function":"CreateShipment",
    "args":["asset1","T-Shirt","blue","medium","ravali","ramya","hyderabad"]
}
```

- Update shipment status
```url: http://localhost:3000/invoke
   sample request body: 
   {
    "chain_code_id":"basic",
    "channel_id":"channel1",
    "function":"UpdateShipment",
    "args":["asset1","shipped_from_manufacturer"]
}
```

- GetAllShipments
```https://localhost:3000/query?channelid=channel1&chaincodeid=basic&function=GetAllShipments```

- Get shipment by assetID
```https://localhost:3000/query?channelid=channel1&chaincodeid=basic&function=GetShipment&args=asset1```

