package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fibonacci Sequence Demo (YAML Config)")
	fmt.Println("=====================================")
	
	// Load YAML configuration
	config, err := loadConfigFromYAML("config-fibonacci.yaml")
	if err != nil {
		fmt.Printf("Error loading config-fibonacci.yaml: %v\n", err)
		return
	}

	fmt.Printf("Loaded %d nodes from configuration\n", len(config.Nodes))

	// Create chain builder
	builder := NewChainBuilder(config)
	
	// Register fibonacci node types
	RegisterFibonacciNodeTypes(builder)

	fmt.Println("\nExecuting fibonacci sequence with reconcile loop...")
	builder.Execute()

	fmt.Println("\n📊 Final Node States:")
	states := builder.GetAllNodeStates()
	for name, state := range states {
		fmt.Printf("  %s: %s (processed: %d)\n",
			name, state.Status, state.Stats.TotalProcessed)
	}

	fmt.Println("\n✅ Fibonacci sequence execution completed!")
}