package grpc

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func QueryState02() (error, error) {
	fmt.Println("QueryState started...")
	myAddress, err := sdk.AccAddressFromBech32(VALIDATOR_ADDRESS)
	if err != nil {
		return err, errors.New("sdk.AccAddressFromBech32")
	}

	// Create a connection to the gRPC server
	grpcConn, err := grpc.Dial(
		GRPC_ADDRESS, // your gRPC server address
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
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
	fmt.Println(bankRes.GetBalance())
	return nil, nil
}
