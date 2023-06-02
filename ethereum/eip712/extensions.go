package eip712

import (
	"math/big"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/cosmos/cosmos-sdk/codec"
	ethermint "github.com/evmos/ethermint/types"
)

type Options struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Amino             *codec.LegacyAmino
	// ChainIDBuilder for constructing a chainId that meets the requirements of ethermint
	ChainIDBuilder func(chainID string) string
}

var chainIDBuilder func(chainID string) string

// SetOptions set the encoding config to the singleton codecs (Amino and Protobuf).
// The process of unmarshaling SignDoc bytes into a SignDoc object requires having a codec
// populated with all relevant message types. As a result, we must call this method on app
// initialization with the app's encoding config.
func SetOptions(opts Options) {
	aminoCodec = opts.Amino
	protoCodec = codec.NewProtoCodec(opts.InterfaceRegistry)
	chainIDBuilder = opts.ChainIDBuilder
}

// ParseChainID override the default ethermint.ParseChainID
func ParseChainID(chainID string) (*big.Int, error) {
	if chainIDBuilder != nil {
		chainID = chainIDBuilder(chainID)
	}
	return ethermint.ParseChainID(chainID)
}
