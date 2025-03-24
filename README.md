# Netbird MCP Server

A [Model Context Protocol](https://modelcontextprotocol.io) (MCP) server for [Netbird](https://netbird.io/).

This project is derived from the [MCP Server for Grafana](https://github.com/grafana/mcp-grafana) by Grafana Labs and is licensed under the same Apache License 2.0.

**Note: this project is still in development.**

## Installing from source

### Clone the repository

```bash
git clone https://github.com/aantti/mcp-netbird
```

### Build and install

```bash
cd mcp-netbird && \
make install
```

## Installing from GitHub

```bash
go install github.com/aantti/mcp-netbird/cmd/mcp-netbird@latest
```

## Configuration

The server requires the following environment variables:

- `NETBIRD_API_TOKEN`: Your Netbird API token
- `NETBIRD_HOST` (optional): The Netbird API host (default is `api.netbird.io`)

## Features

- [x] List Netbird peers with detailed information
  - Connected status
  - Location information
  - System details
  - Group membership
- [x] Configurable API endpoint
- [x] Secure token-based authentication

### Tools

| Tool | Category | Description |
| --- | --- | --- |
| `list_netbird_peers` | Peers | List all peers in your Netbird network |
| `list_netbird_groups` | Groups | List all groups in your Netbird network |
| `list_netbird_policies` | Policies | List all policies in your Netbird network |

## Usage

1. Get your [Netbird API token](https://docs.netbird.io/api/guides/authentication) from the Netbird management console.

2. Install the `mcp-netbird` binary using one of the installation methods above. Make sure the binary is in your PATH.

3. Add the server configuration to your client configuration file. E.g., for Codeium Windsurf add the following to `~/.codeium/windsurf/mcp_config.json`:

   ```json
   {
     "mcpServers": {
       "netbird": {
         "command": "mcp-netbird",
         "args": [],
         "env": {
           "NETBIRD_API_TOKEN": "<your-api-token>"
         }
       }
     }
   }
   ```

For more information on how to add a similar configuration to Claude Desktop, see [here](https://modelcontextprotocol.io/quickstart/user).

> Note: if you see something along the lines of `[netbird] [error] spawn mcp-netbird ENOENT` in Claude Desktop logs, you need to specify the full path to `mcp-netbird`. On macOS Claude Logs are in `~/Library/Logs/Claude`.

4. Try asking questions along the lines of "Can you explain my Netbird peers, groups and policies to me?"

## Development

Contributions are welcome! Please open an issue or submit a pull request if you have any suggestions or improvements.

This project is written in Go. Install Go following the instructions for your platform.

To run the server manually, use:

```bash
export NETBIRD_API_TOKEN=your-token && \
go run cmd/mcp-netbird/main.go
```

Or in SSE mode:

```bash
export NETBIRD_API_TOKEN=your-token && \
go run cmd/mcp-netbird/main.go --transport sse --sse-address :8001
```

### Debugging

The **MCP Inspector** is an interactive developer tool for testing and debugging MCP servers. Read more about it [here](https://modelcontextprotocol.io/docs/tools/inspector).

Here's how to start the MCP Inspector:

```bash
export NETBIRD_API_TOKEN=your-token && \
npx @modelcontextprotocol/inspector
```

Netbird MCP Server can then be tested with either `stdio` or `SSE` transport type.

### Testing

**TODO: add tests**

### Linting

To lint the code, run:

```bash
make lint
```

## License

This project is licensed under the [Apache License, Version 2.0](LICENSE).

This project includes software developed at Grafana Labs (https://grafana.com/).
