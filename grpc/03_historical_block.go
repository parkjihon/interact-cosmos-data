package grpc

import (
	"context"
	"fmt"

	"interact-cosmos-data/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	sdk "github.com/cosmos/cosmos-sdk/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func QueryState03() error {
	config := types.GetConfig()
	myAddress, err := sdk.AccAddressFromBech32(config.Grpc.ValidatorAddress)
	if err != nil {
		return err
	}

	grpcConn, err := grpc.Dial(
		config.Grpc.GrpcAddress,
		grpc.WithInsecure(),
		//grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	bankClient := banktypes.NewQueryClient(grpcConn)
	var header metadata.MD
	bankRes, err := bankClient.Balance(
		metadata.AppendToOutgoingContext(
			context.Background(),
			grpctypes.GRPCBlockHeightHeader,
			"12",
		), // Add metadata to request
		&banktypes.QueryBalanceRequest{
			Address: myAddress.String(),
			Denom:   "uatom",
		},
		grpc.Header(&header), // retrieve header from response
	)
	if err != nil {
		return err
	}
	blockHeight := header.Get(grpctypes.GRPCBlockHeightHeader)
	fmt.Println(blockHeight)
	fmt.Println(len(blockHeight))
	fmt.Println(bankRes)

	return nil
}
