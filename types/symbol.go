package types

import (
	"fmt"
)

// Symbol is the interface for different types of symbols.
type Symbol interface {
	String() string
	Equals(Symbol) bool
	Xor(Symbol) Symbol
}

// StringSymbol is a symbol represented by a string.
type StringSymbol struct {
	Value string
}

// String returns the string representation of the symbol.
func (s StringSymbol) String() string {
	return s.Value
}

// Equals checks if two symbols are equal.
func (s StringSymbol) Equals(other Symbol) bool {
	if otherStr, ok := other.(StringSymbol); ok {
		return s.Value == otherStr.Value
	}
	return false
}

// Xor performs a bitwise XOR operation with another string symbol.
func (s StringSymbol) Xor(other Symbol) Symbol {
	if otherStr, ok := other.(StringSymbol); ok {
		result := make([]byte, len(s.Value))
		for i := range s.Value {
			result[i] = s.Value[i] ^ otherStr.Value[i]
		}
		return StringSymbol{Value: string(result)}
	}
	return s
}

// IntSymbol is a symbol represented by an integer.
type IntSymbol struct {
	Value int
}

// String returns the string representation of the symbol.
func (i IntSymbol) String() string {
	return fmt.Sprintf("%d", i.Value)
}

// Equals checks if two symbols are equal.
func (i IntSymbol) Equals(other Symbol) bool {
	if otherInt, ok := other.(IntSymbol); ok {
		return i.Value == otherInt.Value
	}
	return false
}

// Xor performs a bitwise XOR operation with another integer symbol.
func (i IntSymbol) Xor(other Symbol) Symbol {
	if otherInt, ok := other.(IntSymbol); ok {
		return IntSymbol{Value: i.Value ^ otherInt.Value}
	}
	return i
}
