package ucycle

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/kagwave/universal-cycle/types"
	"github.com/kagwave/universal-cycle/util"
)

// FindEulerianPathParallel finds an Eulerian path in the graph, which visits every edge exactly once, using a worker pool for concurrency.
func FindEulerianPathParallel(graph map[string][][]types.Symbol) ([][]types.Symbol, error) {
	var path [][]types.Symbol
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Empty slice to signal worker termination
	terminationSignal := []types.Symbol{}
	ch := make(chan []types.Symbol, 100)

	// Worker function to process nodes
	worker := func(id int) {
		defer wg.Done()
	forLoop:
		for {
			node, ok := <-ch
			if !ok {
				break forLoop // Channel closed, terminate
			}

			fmt.Printf("Worker %d received node: %v\n", id, node) // Debugging print

			if len(node) == 0 {
				mu.Unlock()
				break forLoop // Termination signal
			}

			nodeKey := util.PermToString(node)
			fmt.Printf("Worker %d processing node: %v\n", id, node)
			for {
				mu.Lock()
				if len(graph[nodeKey]) == 0 {
					mu.Unlock()
					break
				}
				// Sort the nodes to ensure a deterministic path
				sort.Slice(graph[nodeKey], func(i, j int) bool {
					return util.PermToString(graph[nodeKey][i]) < util.PermToString(graph[nodeKey][j])
				})

				// ... rest of the worker logic ...

			}

			mu.Lock()
			path = append(path, node)
			mu.Unlock()
			fmt.Printf("Worker %d appended node: %v to path\n", id, node)
		}
	}
	// Start worker pool
	numWorkers := 4 // Adjust this based on your CPU cores
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Handle empty graph case
	firstKey := util.FirstKey(graph)
	if firstKey == "" {
		return nil, errors.New("graph is empty")
	}

	// Seed the worker pool with the initial node
	initialNode := util.StringToPerm(firstKey)
	fmt.Printf("Initial node sent: %v\n", initialNode)
	ch <- initialNode

	// Signal worker termination (using empty slice)
	ch <- terminationSignal

	// Wait for all workers to finish
	wg.Wait()

	// Close the channel after workers finish
	close(ch)

	// Ensure the path is constructed correctly
	if len(path) == 0 {
		return nil, errors.New("could not complete Eulerian path")
	}

	fmt.Printf("Final path: %v\n", path)
	return util.ReversePath(path), nil
}
