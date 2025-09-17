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

func CreateIPExtCommunity(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return createIPExtCommunity(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		name, ok := args["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("ip extended community name missing")
		}

		location, ok := args["location"].(string)
		if !ok || location == "" {
			return nil, errors.New("location missing")
		}

		propertiesStr, ok := args["properties"].(string)
		if !ok || propertiesStr == "" {
			return nil, errors.New("properties missing")
		}

		var properties armmanagednetworkfabric.IPExtendedCommunityProperties
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

		client, err := armmanagednetworkfabric.NewIPExtendedCommunitiesClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create ip extended communities client: %v", err)
		}

		poller, err := client.BeginCreate(ctx, resourceGroupName, name, armmanagednetworkfabric.IPExtendedCommunity{
			Location:   &location,
			Properties: &properties,
		}, nil)

		if err != nil {
			return nil, fmt.Errorf("failed to begin creating ip extended community: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create ip extended community: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("IP Extended Community '%s' created successfully in resource group '%s'", name, resourceGroupName)), nil
	}
}

func createIPExtCommunity() mcp.Tool {
	return mcp.NewTool(
		CREATE_IP_EXT_COMMUNITY_TOOL_NAME,
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("location",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_LOCATION_DESCRIPTION),
		),
		mcp.WithString("properties",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_PROPERTIES_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Create a new IP extended community"),
	)
}

func DeleteIPExtCommunity(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return deleteIPExtCommunity(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		name, ok := args["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("IP Extended Community name missing")
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

		client, err := armmanagednetworkfabric.NewIPExtendedCommunitiesClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create ip extended communities client: %v", err)
		}

		poller, err := client.BeginDelete(ctx, resourceGroupName, name, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to begin deleting ip extended community: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to delete ip extended community: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("IP Extended Community '%s' deleted successfully from resource group '%s'", name, resourceGroupName)), nil
	}
}

func deleteIPExtCommunity() mcp.Tool {
	return mcp.NewTool(
		DELETE_IP_EXT_COMMUNITY_TOOL_NAME,
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Delete an IP Extended Community"),
	)
}

func PatchIPExtCommunity(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return patchIPExtCommunity(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		name, ok := args["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("IP Extended Community name missing")
		}

		resourceGroupName, ok := args["resourceGroupName"].(string)
		if !ok || resourceGroupName == "" {
			return nil, errors.New("resource group name missing")
		}

		subscriptionId, ok := args["subscriptionId"].(string)
		if !ok || subscriptionId == "" {
			return nil, errors.New("subscription id missing")
		}

		propertiesStr, ok := args["properties"].(string)
		if !ok || propertiesStr == "" {
			return nil, errors.New("properties missing")
		}

		var properties armmanagednetworkfabric.IPExtendedCommunityPatch
		if err := json.Unmarshal([]byte(propertiesStr), &properties); err != nil {
			return nil, fmt.Errorf("error unmarshalling properties: %v", err)
		}

		cred, err := clientRetriever.Get()
		if err != nil {
			return nil, fmt.Errorf("error getting credentials: %v", err)
		}

		client, err := armmanagednetworkfabric.NewIPExtendedCommunitiesClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create ip extended communities client: %v", err)
		}

		poller, err := client.BeginUpdate(ctx, resourceGroupName, name, properties, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to begin updating ip extended community: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to update ip extended community: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("IP Extended Community '%s' updated successfully in resource group '%s'", name, resourceGroupName)), nil
	}
}

func patchIPExtCommunity() mcp.Tool {
	return mcp.NewTool(
		PATCH_IP_EXT_COMMUNITY_TOOL_NAME,
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("properties",
			mcp.Required(),
			mcp.Description("The properties to update on the IP Extended Community. This should be a JSON string."),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Patch an IP Extended Community"),
	)
}

func GetIPExtCommunity(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return getIPExtCommunity(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		name, ok := args["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("IP Extended Community name missing")
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

		client, err := armmanagednetworkfabric.NewIPExtendedCommunitiesClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create ip extended communities client: %v", err)
		}

		res, err := client.Get(ctx, resourceGroupName, name, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get ip extended community: %v", err)
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %v", err)
		}

		return mcp.NewToolResultText(string(resJson)), nil
	}
}

func getIPExtCommunity() mcp.Tool {
	return mcp.NewTool(
		GET_IP_EXT_COMMUNITY_TOOL_NAME,
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(IPEXTCOMMUNITY_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Get an IP Extended Community"),
	)
}
