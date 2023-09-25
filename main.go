package main

import (
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
)

func main() {

	amountStr := decimal.NewFromBigInt(big.NewInt(1000254515021541), -int32(7)).Round(3).String()
	fmt.Println(amountStr)

	// sigchan := make(chan os.Signal, 1)
	// signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// <-sigchan
	// os.Exit(1)
}
