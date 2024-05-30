package ucycle

import (
	"errors"

	"github.com/kagwave/universal-cycle/permutations"
	"github.com/kagwave/universal-cycle/types"
)

// Generate and evaluate permutations for optimization
func OptimizePermutations(symbols []types.Symbol, k int, evalFunc func([]types.Symbol) float64) ([]types.Symbol, float64, error) {
	subPerms := permutations.GenerateKPermutations(symbols, k)
	var bestPerm []types.Symbol
	var bestScore float64

	for _, perm := range subPerms {
		score := evalFunc(perm)
		if bestPerm == nil || score < bestScore {
			bestPerm = perm
			bestScore = score
		}
	}

	if bestPerm == nil {
		return nil, 0, errors.New("no valid permutations found")
	}

	return bestPerm, bestScore, nil
}

// Optimize task scheduling using universal cycles
func ScheduleTasks(tasks []types.Symbol, evalFunc func([]types.Symbol) float64) ([]types.Symbol, float64, error) {
	return OptimizePermutations(tasks, len(tasks), evalFunc)
}
