package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"
)

//go:embed config.toml
var ConfigFS string

type Config struct {
	Version    string     `toml:"version" env:"version"`
	NFTMeta    NFTMeta    `toml:"nft-meta" env:"nft_meta"`
	BlockETL   BlockETL   `toml:"block-etl" env:"block_etl"`
	CloudProxy CloudProxy `toml:"cloud-proxy" env:"cloud_proxy"`
	Gateway    Gateway    `toml:"gateway" env:"gateway"`
	Ranker     Ranker     `toml:"ranker" env:"ranker"`
	Entrance   Entrance   `toml:"entrance" env:"entrance"`
	GenCar     GenCar     `toml:"gen-car" env:"gen_car"`
	Dealer     Dealer     `toml:"dealer" env:"dealer"`
	Retriever  Retriever  `toml:"retriever" env:"retriever"`
	Transform  Transform  `toml:"transform" env:"transform"`
	IPFS       IPFS       `toml:"ipfs" env:"ipfs"`
	MySQL      MySQL      `toml:"mysql" env:"mysql"`
	Pulsar     Pulsar     `toml:"pulsar" env:"pulsar"`
	Redis      Redis      `toml:"redis" env:"redis"`
	Milvus     Milvus     `toml:"milvus" env:"milvus"`
	Minio      Minio      `toml:"minio" env:"minio"`
}

type NFTMeta struct {
	Domain         string `toml:"domain" env:"domain"`
	HTTPPort       int    `toml:"http-port" env:"http_port"`
	GrpcPort       int    `toml:"grpc-port" env:"grpc_port"`
	LogFile        string `toml:"log-file" env:"log_file"`
	CollectionName string `toml:"collection-name" env:"collection_name"`
	Description    string `toml:"description" env:"description"`
}

type BlockETL struct {
	Domain   string `toml:"domain" env:"domain"`
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogFile  string `toml:"log-file" env:"log_file"`
}

type CloudProxy struct {
	Domain   string `toml:"domain" env:"domain"`
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogFile  string `toml:"log-file" env:"log_file"`
}

type Gateway struct {
	Domain   string `toml:"domain" env:"domain"`
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogFile  string `toml:"log-file" env:"log_file"`
}

type Ranker struct {
	Domain   string `toml:"domain" env:"domain"`
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogFile  string `toml:"log-file" env:"log_file"`
}

type Entrance struct {
	Domain   string `toml:"domain" env:"domain"`
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogFile  string `toml:"log-file" env:"log_file"`
}

type GenCar struct {
	Domain   string `toml:"domain" env:"domain"`
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogFile  string `toml:"log-file" env:"log_file"`
	DataDir  string `toml:"data-dir" env:"data_dir"`
}

type Dealer struct {
	Domain    string `toml:"domain" env:"domain"`
	IpfsRepo  string `toml:"ipfs-repo" env:"ipfs_repo"`
	OrbitRepo string `toml:"orbit-repo" env:"orbit_repo"`
	HTTPPort  int    `toml:"http-port" env:"http_port"`
	GrpcPort  int    `toml:"grpc-port" env:"grpc_port"`
	LogFile   string `toml:"log-file" env:"log_file"`
}

type Retriever struct {
	Domain    string `toml:"domain" env:"domain"`
	IpfsRepo  string `toml:"ipfs-repo" env:"ipfs_repo"`
	OrbitRepo string `toml:"orbit-repo" env:"orbit_repo"`
	HTTPPort  int    `toml:"http-port" env:"http_port"`
	GrpcPort  int    `toml:"grpc-port" env:"grpc_port"`
	LogFile   string `toml:"log-file" env:"log_file"`
}

type Transform struct {
	Domain   string `toml:"domain" env:"domain"`
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogFile  string `toml:"log-file" env:"log_file"`
	DataDir  string `toml:"data-dir" env:"data_dir"`
	PyDir    string `toml:"py-dir" env:"py_dir"`
}
type MySQL struct {
	Domain   string `toml:"domain" env:"domain"`
	Port     int    `toml:"port" env:"port"`
	User     string `toml:"user" env:"user"`
	Password string `toml:"password" env:"password"`
	Database string `toml:"database" env:"database"`
}

type Redis struct {
	Address  string `toml:"address" env:"address"`
	Password string `toml:"password" env:"password"`
}

type Pulsar struct {
	Domain              string `toml:"domain" env:"domain"`
	Port                int    `toml:"port" env:"port"`
	OperationTimeout    uint64 `toml:"operation-timeout" env:"operation_timeout"`
	ConnectionTimeout   uint64 `toml:"connection-timeout" env:"connection_timeout"`
	TopicSyncTask       string `toml:"topic-sync-task" env:"topic_sync_task"`
	TopicTransformImage string `toml:"topic-transform-image" env:"topic_transform_image"`
}

type Milvus struct {
	Address string `toml:"address" env:"address"`
}

type Minio struct {
	Address          string `toml:"address" env:"address"`
	AccessKey        string `toml:"access-key" env:"access_key"`
	SecretKey        string `toml:"secret-key" env:"secret_key"`
	Region           string `toml:"region" env:"region"`
	TokenImageBucket string `toml:"token-image-bucket" env:"token_image_bucket"`
}

type IPFS struct {
	HTTPGateway string `toml:"http-gateway" env:"http_gateway"`
}

// set default config
var config = &Config{
	NFTMeta: NFTMeta{
		HTTPPort: 30100,
		GrpcPort: 30101,
	},
}

type envMatcher struct {
	envMap map[string]string
}

func DetectEnv(co *Config) (err error) {
	e := &envMatcher{}
	e.envMap = make(map[string]string)
	ct := reflect.TypeOf(co)
	e.detectEnv(ct, "", "")
	_, err = toml.Decode(e.toToml(), co)
	return err
}

// read environment var
func (e *envMatcher) detectEnv(t reflect.Type, preffix, _preffix string) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		it := t.Field(i)
		envKey := fmt.Sprintf("%v%v", preffix, it.Tag.Get("env"))
		_envKey := fmt.Sprintf("%v%v", _preffix, it.Tag.Get("toml"))
		if it.Type.Kind() != reflect.Struct {
			if envValue, ok := os.LookupEnv(envKey); ok {
				if it.Type.Kind() == reflect.String {
					e.envMap[_envKey] = fmt.Sprintf("\"%v\"", envValue)
				} else {
					e.envMap[_envKey] = envValue
				}
			}
			continue
		}
		envKey = fmt.Sprintf("%v%v_", preffix, it.Tag.Get("env"))
		_envKey = fmt.Sprintf("%v%v.", _preffix, it.Tag.Get("toml"))
		e.detectEnv(it.Type, envKey, _envKey)
	}
}

func (e *envMatcher) toToml() string {
	var b bytes.Buffer

	for v := range e.envMap {
		b.WriteString(fmt.Sprintf("%v=%v\n", v, e.envMap[v]))
	}

	return b.String()
}

func init() {
	md, err := toml.Decode(ConfigFS, config)
	if err != nil {
		panic(fmt.Sprintf("failed to parse config file, %v", err))
	}
	if len(md.Undecoded()) > 0 {
		fmt.Printf("cannot parse [%v] to config\n", md.Undecoded())
	}
	err = DetectEnv(config)
	if err != nil {
		fmt.Printf("environment variable parse failed, %v", err)
	}
}

func GetConfig() *Config {
	return config
}
