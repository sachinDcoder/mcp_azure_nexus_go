package tools

const (
	CREATE_IP_PREFIX_TOOL_NAME          = "create_ipprefix"
	IPREFIX_PARAMETER_DESCRIPTION       = "The name of the IP prefix to be created. If not available, ask the user to provide the name. Do not use a random name of your choice"
	IPREFIX_LOCATION_DESCRIPTION        = "The location of the IP prefix."
	IPREFIX_PROPERTIES_DESCRIPTION      = "The properties of the IP prefix, including IP prefix rules. This should be a JSON string."
	IPREFIX_IP_DESCRIPTION              = "The IP version(s) for the IP prefix, as a JSON string array e.g., [\"ipv6\"]."
	IPREFIX_RESOURCE_GROUP_DESCRIPTION  = "The name of the resource group."
	IPREFIX_SUBSCRIPTION_ID_DESCRIPTION = "The subscription ID for the Azure account."

	CREATE_IP_COMMUNITY_TOOL_NAME           = "create_ipcommunity"
	IPCOMMUNITY_PARAMETER_DESCRIPTION       = "The name of the IP community to be created."
	IPCOMMUNITY_LOCATION_DESCRIPTION        = "The location of the IP community."
	IPCOMMUNITY_PROPERTIES_DESCRIPTION      = "The properties of the IP community, including IP community rules. This should be a JSON string."
	IPCOMMUNITY_RESOURCE_GROUP_DESCRIPTION  = "The name of the resource group."
	IPCOMMUNITY_SUBSCRIPTION_ID_DESCRIPTION = "The subscription ID for the Azure account."

	CREATE_IP_EXT_COMMUNITY_TOOL_NAME          = "create_ipextcommunity"
	IPEXTCOMMUNITY_PARAMETER_DESCRIPTION       = "The name of the IP extended community to be created."
	IPEXTCOMMUNITY_LOCATION_DESCRIPTION        = "The location of the IP extended community."
	IPEXTCOMMUNITY_PROPERTIES_DESCRIPTION      = "The properties of the IP extended community, including IP extended community rules. This should be a JSON string."
	IPEXTCOMMUNITY_RESOURCE_GROUP_DESCRIPTION  = "The name of the resource group."
	IPEXTCOMMUNITY_SUBSCRIPTION_ID_DESCRIPTION = "The subscription ID for the Azure account."

	CREATE_ROUTE_POLICY_TOOL_NAME            = "create_routepolicy"
	ROUTE_POLICY_PARAMETER_DESCRIPTION       = "The name of the Route Policy to be created."
	ROUTE_POLICY_LOCATION_DESCRIPTION        = "The location of the Route Policy."
	ROUTE_POLICY_PROPERTIES_DESCRIPTION      = "The properties of the Route Policy, including statements. This should be a JSON string."
	ROUTE_POLICY_RESOURCE_GROUP_DESCRIPTION  = "The name of the resource group."
	ROUTE_POLICY_SUBSCRIPTION_ID_DESCRIPTION = "The subscription ID for the Azure account."

	CREATE_L3_ISOLATION_DOMAIN_TOOL_NAME            = "create_l3isolationdomain"
	L3_ISOLATION_DOMAIN_PARAMETER_DESCRIPTION       = "The name of the L3 Isolation Domain to be created."
	L3_ISOLATION_DOMAIN_LOCATION_DESCRIPTION        = "The location of the L3 Isolation Domain."
	L3_ISOLATION_DOMAIN_PROPERTIES_DESCRIPTION      = "The properties of the L3 Isolation Domain. This should be a JSON string."
	L3_ISOLATION_DOMAIN_RESOURCE_GROUP_DESCRIPTION  = "The name of the resource group."
	L3_ISOLATION_DOMAIN_SUBSCRIPTION_ID_DESCRIPTION = "The subscription ID for the Azure account."

	ENABLE_L3_ISOLATION_DOMAIN_TOOL_NAME = "enable_l3isolationdomain"

	CREATE_INTERNAL_NETWORK_TOOL_NAME            = "create_internalnetwork"
	INTERNAL_NETWORK_PARAMETER_DESCRIPTION       = "The name of the Internal Network to be created."
	INTERNAL_NETWORK_PROPERTIES_DESCRIPTION      = "The properties of the Internal Network. This should be a JSON string."
	INTERNAL_NETWORK_RESOURCE_GROUP_DESCRIPTION  = "The name of the resource group."
	INTERNAL_NETWORK_SUBSCRIPTION_ID_DESCRIPTION = "The subscription ID for the Azure account."

	DISABLE_L3_ISOLATION_DOMAIN_TOOL_NAME = "disable_l3isolationdomain"
)
