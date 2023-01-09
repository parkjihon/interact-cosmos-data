package grpc

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/types"
	liquiditytypes "github.com/crescent-network/crescent/v3/x/liquidity/types"
	"google.golang.org/grpc"
)

func QueryState07() error {
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

	// create msg for MMOrder
	msg := liquiditytypes.MsgMMOrder{
		Orderer: addr,
		PairId:  4,
		//MaxSellPrice: types.NewDec(0.00535),
		MaxSellPrice:  types.NewDec(5),
		MinSellPrice:  types.NewDec(4),
		SellAmount:    types.NewInt(1000000000), //1,000
		MaxBuyPrice:   types.NewDec(3),
		MinBuyPrice:   types.NewDec(2),
		BuyAmount:     types.NewInt(1000000000),              //1,000
		OrderLifespan: time.Duration(time.Duration.Hours(1)), // 1시간
	}
	fmt.Println("MMOder 메시지 말기")
	fmt.Println(msg)

	return nil
}
