package ucycle

import (
	"testing"
	"time"

	"github.com/kagwave/universal-cycle/types"
)

func TestCreate(t *testing.T) {
	// Define a set of symbols
	symbols := []types.Symbol{
		types.StringSymbol{Value: "α"},
		types.StringSymbol{Value: "β"},
		types.StringSymbol{Value: "γ"},
		types.StringSymbol{Value: "δ"},
		types.StringSymbol{Value: "ε"},
		types.StringSymbol{Value: "ζ"},
	}

	// Set options for non-parallel execution
	options := types.UCycleOptions{
		K:        5,
		Parallel: false,
	}

	start := time.Now()
	fullCycle, err := Create(symbols, options)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("Error generating the universal cycle: %v", err)
	}
	t.Logf("Universal Cycle: %v", fullCycle)
	t.Logf("Time taken to generate the universal cycle: %v", elapsed)

	compressedStart := time.Now()
	compressedCycle := Compress(fullCycle)
	compressedElapsed := time.Since(compressedStart)

	t.Logf("compressed: %v", compressedCycle)
	t.Logf("Time taken to convert to compressed: %v", compressedElapsed)
}

func TestCreateParallel(t *testing.T) {
	// Define a set of symbols
	symbols := []types.Symbol{
		types.StringSymbol{Value: "α"},
		types.StringSymbol{Value: "β"},
		types.StringSymbol{Value: "γ"},
		types.StringSymbol{Value: "δ"},
		types.StringSymbol{Value: "ε"},
		types.StringSymbol{Value: "ζ"},
	}

	// Set options for parallel execution
	options := types.UCycleOptions{
		K:        5,
		Parallel: true,
	}

	start := time.Now()
	fullCycle, err := Create(symbols, options)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("Error generating the universal cycle (parallel): %v", err)
	}
	t.Logf("Universal Cycle: %v", fullCycle)
	t.Logf("Time taken to generate the universal cycle (parallel): %v", elapsed)

	compressedStart := time.Now()
	compressedCycle := Compress(fullCycle)
	compressedElapsed := time.Since(compressedStart)

	t.Logf("compressed: %v", compressedCycle)
	t.Logf("Time taken to convert to compressed (parallel): %v", compressedElapsed)
}
