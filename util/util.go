package util

import (
	"strings"

	"github.com/kagwave/universal-cycle/types"
)

// PermToString converts a permutation of symbols to a string.
func PermToString(perm []types.Symbol) string {
	var sb strings.Builder
	for _, sym := range perm {
		sb.WriteString(sym.String())
	}
	return sb.String()
}

// RotateLeft rotates a slice of symbols to the left by n positions.
func RotateLeft(symbols []types.Symbol, n int) []types.Symbol {
	length := len(symbols)
	rotated := make([]types.Symbol, length)
	for i := 0; i < length; i++ {
		rotated[i] = symbols[(i+n)%length]
	}
	return rotated
}

// FirstKey returns the first key from a map.
func FirstKey(m map[string][][]types.Symbol) string {
	for k := range m {
		return k
	}
	return ""
}

// StringToPerm converts a string key back to a permutation.
func StringToPerm(s string) []types.Symbol {
	var perm []types.Symbol
	for _, r := range s {
		perm = append(perm, types.StringSymbol{Value: string(r)})
	}
	return perm
}

// ReversePath reverses the slice of permutations.
func ReversePath(path [][]types.Symbol) [][]types.Symbol {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

// Helper function to find the index of a symbol in the cycle
func IndexOf(cycle []types.Symbol, symbol types.Symbol) int {
	for i, sym := range cycle {
		if sym.Equals(symbol) {
			return i
		}
	}
	return -1
}
