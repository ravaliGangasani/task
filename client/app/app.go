package app

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type Config struct {
	OrgName      string `json:"org_name"`
	MSPID        string `json:"msp_id"`
	CryptoPath   string `json:"crypto_path"`
	CertPath     string `json:"cert_path"`
	KeyPath      string `json:"key_path"`
	TLSCertPath  string `json:"tls_cert_path"`
	PeerEndpoint string `json:"peer_end_point"`
	GatewayPeer  string `json:"gateway_peer"`
	Gateway      client.Gateway
}

func Serve(config Config) {
	http.HandleFunc("/query", config.Query)
	http.HandleFunc("/invoke", config.Invoke)
	fmt.Println("Listening (http://localhost:3000/)...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
