package client

import (
	"fmt"
	"github.com/ok-chain/gosdk/common/libs/pkg/errors"
	"github.com/ok-chain/gosdk/common/transactParams"
	"github.com/ok-chain/gosdk/crypto/keys"
	"github.com/ok-chain/gosdk/types"
	"github.com/ok-chain/gosdk/types/msg"
	"github.com/ok-chain/gosdk/types/tx"
	"github.com/ok-chain/gosdk/utils"
)

// broadcast mode
const (
	BroadcastBlock = "block"
	BroadcastSync  = "sync"
	BroadcastAsync = "async"
)

func (cli *OKChainClient) Send(fromInfo keys.Info, passWd, toAddr, coinsStr, memo string, accNum, seqNum uint64) (resp types.TxResponse, err error) {
	if !transactParams.IsValidSendParams(fromInfo, passWd, toAddr) {
		return types.TxResponse{}, errors.New("err : params input to send are invalid")
	}

	to, err := types.AccAddressFromBech32(toAddr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Address [%s] error: %s", toAddr, err)
	}

	coins, err := utils.ParseCoins(coinsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := msg.NewMsgTokenSend(fromInfo.GetAddress(), to, coins)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

func (cli *OKChainClient) NewOrder(fromInfo keys.Info, passWd, product, side, price, quantity, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if !transactParams.IsValidNewOrderParams(fromInfo, passWd, product, side, price, quantity, memo, ) {
		return types.TxResponse{}, errors.New("err : params input to pend a order are invalid")
	}
	msg := msg.NewMsgNewOrder(fromInfo.GetAddress(), product, side, price, quantity)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}

func (cli *OKChainClient) CancelOrder(fromInfo keys.Info, passWd, orderID, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if !transactParams.IsValidCancelOrderParams(fromInfo, passWd) {
		return types.TxResponse{}, errors.New("err : params input to cancel a order are invalid")
	}

	msg := msg.NewMsgCancelOrder(fromInfo.GetAddress(), orderID)
	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

func (cli *OKChainClient) MultiSend(fromInfo keys.Info, passWd, transferStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if !transactParams.IsValidMultiSend(fromInfo, passWd, transferStr) {
		return types.TxResponse{}, errors.New("err : params input to multi send are invalid")
	}

	transfers, err := utils.StrToTransfers(transferStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("parse Transfers [%s] error: %s", err, transferStr)
	}

	msg := msg.NewMsgMultiSend(fromInfo.GetAddress(), transfers)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

func (cli *OKChainClient) Mint(fromInfo keys.Info, passWd, symbol string, amount int64, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if !transactParams.IsValidMint(fromInfo, passWd, symbol, amount) {
		return types.TxResponse{}, errors.New("err : params input to mint are invalid")
	}

	msg := msg.NewMsgMint(symbol, amount, fromInfo.GetAddress())

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}
