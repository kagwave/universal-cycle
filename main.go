package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kagwave/universal-cycle/types"
	"github.com/kagwave/universal-cycle/ucycle"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Prompt the user to enter symbols
		fmt.Print("Enter symbols (comma-separated), or 'q' to quit: ")
		symbolsInput, _ := reader.ReadString('\n')
		symbolsInput = strings.TrimSpace(symbolsInput)

		// Check if the user wants to quit
		if symbolsInput == "q" {
			fmt.Println("Quitting the program.")
			break
		}

		symbolsStr := strings.FieldsFunc(symbolsInput, func(r rune) bool {
			return r == ',' || r == ' ' || r == '\t'
		})

		// Convert the symbols to types.Symbol
		symbols := make([]types.Symbol, len(symbolsStr))
		for i, s := range symbolsStr {
			symbols[i] = types.StringSymbol{Value: s}
		}

		// Prompt the user to enter the value of k
		fmt.Print("Enter the value of k: ")
		var k int
		_, err := fmt.Scanf("%d", &k)
		if err != nil {
			fmt.Println("Invalid input for k. Please enter an integer.")
			continue
		}

		// Prompt the user to choose parallel or non-parallel execution
		fmt.Print("Run in parallel? (y/n): ")
		parallelInput, _ := reader.ReadString('\n')
		parallelInput = strings.TrimSpace(parallelInput)
		parallel := parallelInput == "y" || parallelInput == "Y"

		options := types.UCycleOptions{
			K:        k,
			Parallel: parallel,
		}

		// Generate the universal cycle
		start := time.Now()
		fullCycle, err := ucycle.Create(symbols, options)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Time taken to generate the universal cycle:", elapsed)
		fmt.Println("Universal Cycle:", fullCycle)

		// Convert to compressed
		compressedStart := time.Now()
		compressedCycle := ucycle.Compress(fullCycle)
		compressedElapsed := time.Since(compressedStart)

		fmt.Println("Time taken to convert to compressed:", compressedElapsed)
		fmt.Println("compressed:", compressedCycle)
	}
}
