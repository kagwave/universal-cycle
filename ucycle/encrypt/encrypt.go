package encrypt

import (
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/kagwave/universal-cycle/types"
)

// Generate a cryptographic key using the universal cycle
func GenerateKey(cycle []types.Symbol, length int) ([]types.Symbol, error) {
	key := make([]types.Symbol, length)
	cycleLen := big.NewInt(int64(len(cycle)))

	for i := 0; i < length; i++ {
		randPos, err := rand.Int(rand.Reader, cycleLen)
		if err != nil {
			return nil, err
		}
		key[i] = cycle[randPos.Int64()]
	}

	return key, nil
}

// Encrypt data using the generated key
func EncryptData(data []types.Symbol, key []types.Symbol) ([]types.Symbol, error) {
	if len(data) != len(key) {
		return nil, errors.New("data and key must be of the same length")
	}

	encrypted := make([]types.Symbol, len(data))
	for i, d := range data {
		encrypted[i] = d.Xor(key[i])
	}
	return encrypted, nil
}

// Decrypt data using the generated key
func DecryptData(encrypted []types.Symbol, key []types.Symbol) ([]types.Symbol, error) {
	return EncryptData(encrypted, key) // Decryption is the same as encryption for XOR
}
