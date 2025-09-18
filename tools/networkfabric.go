package tools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func CommitNetworkFabric(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return commitNetworkFabric(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		fabricName, ok := args["fabricName"].(string)
		if !ok || fabricName == "" {
			return nil, errors.New("fabric name missing")
		}

		resourceGroupName, ok := args["resourceGroupName"].(string)
		if !ok || resourceGroupName == "" {
			return nil, errors.New("resource group name missing")
		}

		subscriptionId, ok := args["subscriptionId"].(string)
		if !ok || subscriptionId == "" {
			return nil, errors.New("subscription id missing")
		}

		cred, err := clientRetriever.Get()
		if err != nil {
			return nil, fmt.Errorf("error getting credentials: %v", err)
		}

		client, err := armmanagednetworkfabric.NewNetworkFabricsClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create network fabrics client: %v", err)
		}

		poller, err := client.BeginCommitConfiguration(ctx, resourceGroupName, fabricName, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to begin commit configuration on network fabric: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to commit configuration on network fabric: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("Network Fabric '%s' configuration has been committed.", fabricName)), nil
	}
}

func commitNetworkFabric() mcp.Tool {
	return mcp.NewTool(
		COMMIT_NETWORK_FABRIC_TOOL_NAME,
		mcp.WithString("fabricName",
			mcp.Required(),
			mcp.Description(NETWORK_FABRIC_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(NETWORK_FABRIC_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(NETWORK_FABRIC_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Commits the configuration of the network fabric."),
	)
}

func GetNetworkFabric(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return getNetworkFabric(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		fabricName, ok := args["fabricName"].(string)
		if !ok || fabricName == "" {
			return nil, errors.New("fabric name missing")
		}

		resourceGroupName, ok := args["resourceGroupName"].(string)
		if !ok || resourceGroupName == "" {
			return nil, errors.New("resource group name missing")
		}

		subscriptionId, ok := args["subscriptionId"].(string)
		if !ok || subscriptionId == "" {
			return nil, errors.New("subscription id missing")
		}

		cred, err := clientRetriever.Get()
		if err != nil {
			return nil, fmt.Errorf("error getting credentials: %v", err)
		}

		client, err := armmanagednetworkfabric.NewNetworkFabricsClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create network fabrics client: %v", err)
		}

		fabric, err := client.Get(ctx, resourceGroupName, fabricName, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get network fabric: %v", err)
		}

		jsonResult, err := json.Marshal(fabric)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal network fabric result: %v", err)
		}
		return mcp.NewToolResultText(string(jsonResult)), nil
	}
}

func getNetworkFabric() mcp.Tool {
	return mcp.NewTool(
		GET_NETWORK_FABRIC_TOOL_NAME,
		mcp.WithString("fabricName",
			mcp.Required(),
			mcp.Description(NETWORK_FABRIC_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(NETWORK_FABRIC_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(NETWORK_FABRIC_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Gets the configuration of the network fabric."),
	)
}
