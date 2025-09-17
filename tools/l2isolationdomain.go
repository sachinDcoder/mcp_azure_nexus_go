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

func CreateL2IsolationDomain(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return createL2IsolationDomain(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		name, ok := args["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("L2 isolation domain name missing")
		}

		location, ok := args["location"].(string)
		if !ok || location == "" {
			return nil, errors.New("location missing")
		}

		propertiesStr, ok := args["properties"].(string)
		if !ok || propertiesStr == "" {
			return nil, errors.New("properties missing")
		}

		var properties armmanagednetworkfabric.L2IsolationDomainProperties
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

		client, err := armmanagednetworkfabric.NewL2IsolationDomainsClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create L2 isolation domains client: %v", err)
		}

		poller, err := client.BeginCreate(ctx, resourceGroupName, name, armmanagednetworkfabric.L2IsolationDomain{
			Location:   &location,
			Properties: &properties,
		}, nil)

		if err != nil {
			return nil, fmt.Errorf("failed to begin creating L2 isolation domain: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create L2 isolation domain: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("L2 Isolation Domain '%s' created successfully in resource group '%s'", name, resourceGroupName)), nil
	}
}

func createL2IsolationDomain() mcp.Tool {
	return mcp.NewTool(
		CREATE_L2_ISOLATION_DOMAIN_TOOL_NAME,
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("location",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_LOCATION_DESCRIPTION),
		),
		mcp.WithString("properties",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_PROPERTIES_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Create a new L2 Isolation Domain"),
	)
}

func EnableL2IsolationDomain(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return enableL2IsolationDomain(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		name, ok := args["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("L2 isolation domain name missing")
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

		client, err := armmanagednetworkfabric.NewL2IsolationDomainsClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create L2 isolation domains client: %v", err)
		}

		state := armmanagednetworkfabric.EnableDisableState("Enable")
		poller, err := client.BeginUpdateAdministrativeState(ctx, resourceGroupName, name, armmanagednetworkfabric.UpdateAdministrativeState{
			State: &state,
		}, nil)

		if err != nil {
			return nil, fmt.Errorf("failed to begin enabling L2 isolation domain: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to enable L2 isolation domain: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("L2 Isolation Domain '%s' enabled successfully in resource group '%s'", name, resourceGroupName)), nil
	}
}

func enableL2IsolationDomain() mcp.Tool {
	return mcp.NewTool(
		ENABLE_L2_ISOLATION_DOMAIN_TOOL_NAME,
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Enable an L2 Isolation Domain"),
	)
}

func DisableL2IsolationDomain(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return disableL2IsolationDomain(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		name, ok := args["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("L2 isolation domain name missing")
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

		client, err := armmanagednetworkfabric.NewL2IsolationDomainsClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create L2 isolation domains client: %v", err)
		}

		state := armmanagednetworkfabric.EnableDisableState("Disable")
		poller, err := client.BeginUpdateAdministrativeState(ctx, resourceGroupName, name, armmanagednetworkfabric.UpdateAdministrativeState{
			State: &state,
		}, nil)

		if err != nil {
			return nil, fmt.Errorf("failed to begin disabling L2 isolation domain: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to disable L2 isolation domain: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("L2 Isolation Domain '%s' disabled successfully in resource group '%s'", name, resourceGroupName)), nil
	}
}

func disableL2IsolationDomain() mcp.Tool {
	return mcp.NewTool(
		DISABLE_L2_ISOLATION_DOMAIN_TOOL_NAME,
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Disable an L2 Isolation Domain"),
	)
}
