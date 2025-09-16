# MCP server for Azure Nexus using the Go SDK

This is an implementation of a MCP server for Azure Nexus built using its [Go SDK](github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric). It exposes the following tools for interacting with Azure Nexus Resources:
![alt text](images/design.png)

- **Create Ipprefix**: Create Ipprefix with the required parameters.
![alt text](images/image.png)


> The project uses [mcp-go](https://github.com/mark3labs/mcp-go) as the MCP implementation.

## How to run

> Word(s) of caution: As much as I want folks to benefit from this, I have to call out that Large Language Models (LLMs) are non-deterministic by nature and can make mistakes. I would recommend you to **always validate** the results before making any decisions based on them.

```bash
git clone https://github.com/sachinDcoder/mcp_azure_nexus_go.git
cd mcp_azure_nexus_go

go build -o azure-nexus-mcp-server .
```

```
rajasachin@CPC-rajas-TQPGA:~/hackathon/azure-nexus-mcp-server$ go build -o azure-nexus-mcp-server .
rajasachin@CPC-rajas-TQPGA:~/hackathon/azure-nexus-mcp-server$ ./azure-nexus-mcp-server
Welcome to Azure Nexus MCP server!
Registering tools...

```

### Configure the MCP server

This will differ based on the MCP client/tool you use. For VS Code you can [follow these instructions](https://code.visualstudio.com/docs/copilot/chat/mcp-servers#_add-an-mcp-server) on how to configure this server using a `mcp.json` file.

Here is an example of the [mcp.json file](mcp.json):

```json
{
  "servers": {
    "Azure Nexus MCP (Golang)": {
      "type": "stdio",
      "command": "/home/rajasachin/hackathon/azure-nexus-mcp-server/azure-nexus-mcp-server",
    }
  }
}
```

Here is an example of Claude Desktop configuration:

```json
{
  "mcpServers": {
    "Azure Nexus MCP (Golang)": {
      "command": "enter path to binary e.g. /home/rajasachin/hackathon/azure-nexus-mcp-server/azure-nexus-mcp-server",
      "args": []
    }
    //other MCP servers...
  }
}
```

Here is an example of Cline:
```json
{
  "mcpServers": {
    "Azure Nexus MCP (Golang)": {
      "type": "stdio",
      "command": "/home/rajasachin/hackathon/azure-nexus-mcp-server/azure-nexus-mcp-server",
      "timeout": 1800
    }
  }
}
```
