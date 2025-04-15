package tools

import (
	"context"

	mcpnetbird "github.com/aantti/mcp-netbird"
	"github.com/mark3labs/mcp-go/server"
)

type NetbirdNetwork struct {
	ID                string   `json:"id"`
	Routers           []string `json:"routers"`
	RoutingPeersCount int      `json:"routing_peers_count"`
	Resources         []string `json:"resources"`
	Policies          []string `json:"policies"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
}

type ListNetbirdNetworksParams struct{}

func listNetbirdNetworks(ctx context.Context, args ListNetbirdNetworksParams) ([]NetbirdNetwork, error) {
	var client *mcpnetbird.NetbirdClient
	if mcpnetbird.TestNetbirdClient != nil {
		client = mcpnetbird.TestNetbirdClient
	} else {
		client = mcpnetbird.NewNetbirdClient()
	}

	var networks []NetbirdNetwork
	if err := client.Get(ctx, "/networks", &networks); err != nil {
		return nil, err
	}

	return networks, nil
}

var ListNetbirdNetworks = mcpnetbird.MustTool(
	"list_netbird_networks",
	"List all Netbird networks",
	listNetbirdNetworks,
)

func AddNetbirdNetworkTools(mcp *server.MCPServer) {
	ListNetbirdNetworks.Register(mcp)
}
