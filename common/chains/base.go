package chains

import (
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
)

type TokenTransfer struct {
	Contract    string
	TokenType   basetype.TokenType
	TokenID     string
	From        string
	To          string
	Amount      uint64
	BlockNumber uint64
	TxHash      string
	BlockHash   string
}
