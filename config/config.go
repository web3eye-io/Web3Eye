package config

import (
	_ "embed"

	"github.com/BurntSushi/toml"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

//go:embed config.toml
var ConfigFS string

type Config struct {
	Version  string   `toml:"version" json:"version"`
	NFTMeta  NFTMeta  `toml:"nft-meta" json:"nft-meta"`
	BlockETL BlockETL `toml:"block-etl" json:"block-etl"`
	ETH      ETH      `toml:"eth" json:"eth"`
	IPFS     IPFS     `toml:"ipfs" json:"ipfs"`
	MySQL    MySQL    `toml:"mysql" json:"mysql"`
}

type NFTMeta struct {
	HTTPPort int    `toml:"http-port" json:"http-port"`
	GrpcPort int    `toml:"grpc-port" json:"grpc-port"`
	LogDir   string `toml:"log-dir" json:"log-dir"`
}

type BlockETL struct {
	HTTPPort int    `toml:"http-port" json:"http-port"`
	GrpcPort int    `toml:"grpc-port" json:"grpc-port"`
	LogDir   string `toml:"log-dir" json:"log-dir"`
}

type MySQL struct {
	IP       string `toml:"ip" json:"ip"`
	Port     int    `toml:"port" json:"port"`
	User     string `toml:"user" json:"user"`
	Password string `toml:"password" json:"password"`
}

type ETH struct {
	Wallets string `toml:"wallets" json:"wallets"`
}

type IPFS struct {
	HTTPGateway string `toml:"http-gateway" json:"http-gateway"`
}

// set default config
var config = &Config{
	NFTMeta: NFTMeta{
		HTTPPort: 30100,
		GrpcPort: 30101,
		LogDir:   "/var/log",
	},
}

func InitConfig() {
	md, err := toml.Decode(ConfigFS, config)
	if err != nil {
		logger.Sugar().Errorf("failed to parse config file, %v", err)
	}
	if len(md.Undecoded()) > 0 {
		logger.Sugar().Warnf("cannot parse [%v] to config", md.Undecoded())
	}
}

func GetConfig() *Config {
	return config
}
