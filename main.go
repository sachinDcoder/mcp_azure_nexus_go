package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	"github.com/sachinDcoder/mcp_azure_nexus_go/tools"
)

func main() {
	fmt.Println("Welcome to Azure Nexus MCP server!")

	// Create MCP server
	s := server.NewMCPServer(
		"Azure Nexus MCP server 🚀",
		"0.0.1",
		server.WithLogging(),
	)

	fmt.Println("Registering tools...")

	s.AddTool(tools.CreateResourceGroup(tools.ServiceClientRetriever{}))
	s.AddTool(tools.DeleteResourceGroup(tools.ServiceClientRetriever{}))
	s.AddTool(tools.CreateIPPrefix(tools.ServiceClientRetriever{}))
	s.AddTool(tools.CreateIPCommunity(tools.ServiceClientRetriever{}))
	s.AddTool(tools.CreateIPExtCommunity(tools.ServiceClientRetriever{}))
	s.AddTool(tools.CreateRoutePolicy(tools.ServiceClientRetriever{}))
	s.AddTool(tools.CreateL3IsolationDomain(tools.ServiceClientRetriever{}))
	s.AddTool(tools.EnableL3IsolationDomain(tools.ServiceClientRetriever{}))
	s.AddTool(tools.CreateInternalNetwork(tools.ServiceClientRetriever{}))
	s.AddTool(tools.DisableL3IsolationDomain(tools.ServiceClientRetriever{}))
	s.AddTool(tools.CreateL2IsolationDomain(tools.ServiceClientRetriever{}))
	s.AddTool(tools.EnableL2IsolationDomain(tools.ServiceClientRetriever{}))
	s.AddTool(tools.DisableL2IsolationDomain(tools.ServiceClientRetriever{}))
	s.AddTool(tools.CreateExternalNetwork(tools.ServiceClientRetriever{}))
	s.AddTool(tools.GetL2IsolationDomain(tools.ServiceClientRetriever{}))
	s.AddTool(tools.GetL2IsolationDomainAdministrativeState(tools.ServiceClientRetriever{}))
	s.AddTool(tools.GetL2IsolationDomainConfigurationState(tools.ServiceClientRetriever{}))
	s.AddTool(tools.GetL3IsolationDomain(tools.ServiceClientRetriever{}))
	s.AddTool(tools.GetL3IsolationDomainAdministrativeState(tools.ServiceClientRetriever{}))
	s.AddTool(tools.GetL3IsolationDomainConfigurationState(tools.ServiceClientRetriever{}))
	s.AddTool(tools.DeleteL2IsolationDomain(tools.ServiceClientRetriever{}))
	s.AddTool(tools.DeleteL3IsolationDomain(tools.ServiceClientRetriever{}))
	s.AddTool(tools.DeleteIPPrefix(tools.ServiceClientRetriever{}))
	s.AddTool(tools.DeleteIPCommunity(tools.ServiceClientRetriever{}))
	s.AddTool(tools.DeleteIPExtCommunity(tools.ServiceClientRetriever{}))
	s.AddTool(tools.DeleteRoutePolicy(tools.ServiceClientRetriever{}))

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
