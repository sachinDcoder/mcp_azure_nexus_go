package tools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func GetNetworkDevice(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return getNetworkDevice(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		deviceName, ok := args["deviceName"].(string)
		if !ok || deviceName == "" {
			return nil, errors.New("device name missing")
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

		client, err := armmanagednetworkfabric.NewNetworkDevicesClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create network devices client: %v", err)
		}

		device, err := client.Get(ctx, resourceGroupName, deviceName, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get network device: %v", err)
		}

		jsonResult, err := json.Marshal(device)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal network device result: %v", err)
		}

		return mcp.NewToolResultText(string(jsonResult)), nil
	}
}

func getNetworkDevice() mcp.Tool {
	return mcp.NewTool(
		GET_NETWORK_DEVICE_TOOL_NAME,
		mcp.WithString("deviceName",
			mcp.Required(),
			mcp.Description(NETWORK_DEVICE_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(NETWORK_DEVICE_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(NETWORK_DEVICE_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Gets the details of a network device."),
	)
}

func RebootNetworkDevice(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return rebootNetworkDevice(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
		}

		deviceName, ok := args["deviceName"].(string)
		if !ok || deviceName == "" {
			return nil, errors.New("device name missing")
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

		client, err := armmanagednetworkfabric.NewNetworkDevicesClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create network devices client: %v", err)
		}

		device, err := client.Get(ctx, resourceGroupName, deviceName, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get network device: %v", err)
		}

		if device.Properties.SerialNumber != nil && strings.Contains(*device.Properties.SerialNumber, "cEOSLab") {
			return mcp.NewToolResultText(fmt.Sprintf("Skipping reboot for vlab device '%s'.", deviceName)), nil
		}

		rebootProperties := armmanagednetworkfabric.RebootProperties{
			RebootType: to.Ptr(armmanagednetworkfabric.RebootType("Graceful")),
		}
		poller, err := client.BeginReboot(ctx, resourceGroupName, deviceName, rebootProperties, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to begin reboot on network device: %v", err)
		}

		_, err = poller.PollUntilDone(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to reboot network device: %v", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("Network Device '%s' rebooted successfully.", deviceName)), nil
	}
}

func rebootNetworkDevice() mcp.Tool {
	return mcp.NewTool(
		REBOOT_NETWORK_DEVICE_TOOL_NAME,
		mcp.WithString("deviceName",
			mcp.Required(),
			mcp.Description(NETWORK_DEVICE_PARAMETER_DESCRIPTION),
		),
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description(NETWORK_DEVICE_RESOURCE_GROUP_DESCRIPTION),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description(NETWORK_DEVICE_SUBSCRIPTION_ID_DESCRIPTION),
		),
		mcp.WithDescription("Reboots a network device."),
	)
}
