package encrypt

import (
	"testing"

	"github.com/kagwave/universal-cycle/types"
	"github.com/kagwave/universal-cycle/ucycle"
)

func TestGenerateKey(t *testing.T) {
	// Define a set of symbols
	symbols := []types.Symbol{
		types.StringSymbol{Value: "This"},
		types.StringSymbol{Value: "is"},
		types.StringSymbol{Value: "a"},
		types.StringSymbol{Value: "test"},
	}

	// Generate a universal cycle
	options := types.UCycleOptions{
		K:        2,
		Parallel: false,
	}
	fullCycle, err := ucycle.Create(symbols, options)
	if err != nil {
		t.Fatalf("Failed to create universal cycle: %v", err)
	}
	cycle := ucycle.Compress(fullCycle)

	// Generate a key of length 4
	key, err := GenerateKey(cycle, 4)
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}

	// Ensure the key is of the correct length
	if len(key) != 4 {
		t.Errorf("Expected key length of 4, got %d", len(key))
	}

	// Print the key
	t.Log("Generated Key:")
	for _, k := range key {
		t.Log(k.String())
	}
}

func TestEncryptDecryptData(t *testing.T) {
	// Define a set of symbols
	symbols := []types.Symbol{
		types.StringSymbol{Value: "Hello"},
		types.StringSymbol{Value: "world"},
		types.StringSymbol{Value: "this"},
		types.StringSymbol{Value: "is"},
		types.StringSymbol{Value: "a"},
		types.StringSymbol{Value: "test"},
	}

	// Generate a universal cycle
	options := types.UCycleOptions{
		K:        2,
		Parallel: false,
	}
	fullCycle, err := ucycle.Create(symbols, options)
	if err != nil {
		t.Fatalf("Failed to create universal cycle: %v", err)
	}
	cycle := ucycle.Compress(fullCycle)

	// Define data to encrypt and decrypt
	data := []types.Symbol{
		types.StringSymbol{Value: "Hello"},
		types.StringSymbol{Value: "world"},
		types.StringSymbol{Value: "this"},
		types.StringSymbol{Value: "is"},
	}

	// Generate a key of the same length as the data
	key, err := GenerateKey(cycle, len(data))
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}

	// Print the key
	t.Log("Generated Key:")
	for _, k := range key {
		t.Log(k.String())
	}

	// Print lengths for debugging
	t.Logf("Data length: %d, Key length: %d", len(data), len(key))

	// Encrypt the data
	encrypted, err := EncryptData(data, key)
	if err != nil {
		t.Fatalf("Failed to encrypt data: %v", err)
	}

	// Print the encrypted data
	t.Log("Encrypted Data:")
	for _, e := range encrypted {
		t.Log(e.String())
	}

	// Ensure encrypted data is not the same as the original
	for i, v := range encrypted {
		if v.Equals(data[i]) {
			t.Errorf("Expected encrypted data to differ from original at index %d", i)
		}
	}

	// Decrypt the data
	decrypted, err := DecryptData(encrypted, key)
	if err != nil {
		t.Fatalf("Failed to decrypt data: %v", err)
	}

	// Print the decrypted data
	t.Log("Decrypted Data:")
	for _, d := range decrypted {
		t.Log(d.String())
	}

	// Ensure decrypted data matches the original data
	for i, v := range decrypted {
		if !v.Equals(data[i]) {
			t.Errorf("Expected decrypted data to match original at index %d, got %v", i, v)
		}
	}
}
