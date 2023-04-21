package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hyperedger/task/client/app"
)

func main() {
	//Initialize setup for Org1
	fmt.Println(len(os.Args), os.Args)
	if len(os.Args) != 2 {
		panic("config file path not provided")
	}

	bz, err := ioutil.ReadFile(os.Args[1])
	var config app.Config
	err = json.Unmarshal(bz, &config)
	if err != nil {
		fmt.Println("error while unmarshalling data")
		return
	}

	orgConfig, err := app.Initialize(config)
	if err != nil {
		fmt.Println("Error initializing setup for Org1: ", err)
	}
	app.Serve(app.Config(*orgConfig))
}
