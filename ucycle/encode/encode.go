package compress

import (
	"errors"

	"github.com/kagwave/universal-cycle/types"
	"github.com/kagwave/universal-cycle/util"
)

// Encode a data stream using the universal cycle
func EncodeData(data []types.Symbol, cycle []types.Symbol) ([]int, error) {
	encoded := make([]int, len(data))
	for i, d := range data {
		pos := util.IndexOf(cycle, d)
		if pos == -1 {
			return nil, errors.New("symbol not found in cycle")
		}
		encoded[i] = pos
	}
	return encoded, nil
}

// Decode a data stream from the universal cycle positions
func DecodeData(encoded []int, cycle []types.Symbol) ([]types.Symbol, error) {
	decoded := make([]types.Symbol, len(encoded))
	for i, pos := range encoded {
		if pos < 0 || pos >= len(cycle) {
			return nil, errors.New("invalid position in cycle")
		}
		decoded[i] = cycle[pos]
	}
	return decoded, nil
}
