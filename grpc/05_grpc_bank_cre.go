package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func QueryState05() error {
	//myAddress, err := sdk.AccAddressFromBech32("cosmos1zgwx3cwyyx8np35hlzngmkfdalnrjxj2450sul")
	//myAddress, err := sdk.AccAddressFromBech32("cre1zgwx3cwyyx8np35hlzngmkfdalnrjxj23uu4fj")
	// if err != nil {
	// 	return err
	// }

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		//"testnet-endpoint.crescent.network:9090", // your gRPC server address.
		"localhost:9090",
		grpc.WithInsecure(), // The SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		fmt.Println("하 시발", err)
	}
	defer grpcConn.Close()

	// This creates a gRPC client to query the x/bank service.
	bankClient := banktypes.NewQueryClient(grpcConn)
	bankRes, err := bankClient.Balance(
		context.Background(),
		//&banktypes.QueryBalanceRequest{Address: myAddress.String(), Denom: "atom"},
		&banktypes.QueryBalanceRequest{Address: "cre1zgwx3cwyyx8np35hlzngmkfdalnrjxj23uu4fj", Denom: "ucre"},
	)
	if err != nil {
		fmt.Println("하 시발222", err)
		return err
	}

	fmt.Println(bankRes.GetBalance()) // Prints the account balance

	return nil
}
