package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/web3eye-io/cyber-tracer/config"
)

func main() {
	fmt.Println(config.GetConfig().MySQL.IP)
	fmt.Println(config.GetConfig().MySQL.Port)
	fmt.Println(config.GetConfig().MySQL.Password)
	fmt.Println(config.GetConfig().BlockETL.HTTPPort)
}
