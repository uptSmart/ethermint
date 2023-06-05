package eip712

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

type Options struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Amino             *codec.LegacyAmino
}

// SetOptions set the encoding config to the singleton codecs (Amino and Protobuf).
// The process of unmarshaling SignDoc bytes into a SignDoc object requires having a codec
// populated with all relevant message types. As a result, we must call this method on app
// initialization with the app's encoding config.
func SetOptions(opts Options) {
	aminoCodec = opts.Amino
	protoCodec = codec.NewProtoCodec(opts.InterfaceRegistry)
}
