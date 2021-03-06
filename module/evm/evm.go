package evm

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/evm/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	evm "github.com/okex/okexchain/x/evm/types"
)

var _ gosdktypes.Module = (*evmClient)(nil)

type evmClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in evm module
func (ec evmClient) RegisterCodec(cdc *codec.Codec) {
	evm.RegisterCodec(cdc)
}

// Name returns the module name
func (evmClient) Name() string {
	return types.ModuleName
}

// NewEvmClient creates a new instance of evm client as implement
func NewEvmClient(baseClient gosdktypes.BaseClient) exposed.Evm {
	return evmClient{baseClient}
}
