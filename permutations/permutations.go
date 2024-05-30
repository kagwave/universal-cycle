package permutations

import (
	"sync"

	"github.com/kagwave/universal-cycle/types"
)

// GenerateKPermutations generates all sub-permutations of length k.
func GenerateKPermutations(symbols []types.Symbol, k int) [][]types.Symbol {
	if k == 1 {
		var singlePerms [][]types.Symbol
		for _, sym := range symbols {
			singlePerms = append(singlePerms, []types.Symbol{sym})
		}
		return singlePerms
	}

	var perms [][]types.Symbol
	for i, sym := range symbols {
		remaining := append([]types.Symbol{}, symbols[:i]...)
		remaining = append(remaining, symbols[i+1:]...)
		for _, perm := range GenerateKPermutations(remaining, k-1) {
			newPerm := append([]types.Symbol{sym}, perm...)
			perms = append(perms, newPerm)
		}
	}
	return perms
}

// GenerateSubPermutationsParallel generates all sub-permutations of length k in parallel.
func GenerateKPermutationsParallel(symbols []types.Symbol, k int) [][]types.Symbol {
	var perms [][]types.Symbol
	var wg sync.WaitGroup
	var mu sync.Mutex

	ch := make(chan []types.Symbol, 100)

	// Worker function to generate permutations
	worker := func(_ int) {
		defer wg.Done()
		for perm := range ch {
			mu.Lock()
			perms = append(perms, perm)
			mu.Unlock()
			// Log the worker ID and permutation for debugging
			//fmt.Printf("Worker %d processed permutation: %v\n", workerID, perm)
		}
	}

	// Start workers
	numWorkers := 4 // Adjust this based on your CPU cores
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Recursive function to generate permutations
	var generate func([]types.Symbol, []types.Symbol, int)
	generate = func(perm []types.Symbol, remaining []types.Symbol, k int) {
		if k == 0 {
			ch <- perm
			return
		}

		for i, sym := range remaining {
			newPerm := append([]types.Symbol{}, perm...)
			newPerm = append(newPerm, sym)
			newRemaining := append([]types.Symbol{}, remaining[:i]...)
			newRemaining = append(newRemaining, remaining[i+1:]...)
			generate(newPerm, newRemaining, k-1)
		}
	}

	go func() {
		defer close(ch)
		generate([]types.Symbol{}, symbols, k)
	}()

	wg.Wait()
	return perms
}

func GenerateAllPermutations(symbols []types.Symbol) [][]types.Symbol {
	var helper func([]types.Symbol, int)
	res := [][]types.Symbol{}

	helper = func(arr []types.Symbol, n int) {
		if n == 1 {
			tmp := make([]types.Symbol, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				} else {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				}
			}
		}
	}

	helper(symbols, len(symbols))
	return res
}
