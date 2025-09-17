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

func GetL2IsolationDomain(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return getL2IsolationDomain(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		res, err := client.Get(ctx, resourceGroupName, name, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get L2 isolation domain: %v", err)
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %v", err)
		}

		return mcp.NewToolResultText(string(resJson)), nil
	}
}

func getL2IsolationDomain() mcp.Tool {
	return mcp.NewTool(
		GET_L2_ISOLATION_DOMAIN_TOOL_NAME,
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
		mcp.WithDescription("Get an L2 Isolation Domain"),
	)
}

func GetL2IsolationDomainAdministrativeState(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return getL2IsolationDomainAdministrativeState(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		res, err := client.Get(ctx, resourceGroupName, name, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get L2 isolation domain: %v", err)
		}

		if res.Properties == nil || res.Properties.AdministrativeState == nil {
			return nil, fmt.Errorf("administrative state not found for L2 isolation domain '%s'", name)
		}

		return mcp.NewToolResultText(string(*res.Properties.AdministrativeState)), nil
	}
}

func getL2IsolationDomainAdministrativeState() mcp.Tool {
	return mcp.NewTool(
		GET_L2_ISOLATION_DOMAIN_ADMINISTRATIVE_STATE_TOOL_NAME,
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
		mcp.WithDescription("Get the administrative state of an L2 Isolation Domain"),
	)
}

func GetL2IsolationDomainConfigurationState(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return getL2IsolationDomainConfigurationState(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		res, err := client.Get(ctx, resourceGroupName, name, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get L2 isolation domain: %v", err)
		}

		if res.Properties == nil || res.Properties.ConfigurationState == nil {
			return nil, fmt.Errorf("configuration state not found for L2 isolation domain '%s'", name)
		}

		return mcp.NewToolResultText(string(*res.Properties.ConfigurationState)), nil
	}
}

func getL2IsolationDomainConfigurationState() mcp.Tool {
	return mcp.NewTool(
		GET_L2_ISOLATION_DOMAIN_CONFIGURATION_STATE_TOOL_NAME,
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
		mcp.WithDescription("Get the configuration state of an L2 Isolation Domain"),
	)
}

func DeleteL2IsolationDomain(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return deleteL2IsolationDomain(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		poller, err := client.BeginDelete(ctx, resourceGroupName, name, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to begin deleting L2 isolation domain: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to delete L2 isolation domain: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("L2 Isolation Domain '%s' deleted successfully from resource group '%s'", name, resourceGroupName)), nil
	}
}

func deleteL2IsolationDomain() mcp.Tool {
	return mcp.NewTool(
		DELETE_L2_ISOLATION_DOMAIN_TOOL_NAME,
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
		mcp.WithDescription("Delete an L2 Isolation Domain"),
	)
}

func PatchL2IsolationDomain(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return patchL2IsolationDomain(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		name, ok := args["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("L2 Isolation Domain name missing")
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

		var patchProps armmanagednetworkfabric.L2IsolationDomainPatchProperties
		if err := json.Unmarshal([]byte(propertiesStr), &patchProps); err != nil {
			return nil, fmt.Errorf("error unmarshalling properties: %v", err)
		}
		properties := armmanagednetworkfabric.L2IsolationDomainPatch{
			Properties: &patchProps,
		}

		cred, err := clientRetriever.Get()
		if err != nil {
			return nil, fmt.Errorf("error getting credentials: %v", err)
		}

		client, err := armmanagednetworkfabric.NewL2IsolationDomainsClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create l2 isolation domains client: %v", err)
		}

		poller, err := client.BeginUpdate(ctx, resourceGroupName, name, properties, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to begin updating l2 isolation domain: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to update l2 isolation domain: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("L2 Isolation Domain '%s' updated successfully in resource group '%s'", name, resourceGroupName)), nil
	}
}

func patchL2IsolationDomain() mcp.Tool {
	return mcp.NewTool(
		PATCH_L2_ISOLATION_DOMAIN_TOOL_NAME,
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("properties",
			mcp.Required(),
			mcp.Description("The properties to update on the L2 Isolation Domain. This should be a JSON string."),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(L2_ISOLATION_DOMAIN_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Patch an L2 Isolation Domain"),
	)
}
