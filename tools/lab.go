package tools

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type LabStatus struct {
	FabricStatus []FabricStatus `json:"fabricStatus"`
}

type FabricStatus struct {
	Name                string         `json:"name"`
	ProvisioningState   string         `json:"provisioningState"`
	AdministrativeState string         `json:"administrativeState"`
	ConfigurationState  string         `json:"configurationState"`
	DeviceStatus        []DeviceStatus `json:"deviceStatus"`
}

type DeviceStatus struct {
	Name                string `json:"name"`
	ProvisioningState   string `json:"provisioningState"`
	AdministrativeState string `json:"administrativeState"`
	ConfigurationState  string `json:"configurationState"`
}

func GetLabStatus(clientRetriever ServiceClientRetriever) (mcp.Tool, server.ToolHandlerFunc) {
	return getLabStatus(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return nil, errors.New("invalid arguments format")
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

		fabricsClient, err := armmanagednetworkfabric.NewNetworkFabricsClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create network fabrics client: %v", err)
		}

		devicesClient, err := armmanagednetworkfabric.NewNetworkDevicesClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create network devices client: %v", err)
		}

		var labStatus LabStatus
		fabricPager := fabricsClient.NewListByResourceGroupPager(resourceGroupName, nil)
		for fabricPager.More() {
			page, err := fabricPager.NextPage(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get next page of fabrics: %v", err)
			}
			for _, fabric := range page.Value {
				var fabricStatus FabricStatus
				fabricStatus.Name = *fabric.Name
				fabricStatus.ProvisioningState = string(*fabric.Properties.ProvisioningState)
				fabricStatus.AdministrativeState = string(*fabric.Properties.AdministrativeState)
				fabricStatus.ConfigurationState = string(*fabric.Properties.ConfigurationState)

				deviceIds, err := getDeviceIdsForFabric(ctx, clientRetriever, subscriptionId, resourceGroupName, *fabric.Name)
				if err != nil {
					return nil, fmt.Errorf("failed to get device IDs for fabric %s: %v", *fabric.Name, err)
				}

				for _, deviceId := range deviceIds {
					device, err := devicesClient.Get(ctx, resourceGroupName, deviceId, nil)
					if err != nil {
						return nil, fmt.Errorf("failed to get device %s: %v", deviceId, err)
					}
					var deviceStatus DeviceStatus
					deviceStatus.Name = *device.Name
					deviceStatus.ProvisioningState = string(*device.Properties.ProvisioningState)
					deviceStatus.AdministrativeState = string(*device.Properties.AdministrativeState)
					deviceStatus.ConfigurationState = string(*device.Properties.ConfigurationState)
					fabricStatus.DeviceStatus = append(fabricStatus.DeviceStatus, deviceStatus)
				}
				labStatus.FabricStatus = append(labStatus.FabricStatus, fabricStatus)
			}
		}

		isHealthy := true
		if len(labStatus.FabricStatus) == 0 {
			isHealthy = false
		}

		for _, fabric := range labStatus.FabricStatus {
			if fabric.ProvisioningState != "Succeeded" || fabric.AdministrativeState != "Enabled" || (fabric.ConfigurationState != "Succeeded" && fabric.ConfigurationState != "Provisioned") {
				isHealthy = false
				break
			}
			for _, device := range fabric.DeviceStatus {
				if device.ProvisioningState != "Succeeded" || device.AdministrativeState != "Enabled" || device.ConfigurationState != "Succeeded" {
					isHealthy = false
					break
				}
			}
			if !isHealthy {
				break
			}
		}

		var resultString string
		if isHealthy {
			resultString = "The lab is in a healthy state.\n\n"
		} else {
			resultString = "The lab is not in a healthy state.\n\n"
		}

		resultString += "Fabric Status:\n"
		for _, fabric := range labStatus.FabricStatus {
			resultString += fmt.Sprintf("- %s (Provisioning State: %s, Administrative State: %s, Configuration State: %s)\n", fabric.Name, fabric.ProvisioningState, fabric.AdministrativeState, fabric.ConfigurationState)
			resultString += "  Devices:\n"
			resultString += "    | Name | Provisioning State | AdministrativeState | Configuration State |\n"
			resultString += "    | :--- | :--- | :--- | :--- |\n"
			for _, device := range fabric.DeviceStatus {
				resultString += fmt.Sprintf("    | %s | %s | %s | %s |\n", device.Name, device.ProvisioningState, device.AdministrativeState, device.ConfigurationState)
			}
			resultString += "\n"
		}

		return mcp.NewToolResultText(resultString), nil
	}
}

func getLabStatus() mcp.Tool {
	return mcp.NewTool(
		GET_LAB_STATUS_TOOL_NAME,
		mcp.WithString("resourceGroupName",
			mcp.Required(),
			mcp.Description("The name of the resource group."),
		),
		mcp.WithString("subscriptionId",
			mcp.Required(),
			mcp.Description("The subscription ID for the Azure account."),
		),
		mcp.WithDescription("Gets the status of the lab, including network fabrics and devices."),
	)
}

func getDeviceIdsForFabric(ctx context.Context, clientRetriever ServiceClientRetriever, subscriptionId, resourceGroupName, fabricName string) ([]string, error) {
	cred, err := clientRetriever.Get()
	if err != nil {
		return nil, fmt.Errorf("error getting credentials: %v", err)
	}

	fabricsClient, err := armmanagednetworkfabric.NewNetworkFabricsClient(subscriptionId, cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create network fabrics client: %v", err)
	}

	fabric, err := fabricsClient.Get(ctx, resourceGroupName, fabricName, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get network fabric: %v", err)
	}

	var fabricDeviceIds []string
	if fabric.Properties.Racks != nil {
		racksClient, err := armmanagednetworkfabric.NewNetworkRacksClient(subscriptionId, cred, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create network racks client: %v", err)
		}
		for _, rackId := range fabric.Properties.Racks {
			rackName := getNameFromID(*rackId)
			rackResp, err := racksClient.Get(ctx, resourceGroupName, rackName, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to get network rack %s: %v", rackName, err)
			}
			// Check rack provisioning state for health check
			if *rackResp.NetworkRack.Properties.ProvisioningState != armmanagednetworkfabric.ProvisioningStateSucceeded {
				// We can return an error or handle it to mark the lab as unhealthy.
				// For now, we'll just not add its devices if the rack is not ready.
				continue
			}
			if rackResp.NetworkRack.Properties.NetworkDevices != nil {
				for _, deviceId := range rackResp.NetworkRack.Properties.NetworkDevices {
					fabricDeviceIds = append(fabricDeviceIds, getNameFromID(*deviceId))
				}
			}
		}
	}
	return fabricDeviceIds, nil
}
