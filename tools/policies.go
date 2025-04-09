package tools

import (
	"context"

	mcpnetbird "github.com/aantti/mcp-netbird"
	"github.com/mark3labs/mcp-go/server"
)

type NetbirdPolicyRule struct {
	Action        string             `json:"action"`
	Bidirectional bool               `json:"bidirectional"`
	Description   string             `json:"description"`
	Destinations  []NetbirdPeerGroup `json:"destinations"`
	Enabled       bool               `json:"enabled"`
	ID            string             `json:"id"`
	Name          string             `json:"name"`
	Protocol      string             `json:"protocol"`
	Sources       []NetbirdPeerGroup `json:"sources"`
}

type NetbirdPolicy struct {
	Description         string              `json:"description"`
	Enabled             bool                `json:"enabled"`
	ID                  string              `json:"id"`
	Name                string              `json:"name"`
	Rules               []NetbirdPolicyRule `json:"rules"`
	SourcePostureChecks any                 `json:"source_posture_checks"`
}

type ListNetbirdPoliciesParams struct{}

func listNetbirdPolicies(ctx context.Context, args ListNetbirdPoliciesParams) ([]NetbirdPolicy, error) {
	client := mcpnetbird.NewNetbirdClient()

	var policies []NetbirdPolicy
	if err := client.Get(ctx, "/policies", &policies); err != nil {
		return nil, err
	}

	return policies, nil
}

var ListNetbirdPolicies = mcpnetbird.MustTool(
	"list_netbird_policies",
	"List all Netbird policies",
	listNetbirdPolicies,
)

func AddNetbirdPolicyTools(mcp *server.MCPServer) {
	ListNetbirdPolicies.Register(mcp)
}
