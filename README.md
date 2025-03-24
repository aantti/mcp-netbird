# Netbird MCP Server

A [Model Context Protocol](https://modelcontextprotocol.io) (MCP) server for [Netbird](https://netbird.io/).

This project is derived from the [MCP Server for Grafana](https://github.com/grafana/mcp-grafana) by Grafana Labs and is licensed under the same Apache License 2.0.

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/aantti/mcp-netbird
cd mcp-netbird

# Build and install
make install
```

### From GitHub

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

## Usage

1. Get your Netbird API token from the Netbird management console.

2. Install the `mcp-netbird` binary using one of the installation methods above.

3. Add the server configuration to your client configuration file. For example, for Claude Desktop:

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

> Note: if you see `Error: spawn mcp-netbird ENOENT` in Claude Desktop, you need to specify the full path to `mcp-netbird`.

## Development

Contributions are welcome! Please open an issue or submit a pull request if you have any suggestions or improvements.

This project is written in Go. Install Go following the instructions for your platform.

To run the server manually, use:

```bash
NETBIRD_API_TOKEN=your-token make run
```

Or in SSE mode:

```bash
NETBIRD_API_TOKEN=your-token make run-sse
```

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
