package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/log"
	"github.com/zhulida1234/mix-chain-account/chaindispatcher"
	"github.com/zhulida1234/mix-chain-account/config"
	wallet2 "github.com/zhulida1234/mix-chain-account/rpc/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	var f = flag.String("c", "config.yml", "config path")
	flag.Parse()
	conf, err := config.NewConfig(*f)
	if err != nil {
		panic(err)
	}
	dispatcher, err := chaindispatcher.NewDispatcher(conf)
	if err != nil {
		log.Error("Setup dispatcher failed", "err", err)
		panic(err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(dispatcher.Interceptor))
	defer grpcServer.GracefulStop()

	// 需要注册grpc protobuf
	wallet2.RegisterWalletAccountServiceServer(grpcServer, dispatcher)

	listen, err := net.Listen("tcp", ":"+conf.Server.Port)
	if err != nil {
		log.Error("Listen failed", "err", err)
		panic(err)
	}

	reflection.Register(grpcServer)

	log.Info("mix wallet rpc services start success", "port", conf.Server.Port)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Error("Serve failed", "err", err)
		panic(err)
	}
}
