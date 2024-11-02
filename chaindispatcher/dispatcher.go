package chaindispatcher

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/zhulida1234/mix-chain-account/chain"
	"github.com/zhulida1234/mix-chain-account/chain/ethereum"
	"github.com/zhulida1234/mix-chain-account/config"
	"github.com/zhulida1234/mix-chain-account/rpc/account"
	"github.com/zhulida1234/mix-chain-account/rpc/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"
	"strings"
)

type CommonRequest interface {
	GetChain() string
}

type CommonReply = account.SupportChainsResponse

type Dispatcher struct {
	Registry map[string]chain.IChainAdaptor
}

func NewDispatcher(conf *config.Config) (*Dispatcher, error) {
	dispatcher := Dispatcher{
		Registry: make(map[string]chain.IChainAdaptor),
	}

	chainAdaptorFactoryMap := map[string]func(conf *config.Config) (chain.IChainAdaptor, error){
		ethereum.ChainName: ethereum.NewChainAdaptor,
	}

	supportedChains := []string{
		ethereum.ChainName,
	}

	for _, c := range conf.Chains {
		if factory, ok := chainAdaptorFactoryMap[c]; ok {
			adaptor, err := factory(conf)
			if err != nil {
				log.Crit("failed to create chain adaptor", "error", err)
			}
			dispatcher.Registry[c] = adaptor
		} else {
			log.Error("unsupported chain", "chain", c, "supportedChains", supportedChains)
		}
	}
	return &dispatcher, nil
}

func (d *Dispatcher) Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Error("panic error", "msg", e)
			log.Debug(string(debug.Stack()))
			err = status.Errorf(codes.Internal, "panic error", e)
		}
	}()

	pos := strings.LastIndex(info.FullMethod, "/")
	method := info.FullMethod[pos+1:]

	chainName := req.(CommonRequest).GetChain()
	log.Info(method, "chain", chainName, "req", req)

	resp, err = handler(ctx, req)
	log.Debug("Finish handling", "resp", resp, "err", err)
	return resp, err
}

func (d *Dispatcher) preHandler(req interface{}) (resp *CommonReply) {
	chainName := req.(CommonRequest).GetChain()
	if _, ok := d.Registry[chainName]; !ok {
		return &CommonReply{
			Code:    common.ReturnCode_ERROR,
			Msg:     config.UnsupportedOperation,
			Support: false,
		}
	}
	return nil
}

func (d *Dispatcher) GetSupportChains(ctx context.Context, request *account.SupportChainsRequest) (*account.SupportChainsResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.SupportChainsResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  config.UnsupportedOperation,
		}, nil
	}
	return d.Registry[request.Chain].GetSupportChains(request)
}

func (d *Dispatcher) ConvertAddress(ctx context.Context, request *account.ConvertAddressRequest) (*account.ConvertAddressResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.ConvertAddressResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "covert address fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].ConvertAddress(request)
}

func (d *Dispatcher) ValidAddress(ctx context.Context, request *account.ValidAddressRequest) (*account.ValidAddressResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Dispatcher) GetBlockByNumber(ctx context.Context, request *account.BlockNumberRequest) (*account.BlockResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.BlockResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get block number fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].GetBlockByNumber(request)
}

func (d *Dispatcher) GetBlockByHash(ctx context.Context, request *account.BlockHashRequest) (*account.BlockResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.BlockResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get block by hash fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].GetBlockByHash(request)
}

func (d *Dispatcher) GetBlockHeaderByHash(ctx context.Context, request *account.BlockHeaderHashRequest) (*account.BlockHeaderResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.BlockHeaderResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get block header by hash fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].GetBlockHeaderByHash(request)
}

func (d *Dispatcher) GetBlockHeaderByNumber(ctx context.Context, request *account.BlockHeaderNumberRequest) (*account.BlockHeaderResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.BlockHeaderResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get block header by number fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].GetBlockHeaderByNumber(request)
}

func (d *Dispatcher) GetAccount(ctx context.Context, request *account.AccountRequest) (*account.AccountResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.AccountResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get account information fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].GetAccount(request)
}

func (d *Dispatcher) GetFee(ctx context.Context, request *account.FeeRequest) (*account.FeeResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.FeeResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get fee fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].GetFee(request)
}

func (d *Dispatcher) SendTx(ctx context.Context, request *account.SendTxRequest) (*account.SendTxResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.SendTxResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "send tx fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].SendTx(request)
}

func (d *Dispatcher) GetTxByAddress(ctx context.Context, request *account.TxAddressRequest) (*account.TxAddressResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.TxAddressResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get tx by address fail pre handle",
		}, nil
	}
	return d.Registry[request.Chain].GetTxByAddress(request)
}

func (d *Dispatcher) GetTxByHash(ctx context.Context, request *account.TxHashRequest) (*account.TxHashResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.TxHashResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get tx by hash fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].GetTxByHash(request)
}

func (d *Dispatcher) GetBlockByRange(ctx context.Context, request *account.BlockByRangeRequest) (*account.BlockByRangeResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.BlockByRangeResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get blcok by range fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].GetBlockByRange(request)
}

func (d *Dispatcher) CreateUnSignTransaction(ctx context.Context, request *account.UnSignTransactionRequest) (*account.UnSignTransactionResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.UnSignTransactionResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get un sign tx fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].CreateUnSignTransaction(request)
}

func (d *Dispatcher) BuildSignedTransaction(ctx context.Context, request *account.SignedTransactionRequest) (*account.SignedTransactionResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.SignedTransactionResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "signed tx fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].BuildSignedTransaction(request)
}

func (d *Dispatcher) DecodeTransaction(ctx context.Context, request *account.DecodeTransactionRequest) (*account.DecodeTransactionResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.DecodeTransactionResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "decode tx fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].DecodeTransaction(request)
}

func (d *Dispatcher) VerifySignedTransaction(ctx context.Context, request *account.VerifyTransactionRequest) (*account.VerifyTransactionResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.VerifyTransactionResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "verify tx fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].VerifySignedTransaction(request)
}

func (d *Dispatcher) GetExtraData(ctx context.Context, request *account.ExtraDataRequest) (*account.ExtraDataResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.ExtraDataResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get extra data fail at pre handle",
		}, nil
	}
	return d.Registry[request.Chain].GetExtraData(request)
}
