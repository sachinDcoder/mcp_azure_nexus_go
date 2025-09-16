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

func CreateInternalNetwork(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return createInternalNetwork(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		l3IsolationDomainName, ok := args["l3IsolationDomainName"].(string)
		if !ok || l3IsolationDomainName == "" {
			return nil, errors.New("L3 isolation domain name missing")
		}

		internalNetworkName, ok := args["internalNetworkName"].(string)
		if !ok || internalNetworkName == "" {
			return nil, errors.New("internal network name missing")
		}

		propertiesStr, ok := args["properties"].(string)
		if !ok || propertiesStr == "" {
			return nil, errors.New("properties missing")
		}

		var properties armmanagednetworkfabric.InternalNetworkProperties
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

		client, err := armmanagednetworkfabric.NewInternalNetworksClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create internal networks client: %v", err)
		}

		poller, err := client.BeginCreate(ctx, resourceGroupName, l3IsolationDomainName, internalNetworkName, armmanagednetworkfabric.InternalNetwork{
			Properties: &properties,
		}, nil)

		if err != nil {
			return nil, fmt.Errorf("failed to begin creating internal network: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create internal network: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("Internal Network '%s' created successfully in resource group '%s'", internalNetworkName, resourceGroupName)), nil
	}
}

func createInternalNetwork() mcp.Tool {
	return mcp.NewTool(
		CREATE_INTERNAL_NETWORK_TOOL_NAME,
		mcp.WithString("l3IsolationDomainName",
			mcp.Required(),
			mcp.Description(L3_ISOLATION_DOMAIN_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("internalNetworkName",
			mcp.Required(),
			mcp.Description(INTERNAL_NETWORK_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("properties",
			mcp.Required(),
			mcp.Description(INTERNAL_NETWORK_PROPERTIES_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(INTERNAL_NETWORK_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(INTERNAL_NETWORK_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Create a new Internal Network"),
	)
}
