package dex

import (
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	"github.com/okex/okchain-go-sdk/module/auth"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/utils"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestDexClient_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)
	res, err := mockCli.Dex().List(fromInfo, passWd, "btc", "okt", "1.024", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)
}

func TestDexClient_Deposit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(),
		accInfo.GetSequence()).Return(mocks.DefaultMockSuccessTxResponse(), nil)
	res, err := mockCli.Dex().Deposit(fromInfo, passWd, product, "10.24okt", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)
}

func TestDexClient_Withdraw(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(),
		accInfo.GetSequence()).Return(mocks.DefaultMockSuccessTxResponse(), nil)
	res, err := mockCli.Dex().Withdraw(fromInfo, passWd, product, "1.024okt", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)
}

func TestDexClient_TransferOwnership(t *testing.T) {
	// set up signedTx.json
	err := ioutil.WriteFile(signedPath, []byte(expectedSignedTxJSON), 0644)
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "1024okt", 11, 22)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(),
		accInfo.GetSequence()).Return(mocks.DefaultMockSuccessTxResponse(), nil)
	res, err := mockCli.Dex().TransferOwnership(fromInfo, passWd, signedPath, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	// remove the temporary file: signedTx.json
	err = os.Remove(signedPath)
	require.NoError(t, err)

}
