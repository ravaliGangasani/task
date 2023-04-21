package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type RequestBody struct {
	ChaincodeID string   `json:"chain_code_id"`
	ChannelID   string   `json:"channel_id"`
	Function    string   `json:"function"`
	Args        []string `json:"args"`
}

// Invoke handles chaincode invoke requests.
func (c *Config) Invoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "err: %s", err)
		return
	}

	var request RequestBody
	err = json.Unmarshal(body, &request)
	if err != nil {
		fmt.Fprintf(w, "err: %s", err)
		return
	}

	network := c.Gateway.GetNetwork(request.ChannelID)
	contract := network.GetContract(request.ChaincodeID)
	txn_proposal, err := contract.NewProposal(request.Function, client.WithArguments(request.Args...))
	if err != nil {
		fmt.Fprintf(w, "Error creating txn proposal: %s", err)
		return
	}
	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		fmt.Fprintf(w, "Error endorsing txn: %s", err)
		return
	}
	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		fmt.Fprintf(w, "Error submitting transaction: %s", err)
		return
	}
	fmt.Fprintf(w, "Transaction ID : %s", txn_committed.TransactionID())
}
