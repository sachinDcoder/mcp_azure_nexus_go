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

func CreateIPCommunity(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return createIPCommunity(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		name, ok := args["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("ip community name missing")
		}

		location, ok := args["location"].(string)
		if !ok || location == "" {
			return nil, errors.New("location missing")
		}

		propertiesStr, ok := args["properties"].(string)
		if !ok || propertiesStr == "" {
			return nil, errors.New("properties missing")
		}

		var properties armmanagednetworkfabric.IPCommunityProperties
		if err := json.Unmarshal([]byte(propertiesStr), &properties); err != nil {
			return nil, fmt.Errorf("error unmarshalling properties: %v", err)
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

		client, err := armmanagednetworkfabric.NewIPCommunitiesClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create ip communities client: %v", err)
		}

		poller, err := client.BeginCreate(ctx, resourceGroupName, name, armmanagednetworkfabric.IPCommunity{
			Location:   &location,
			Properties: &properties,
		}, nil)

		if err != nil {
			return nil, fmt.Errorf("failed to begin creating ip community: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create ip community: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("IP Community '%s' created successfully in resource group '%s'", name, resourceGroupName)), nil
	}
}

func createIPCommunity() mcp.Tool {
	return mcp.NewTool(
		CREATE_IP_COMMUNITY_TOOL_NAME,
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description(IPCOMMUNITY_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("location",
			mcp.Required(),
			mcp.Description(IPCOMMUNITY_LOCATION_DESCRIPTION),
		),
		mcp.WithString("properties",
			mcp.Required(),
			mcp.Description(IPCOMMUNITY_PROPERTIES_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(IPCOMMUNITY_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(IPCOMMUNITY_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Create a new IP community"),
	)
}

func DeleteIPCommunity(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return deleteIPCommunity(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		name, ok := args["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("IP Community name missing")
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

		client, err := armmanagednetworkfabric.NewIPCommunitiesClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create ip communities client: %v", err)
		}

		poller, err := client.BeginDelete(ctx, resourceGroupName, name, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to begin deleting ip community: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to delete ip community: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("IP Community '%s' deleted successfully from resource group '%s'", name, resourceGroupName)), nil
	}
}

func deleteIPCommunity() mcp.Tool {
	return mcp.NewTool(
		DELETE_IP_COMMUNITY_TOOL_NAME,
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description(IPCOMMUNITY_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(IPCOMMUNITY_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(IPCOMMUNITY_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Delete an IP Community"),
	)
}
