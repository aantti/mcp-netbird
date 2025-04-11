package tools

import (
	"context"

	mcpnetbird "github.com/aantti/mcp-netbird"
	"github.com/mark3labs/mcp-go/server"
)

type PortRangeMapping struct {
	TranslatedStart int    `json:"translated_start"`
	TranslatedEnd   int    `json:"translated_end"`
	IngressStart    int    `json:"ingress_start"`
	IngressEnd      int    `json:"ingress_end"`
	Protocol        string `json:"protocol"`
}

type NetbirdPortAllocations struct {
	ID                string             `json:"id"`
	Name              string             `json:"name"`
	IngressPeerID     string             `json:"ingress_peer_id"`
	Region            string             `json:"region"`
	Enabled           bool               `json:"enabled"`
	IngressIP         string             `json:"ingress_ip"`
	PortRangeMappings []PortRangeMapping `json:"port_range_mappings"`
}

type ListNetbirdPortAllocationsParams struct {
	// PeerID is the ID of the peer to get port allocations for
	// This field is required and must match a valid peer ID from NetbirdPeer in tools/peers.go
	PeerID string `mcp:"peer_id" validate:"required"`
}

func listNetbirdPortAllocations(ctx context.Context, args ListNetbirdPortAllocationsParams) ([]NetbirdPortAllocations, error) {
	client := mcpnetbird.NewNetbirdClient()

	var allocations []NetbirdPortAllocations
	if err := client.Get(ctx, "/peers/"+args.PeerID+"/ingress/ports", &allocations); err != nil {
		return nil, err
	}

	return allocations, nil
}

var ListNetbirdPortAllocations = mcpnetbird.MustTool(
	"list_netbird_port_allocations",
	"List all Netbird port allocations",
	listNetbirdPortAllocations,
)

func AddNetbirdPortAllocationTools(mcp *server.MCPServer) {
	ListNetbirdPortAllocations.Register(mcp)
}
