package types

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"

	errorsmod "cosmossdk.io/errors"
)

var (
	regexChainID         = `[a-z]{1,}`
	regexEIP155Separator = `_{1}`
	regexEIP155          = `[1-9][0-9]*`
	regexEpochSeparator  = `-{1}`
	regexEpoch           = `[1-9][0-9]*`
	ethermintChainID     = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)%s(%s)$`,
		regexChainID,
		regexEIP155Separator,
		regexEIP155,
		regexEpochSeparator,
		regexEpoch))

	chainIDBuilder func(chainID string) string
)

// SetChainIDBuilder return a chainId that meets the requirements of ethermint
func SetChainIDBuilder(builder func(chainID string) string) {
	chainIDBuilder = builder
}

// IsValidChainID returns false if the given chain identifier is incorrectly formatted.
func IsValidChainID(chainID string) bool {
	if len(chainID) > 48 {
		return false
	}

	return ethermintChainID.MatchString(chainID)
}

// ParseChainID parses a string chain identifier's epoch to an Ethereum-compatible
// chain-id in *big.Int format. The function returns an error if the chain-id has an invalid format
func ParseChainID(chainID string) (*big.Int, error) {
	chainID = strings.TrimSpace(chainID)
	if len(chainID) > 48 {
		return nil, errorsmod.Wrapf(
			ErrInvalidChainID,
			"chain-id '%s' cannot exceed 48 chars",
			chainID,
		)
	}

	chainID = buildChainID(chainID)

	matches := ethermintChainID.FindStringSubmatch(chainID)
	if matches == nil || len(matches) != 4 || matches[1] == "" {
		return nil, errorsmod.Wrapf(ErrInvalidChainID, "%s: %v", chainID, matches)
	}

	// verify that the chain-id entered is a base 10 integer
	chainIDInt, ok := new(big.Int).SetString(matches[2], 10)
	if !ok {
		return nil, errorsmod.Wrapf(
			ErrInvalidChainID,
			"epoch %s must be base-10 integer format",
			matches[2],
		)
	}

	return chainIDInt, nil
}

func buildChainID(chainID string) string {
	if chainIDBuilder == nil || IsValidChainID(chainID) {
		return chainID
	}
	return chainIDBuilder(chainID)
}
