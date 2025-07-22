package opPool

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Operation struct {
	ExtensionID string         `json:"extension_id"`
	To          common.Address `json:"to"`
	Gas         uint           `json:"gas"`
	GasPrice    *big.Int       `json:"gas_price"`
	Data        []byte         `json:"data"`
	Sig         []byte         `json:"sig"`
	BlockHash   []byte         `json:"block_hash"` // block hash upon op submission to extension
}

type Pool struct {
}
