package tools

import (
	"context"

	mcpnetbird "github.com/aantti/mcp-netbird"
	"github.com/mark3labs/mcp-go/server"
)

type Nameserver struct {
	IP     string `json:"ip"`
	NSType string `json:"ns_type"`
	Port   int    `json:"port"`
}

type NetbirdNameservers struct {
	ID                   string       `json:"id"`
	Name                 string       `json:"name"`
	Description          string       `json:"description"`
	Nameservers          []Nameserver `json:"nameservers"`
	Enabled              bool         `json:"enabled"`
	Groups               []string     `json:"groups"`
	Primary              bool         `json:"primary"`
	Domains              []string     `json:"domains"`
	SearchDomainsEnabled bool         `json:"search_domains_enabled"`
}

type ListNetbirdNameserversParams struct{}

func listNetbirdNameservers(ctx context.Context, args ListNetbirdNameserversParams) ([]NetbirdNameservers, error) {
	var client *mcpnetbird.NetbirdClient
	if mcpnetbird.TestNetbirdClient != nil {
		client = mcpnetbird.TestNetbirdClient
	} else {
		client = mcpnetbird.NewNetbirdClient()
	}

	var nameservers []NetbirdNameservers
	if err := client.Get(ctx, "/dns/nameservers", &nameservers); err != nil {
		return nil, err
	}

	return nameservers, nil
}

var ListNetbirdNameservers = mcpnetbird.MustTool(
	"list_netbird_nameservers",
	"List all Netbird nameservers",
	listNetbirdNameservers,
)

func AddNetbirdNameserverTools(mcp *server.MCPServer) {
	ListNetbirdNameservers.Register(mcp)
}
