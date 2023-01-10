package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/b-harvest/modules-test-tool/wallet"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	txtype "github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	liquiditytypes "github.com/crescent-network/crescent/v3/x/liquidity/types"

	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"google.golang.org/grpc"
)

func QueryState07() error {
	myMne := "assist field six actual frown fury diet bulb address hello van swamp sustain dolphin ordinary short teach mammal digital garage legal labor ride pioneer"
	_, privKey, err := wallet.RecoverAccountFromMnemonic(myMne, "")
	if err != nil {
		return err
	}
	priv := cryptotypes.PrivKey(privKey)

	// Create msg for MMOrder
	tmpMaxSellPrice, _ := types.NewDecFromStr("0.0054100")
	tmpMinSellPrice, _ := types.NewDecFromStr("0.0053801")
	tmpMaxBuyPrice, _ := types.NewDecFromStr("0.0053800")
	tmpMinBuyPrice, _ := types.NewDecFromStr("0.0053501")
	msg := liquiditytypes.MsgMMOrder{
		Orderer:       addr,
		PairId:        4,
		MaxSellPrice:  tmpMaxSellPrice,
		MinSellPrice:  tmpMinSellPrice,
		SellAmount:    types.NewInt(4000000000), //500
		MaxBuyPrice:   tmpMaxBuyPrice,
		MinBuyPrice:   tmpMinBuyPrice,
		BuyAmount:     types.NewInt(4000000000), //500
		OrderLifespan: time.Hour,                // 1시간
	}

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",    // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()

	// we use Protobuf, given by the following function.
	encCfg := simapp.MakeTestEncodingConfig()
	// Create a new TxBuilder.
	txBuilder := encCfg.TxConfig.NewTxBuilder()
	if err := txBuilder.SetMsgs(&msg); err != nil {
		return err
	}
	txBuilder.SetGasLimit(uint64(1000000))

	// To find accounts' number & seq
	authClient := authtypes.NewQueryClient(grpcConn)
	queryAccountReq := authtypes.QueryAccountRequest{
		Address: addr,
	}
	queryAccountResp, err := authClient.Account(
		context.Background(),
		&queryAccountReq,
	)
	if err != nil {
		return err
	}
	var baseAccount authtypes.BaseAccount
	err = baseAccount.Unmarshal(queryAccountResp.GetAccount().Value)
	if err != nil {
		return err
	}
	accNum := baseAccount.GetAccountNumber()
	accSeq := baseAccount.GetSequence()

	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	sigV2 := signing.SignatureV2{
		PubKey: priv.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: accSeq,
	}
	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}

	// Second round: all signer infos are set, so each signer can sign.
	sigV2 = signing.SignatureV2{}
	signerData := xauthsigning.SignerData{
		ChainID:       "mooncat-2-external",
		AccountNumber: accNum,
		Sequence:      accSeq,
	}
	sigV2, err = tx.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, priv, encCfg.TxConfig, accSeq)
	if err != nil {
		return err
	}
	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}

	// Generated Protobuf-encoded bytes.
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return err
	}

	// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx service.
	txClient := txtype.NewServiceClient(grpcConn)
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&txtype.BroadcastTxRequest{
			Mode:    txtype.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		return err
	}
	fmt.Println(grpcRes.TxResponse) // Code Should be `0` if the tx is successful

	return nil
}
