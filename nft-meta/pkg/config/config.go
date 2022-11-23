package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	KeyLogDir          = "logdir"
	KeyAppID           = "appid"
	KeyHTTPPort        = "http_port"
	KeyGRPCPort        = "grpc_port"
	KeyPrometheusPort  = "prometheus_port"
	KeySphinxProxyAddr = "sphinx_proxy_addr"
	KeyContract        = "contract"
	rootConfig         = "config"
)

func Init(configPath, appName string) error {
	viper.SetConfigName(fmt.Sprintf("%s.viper", appName))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath("./")
	viper.AddConfigPath(fmt.Sprintf("/etc/%v", appName))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%v", appName))
	viper.AddConfigPath(".")

	// Following're must for every service
	// config:
	//   hostname: my-service.npool.top
	//   http_port: 32759
	//   grpc_port: 32789
	//   prometheus_port: 32799
	//   appid: "89089012783789789719823798127398",
	//   logdir: "/var/log"
	//
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("fail to init config: %v", err)
	}

	appID, ok := viper.GetStringMap(rootConfig)[KeyAppID].(string)
	if !ok {
		return errors.New("fail to get init config KeyAppID not a string value")
	}

	logDir, ok := viper.GetStringMap(rootConfig)[KeyLogDir].(string)
	if !ok {
		return errors.New("fail to get init config KeyLogDir not a string value")
	}

	fmt.Printf("appid: %v\n", appID)
	fmt.Printf("logdir: %v\n", logDir)
	return nil
}

func GetString(key string) string {
	return viper.GetStringMap(rootConfig)[key].(string)
}

func GetInt(key string) int {
	return viper.GetStringMap(rootConfig)[key].(int)
}

var global *ENVInfo

type ENVInfo struct {
	LogDir string
}

func LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

func SetENV(info *ENVInfo) {
	global = info
}

func GetENV() *ENVInfo {
	return global
}
