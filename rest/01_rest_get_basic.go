package rest

import (
	"fmt"
	"interact-cosmos-data/types"
	"net/http"
)

func QueryRest01() error {
	config := types.GetConfig()
	fmt.Println(config.Eth.Etherscan_Key)
	fmt.Println(config.Eth.Etherscan_Mainnet_Endpoint)

	// GET
	resp, err := http.Get(config.Eth.Etherscan_Mainnet_Endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// result handling
	//data, err := ioutil.ReadAll

	return nil
}
