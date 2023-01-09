package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	liquiditytypes "github.com/crescent-network/crescent/v3/x/liquidity/types"
)

var addr string = "cre1zgwx3cwyyx8np35hlzngmkfdalnrjxj23uu4fj"

func QueryState06() error {
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

	liquidityClient := liquiditytypes.NewQueryClient(grpcConn)

	lpPairs, err := liquidityClient.Pairs(
		context.Background(),
		&liquiditytypes.QueryPairsRequest{},
	)
	if err != nil {
		fmt.Println("pair 찾기", err)
		return err
	}
	fmt.Println(lpPairs)

	lpOrders, err := liquidityClient.OrdersByOrderer(
		context.Background(),
		&liquiditytypes.QueryOrdersByOrdererRequest{
			Orderer: addr,
			PairId:  4,
		},
	)
	if err != nil {
		fmt.Println("orders 찾기", err)
		return err
	}
	fmt.Println("내가 깐 주문")
	fmt.Println(lpOrders.GetOrders())

	return nil
}
