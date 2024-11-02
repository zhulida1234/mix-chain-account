package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/zhulida1234/mix-chain-account/chain"
	"github.com/zhulida1234/mix-chain-account/config"
	"github.com/zhulida1234/mix-chain-account/rpc/account"
	common2 "github.com/zhulida1234/mix-chain-account/rpc/common"
	"math/big"
	"time"
)

const ChainName = "Ethereum"

type EthereumAdaptor struct {
	ethClient  EthClient
	DataClient *EthData
}

func (c EthereumAdaptor) ValidAddress(req *account.ValidAddressRequest) (*account.ValidAddressResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) GetBlockByHash(req *account.BlockHashRequest) (*account.BlockResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) GetBlockHeaderByHash(req *account.BlockHeaderHashRequest) (*account.BlockHeaderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) GetBlockHeaderByNumber(req *account.BlockHeaderNumberRequest) (*account.BlockHeaderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) GetAccount(req *account.AccountRequest) (*account.AccountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) GetFee(req *account.FeeRequest) (*account.FeeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) SendTx(req *account.SendTxRequest) (*account.SendTxResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) GetTxByAddress(req *account.TxAddressRequest) (*account.TxAddressResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) GetTxByHash(req *account.TxHashRequest) (*account.TxHashResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) GetBlockByRange(req *account.BlockByRangeRequest) (*account.BlockByRangeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) CreateUnSignTransaction(req *account.UnSignTransactionRequest) (*account.UnSignTransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) BuildSignedTransaction(req *account.SignedTransactionRequest) (*account.SignedTransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) DecodeTransaction(req *account.DecodeTransactionRequest) (*account.DecodeTransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) VerifySignedTransaction(req *account.VerifyTransactionRequest) (*account.VerifyTransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c EthereumAdaptor) GetExtraData(req *account.ExtraDataRequest) (*account.ExtraDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewChainAdaptor(conf *config.Config) (chain.IChainAdaptor, error) {
	client, err := DialEthClient(context.Background(), conf.WalletNode.Eth.RPCs[0].RpcUrl)
	if err != nil {
		log.Error("初始化EthClient 失败", "err", err)
		return nil, err
	}
	dataClient, err := NewEthDataClient(conf.WalletNode.Eth.DataApiUrl, conf.WalletNode.Eth.DataApiKey, time.Duration(conf.WalletNode.Eth.Timeout))
	if err != nil {
		log.Error("初始化DataClient 失败", "err", err)
		return nil, err
	}
	return &EthereumAdaptor{
		ethClient:  client,
		DataClient: dataClient,
	}, nil
}

func (c EthereumAdaptor) GetSupportChains(req *account.SupportChainsRequest) (*account.SupportChainsResponse, error) {
	return &account.SupportChainsResponse{
		Code:    common2.ReturnCode_SUCCESS,
		Msg:     "Support this chain",
		Support: true,
	}, nil
}

func (c EthereumAdaptor) ConvertAddress(req *account.ConvertAddressRequest) (*account.ConvertAddressResponse, error) {
	addressCommon := common.BytesToAddress(crypto.Keccak256(req.PublicKey[1:])[12:])
	return &account.ConvertAddressResponse{
		Code:    common2.ReturnCode_SUCCESS,
		Msg:     "convert address successs",
		Address: addressCommon.String(),
	}, nil
}

func (c EthereumAdaptor) GetBlockByNumber(req *account.BlockNumberRequest) (*account.BlockResponse, error) {
	block, err := c.ethClient.BlockByNumber(big.NewInt(req.Height))
	if err != nil {
		log.Error("block by number error", err)
		return &account.BlockResponse{
			Code: common2.ReturnCode_ERROR,
			Msg:  "block by number error",
		}, nil
	}
	var txListRet []*account.BlockInfoTransactionList
	for _, v := range block.Transactions {
		bitlItem := &account.BlockInfoTransactionList{
			From:   "0x000",
			To:     v.To,
			Hash:   v.Hash,
			Time:   "0",
			Amount: "10",
			Fee:    "0",
			Status: "1",
		}
		txListRet = append(txListRet, bitlItem)
	}
	return &account.BlockResponse{
		Code:         common2.ReturnCode_SUCCESS,
		Msg:          "block by number success",
		Hash:         block.Hash.String(),
		BaseFee:      block.BaseFee,
		Transactions: txListRet,
	}, nil
}
