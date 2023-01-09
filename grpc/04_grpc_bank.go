package grpc

import (
	"context"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"

	"interact-cosmos-data/types"
)

func QueryState04() (error, error) {
	config := types.GetConfig()
	var (
		GRPC_ADDRESS      = config.Grpc.GrpcAddress
		VALIDATOR_ADDRESS = config.Grpc.ValidatorAddress
	)
	fmt.Println("QueryState started...")
	// get config & address
	myAddress, err := sdk.AccAddressFromBech32(VALIDATOR_ADDRESS)
	if err != nil {
		return err, errors.New("sdk.AccAddressFromBech32")
	}

	// Create a connection to the gRPC server
	grpcConn, err := grpc.Dial(
		GRPC_ADDRESS, // your gRPC server address
		grpc.WithInsecure(),
		//grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		return err, errors.New("grpc.Dial")
	}
	defer grpcConn.Close()

	// Create a gRPC client to query the x/bank service.
	bankClient := banktypes.NewQueryClient(grpcConn)
	bankRes, err := bankClient.Balance(
		context.Background(),
		&banktypes.QueryBalanceRequest{
			Address: myAddress.String(),
			Denom:   "uatom",
		},
	)
	if err != nil {
		return err, errors.New("NewQueryClient")
	}

	// 한 계정의 모든 coin 추출
	bankAllBalances, err := bankClient.AllBalances(
		context.Background(),
		&banktypes.QueryAllBalancesRequest{
			Address: myAddress.String(),
		},
	)

	bankDenomMetadata, err := bankClient.DenomMetadata(
		context.Background(),
		&banktypes.QueryDenomMetadataRequest{
			Denom: "uatom",
		},
	)

	fmt.Println("GetBalance   : ", bankRes.GetBalance())
	fmt.Println("AllBalances  : ", bankAllBalances.GetBalances())
	fmt.Println("DenomMetadata: ", bankDenomMetadata.GetMetadata())

	return nil, nil
}
