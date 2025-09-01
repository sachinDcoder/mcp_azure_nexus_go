package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id)

	// Create MCP server
	s := server.NewMCPServer(
		"Azure Nexus MCP server 🚀",
		"0.0.1",
		server.WithLogging(),
	)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
