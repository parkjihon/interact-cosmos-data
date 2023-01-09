package types

import (
	"github.com/BurntSushi/toml"
)

type configStruct struct {
	Title string
	Grpc  grpcStruct
	Eth   ethStruct
}

type grpcStruct struct {
	GrpcAddress      string `toml:"grpc_address"`
	ValidatorAddress string `toml:"validator_address"`
}

type ethStruct struct {
	Etherscan_Mainnet_Endpoint string `toml:"etherscan_mainnet_endpoint"`
	Etherscan_Key              string `toml:"etherscan_key"`
}

func GetConfig() configStruct {
	var config configStruct
	_, err := toml.DecodeFile("config/config.toml", &config)
	if err != nil {
		panic(err)
	}

	return config
}
