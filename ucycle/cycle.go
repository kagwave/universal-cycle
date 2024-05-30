package ucycle

import (
	"errors"
	"sort"

	"github.com/kagwave/universal-cycle/permutations"
	"github.com/kagwave/universal-cycle/types"
	"github.com/kagwave/universal-cycle/util"
)

// BuildGraph constructs a directed graph from the sub-permutations.
func BuildGraph(perms [][]types.Symbol) map[string][][]types.Symbol {
	graph := make(map[string][][]types.Symbol)
	for _, perm := range perms {
		key := util.PermToString(perm[:len(perm)-1])
		value := perm[1:]
		graph[key] = append(graph[key], value)
	}
	return graph
}

// FindEulerianPath finds an Eulerian path in the graph using Hierholzer's algorithm.
func FindEulerianPath(graph map[string][][]types.Symbol) ([][]types.Symbol, error) {
	path := [][]types.Symbol{}
	stack := [][]types.Symbol{util.StringToPerm(util.FirstKey(graph))}

	// Perform DFS to find the Eulerian path
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		nodeKey := util.PermToString(node)
		if len(graph[nodeKey]) > 0 {
			// Sort the nodes to ensure a deterministic path
			sort.Slice(graph[nodeKey], func(i, j int) bool {
				return util.PermToString(graph[nodeKey][i]) < util.PermToString(graph[nodeKey][j])
			})
			nextNode := graph[nodeKey][0]
			graph[nodeKey] = graph[nodeKey][1:]
			stack = append(stack, nextNode)
		} else {
			path = append(path, node)
			stack = stack[:len(stack)-1]
		}
	}

	if len(path) == 0 {
		return nil, errors.New("could not complete Eulerian path")
	}

	return util.ReversePath(path), nil
}

// Compress converts the 2d array elements of the full u-cycle to one array representing the word.
func Compress(fullCycle [][]types.Symbol) []types.Symbol {
	if len(fullCycle) == 0 {
		return []types.Symbol{}
	}

	compressed := []types.Symbol{fullCycle[0][0]}

	for i := 1; i < len(fullCycle); i++ {
		if len(fullCycle[i]) > 1 && !fullCycle[i][0].Equals(compressed[len(compressed)-1]) {
			compressed = append(compressed, fullCycle[i][0])
		}
	}

	return compressed
}

func Create(symbols []types.Symbol, options types.UCycleOptions) ([][]types.Symbol, error) {
	if options.K < 1 || options.K > len(symbols) {
		return nil, errors.New("invalid permutation length")
	}

	var eulerianPath [][]types.Symbol
	var err error

	if options.Parallel {
		// Generate all sub-permutations of length k
		subPerms := permutations.GenerateKPermutations(symbols, options.K)

		// Build the graph from the sub-permutations
		graph := BuildGraph(subPerms)

		// Find the Eulerian path in the graph
		eulerianPath, err = FindEulerianPath(graph)
		if err != nil {
			return nil, err
		}
	} else {
		// Generate all sub-permutations of length k
		subPerms := permutations.GenerateKPermutations(symbols, options.K)

		// Build the graph from the sub-permutations
		graph := BuildGraph(subPerms)

		// Find the Eulerian path in the graph
		eulerianPath, err = FindEulerianPath(graph)
		if err != nil {
			return nil, err
		}
	}

	// Construct the full cycle
	fullCycle := [][]types.Symbol{}
	for i := 0; i < len(eulerianPath)-1; i++ {
		fullCycle = append(fullCycle, append(eulerianPath[i], eulerianPath[i+1][len(eulerianPath[i+1])-1]))
	}

	return fullCycle, nil
}
