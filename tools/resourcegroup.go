package tools

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func CreateResourceGroup(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return createResourceGroup(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		name, ok := args["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("resource group name missing")
		}

		location, ok := args["location"].(string)
		if !ok || location == "" {
			return nil, errors.New("location missing")
		}

		subscriptionId, ok := args["subscriptionId"].(string)
		if !ok || subscriptionId == "" {
			return nil, errors.New("subscription id missing")
		}

		cred, err := clientRetriever.Get()
		if err != nil {
			return nil, fmt.Errorf("error getting credentials: %v", err)
		}

		client, err := armresources.NewResourceGroupsClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create resource groups client: %v", err)
		}

		_, err = client.CreateOrUpdate(ctx, name, armresources.ResourceGroup{
			Location: &location,
		}, nil)

		if err != nil {
			return nil, fmt.Errorf("failed to create resource group: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("Resource Group '%s' created successfully in location '%s'", name, location)), nil
	}
}

func createResourceGroup() mcp.Tool {
	return mcp.NewTool(
		CREATE_RESOURCE_GROUP_TOOL_NAME,
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description(RESOURCE_GROUP_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("location",
			mcp.Required(),
			mcp.Description(RESOURCE_GROUP_LOCATION_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(RESOURCE_GROUP_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Create a new Resource Group"),
	)
}

func DeleteResourceGroup(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return deleteResourceGroup(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		name, ok := args["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("Resource Group name missing")
		}

		subscriptionId, ok := args["subscriptionId"].(string)
		if !ok || subscriptionId == "" {
			return nil, errors.New("subscription id missing")
		}

		cred, err := clientRetriever.Get()
		if err != nil {
			return nil, fmt.Errorf("error getting credentials: %v", err)
		}

		client, err := armresources.NewResourceGroupsClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create resource groups client: %v", err)
		}

		poller, err := client.BeginDelete(ctx, name, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to begin deleting resource group: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to delete resource group: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("Resource Group '%s' deleted successfully", name)), nil
	}
}

func deleteResourceGroup() mcp.Tool {
	return mcp.NewTool(
		DELETE_RESOURCE_GROUP_TOOL_NAME,
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description(RESOURCE_GROUP_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(RESOURCE_GROUP_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Delete a Resource Group"),
	)
}
