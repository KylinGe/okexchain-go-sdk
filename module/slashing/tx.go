package slashing

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	"github.com/okex/okexchain/x/slashing"
)

// Unjail unjails the own validator which was jailed by slashing module
func (sc slashingClient) Unjail(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (resp sdk.TxResponse,
	err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	msg := slashing.NewMsgUnjail(sdk.ValAddress(fromInfo.GetAddress()))
	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}
