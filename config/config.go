package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/log"
)

type Server struct {
	Port string `yaml:"port"`
}

type RPC struct {
	RpcUrl string `yaml:"rpc_url"`
}

type Node struct {
	RPCs         []*RPC `yaml:"rpcs"`
	DataApiUrl   string `yaml:"data_api_url"`
	DataApiKey   string `yaml:"data_api_key"`
	DataApiToken string `yaml:"data_api_token"`
	Timeout      int    `yaml:"timeout"`
}

type WalletNode struct {
	Eth    Node `yaml:"eth"`
	Solana Node `yaml:"solana"`
}

type Config struct {
	Server     Server     `yaml:"server"`
	RPC        RPC        `yaml:"rpc"`
	Network    string     `yaml:"network"`
	WalletNode WalletNode `yaml:"wallet_node"`
	Chains     []string   `yaml:"chains"`
}

func NewConfig(path string) (*Config, error) {
	var config = new(Config)
	h := log.NewTerminalHandler(os.Stdout, true)
	log.SetDefault(log.NewLogger(h))

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

const UnsupportedChain = "Unsupport chain"
const UnsupportedOperation = UnsupportedChain
