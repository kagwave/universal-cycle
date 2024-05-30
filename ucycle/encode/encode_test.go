package encode

import (
	"fmt"
	"testing"

	"github.com/kagwave/universal-cycle/types"
	"github.com/kagwave/universal-cycle/ucycle"
)

func TestEncodeDecode(t *testing.T) {
	// Define a specific set of symbols (e.g., related to a specific topic)
	symbols := []types.Symbol{
		types.StringSymbol{Value: "This"},
		types.StringSymbol{Value: "is"},
		types.StringSymbol{Value: "a"},
		types.StringSymbol{Value: "test."},
		types.StringSymbol{Value: "only"},
	}

	// Generate a universal cycle with k = 3 for more complexity
	options := types.UCycleOptions{
		K:        3,
		Parallel: false,
	}
	fullCycle, err := ucycle.Create(symbols, options)
	if err != nil {
		t.Fatalf("Failed to create universal cycle: %v", err)
	}
	cycle := ucycle.Compress(fullCycle)

	// Define a specific sentence to encode and decode
	data := []types.Symbol{
		types.StringSymbol{Value: "This"},
		types.StringSymbol{Value: "is"},
		types.StringSymbol{Value: "a"},
		types.StringSymbol{Value: "test."},
		types.StringSymbol{Value: "This"},
		types.StringSymbol{Value: "is"},
		types.StringSymbol{Value: "only"},
		types.StringSymbol{Value: "a"},
		types.StringSymbol{Value: "test."},
	}

	// Encode data
	encoded, err := EncodeData(data, cycle)
	if err != nil {
		t.Fatalf("Failed to encode data: %v", err)
	}

	// Print the encoded data
	fmt.Println("Encoded Data:", encoded)

	// Decode data
	decoded, err := DecodeData(encoded, cycle)
	if err != nil {
		t.Fatalf("Failed to decode data: %v", err)
	}

	// Print the decoded data
	fmt.Print("Decoded Data:")
	for _, v := range decoded {
		fmt.Printf(" %s", v.String())
	}
	fmt.Println()

	// Verify that decoded data matches original data
	for i, v := range decoded {
		if !v.Equals(data[i]) {
			t.Errorf("Expected %v at index %d, got %v", data[i], i, v)
		}
	}
}

func TestEncodeDataWithMixedSymbolTypes(t *testing.T) {
	// Define symbols including strings and integers
	symbols := []types.Symbol{
		types.StringSymbol{Value: "sensor"},
		types.StringSymbol{Value: "data"},
		types.IntSymbol{Value: 1},
		types.IntSymbol{Value: 2024},
		types.IntSymbol{Value: 42},
	}

	// Generate a universal cycle with k = 2
	options := types.UCycleOptions{
		K:        2,
		Parallel: false,
	}
	fullCycle, err := ucycle.Create(symbols, options)
	if err != nil {
		t.Fatalf("Failed to create universal cycle: %v", err)
	}
	cycle := ucycle.Compress(fullCycle)

	// Define data to encode and decode
	data := []types.Symbol{
		types.StringSymbol{Value: "sensor"},
		types.IntSymbol{Value: 2024},
		types.StringSymbol{Value: "data"},
		types.IntSymbol{Value: 42},
	}

	// Encode data
	encoded, err := EncodeData(data, cycle)
	if err != nil {
		t.Fatalf("Failed to encode data: %v", err)
	}

	// Print the encoded data
	fmt.Println("Encoded Data (Mixed Types):", encoded)

	// Decode data
	decoded, err := DecodeData(encoded, cycle)
	if err != nil {
		t.Fatalf("Failed to decode data: %v", err)
	}

	// Print the decoded data
	fmt.Print("Decoded Data (Mixed Types):")
	for _, v := range decoded {
		fmt.Printf(" %s", v.String())
	}
	fmt.Println()

	// Verify that decoded data matches original data
	for i, v := range decoded {
		if !v.Equals(data[i]) {
			t.Errorf("Expected %v at index %d, got %v", data[i], i, v)
		}
	}
}

func TestEncodeDataSymbolNotFound(t *testing.T) {
	// Define a cycle with fewer symbols
	cycle := []types.Symbol{
		types.StringSymbol{Value: "sensor"},
		types.StringSymbol{Value: "data"},
		types.StringSymbol{Value: "reading"},
	}

	// Data contains a symbol not in the cycle
	data := []types.Symbol{
		types.StringSymbol{Value: "sensor"},
		types.StringSymbol{Value: "value"}, // 'value' is not in cycle
	}

	_, err := EncodeData(data, cycle)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestDecodeDataInvalidPosition(t *testing.T) {
	// Define a cycle
	cycle := []types.Symbol{
		types.StringSymbol{Value: "sensor"},
		types.StringSymbol{Value: "data"},
		types.StringSymbol{Value: "reading"},
	}

	// Encoded data contains an invalid position
	encoded := []int{0, 1, 4} // 4 is an invalid position

	_, err := DecodeData(encoded, cycle)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestCycleIntegrity(t *testing.T) {
	// Define symbols
	symbols := []types.Symbol{
		types.StringSymbol{Value: "sensor"},
		types.StringSymbol{Value: "data"},
		types.StringSymbol{Value: "reading"},
		types.StringSymbol{Value: "value"},
		types.StringSymbol{Value: "measurement"},
	}

	// Generate a universal cycle with k = 2
	options := types.UCycleOptions{
		K:        2,
		Parallel: false,
	}
	fullCycle, err := ucycle.Create(symbols, options)
	if err != nil {
		t.Fatalf("Failed to create universal cycle: %v", err)
	}
	cycle := ucycle.Compress(fullCycle)

	// Ensure that the cycle contains all symbols
	for _, sym := range symbols {
		found := false
		for _, c := range cycle {
			if c.Equals(sym) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Symbol %v not found in the cycle", sym)
		}
	}
}
