package tools

import (
	"context"
	"time"

	mcpnetbird "github.com/aantti/mcp-netbird"
	"github.com/mark3labs/mcp-go/server"
)

type NetbirdGroup struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PeersCount     int    `json:"peers_count"`
	ResourcesCount int    `json:"resources_count"`
}

type NetbirdPeer struct {
	AccessiblePeersCount        int            `json:"accessible_peers_count"`
	ApprovalRequired            bool           `json:"approval_required"`
	CityName                    string         `json:"city_name"`
	Connected                   bool           `json:"connected"`
	ConnectionIP                string         `json:"connection_ip"`
	CountryCode                 string         `json:"country_code"`
	DNSLabel                    string         `json:"dns_label"`
	ExtraDNSLabels              []string       `json:"extra_dns_labels"`
	GeonameID                   int            `json:"geoname_id"`
	Groups                      []NetbirdGroup `json:"groups"`
	Hostname                    string         `json:"hostname"`
	ID                          string         `json:"id"`
	InactivityExpirationEnabled bool           `json:"inactivity_expiration_enabled"`
	IP                          string         `json:"ip"`
	KernelVersion               string         `json:"kernel_version"`
	LastLogin                   time.Time      `json:"last_login"`
	LastSeen                    time.Time      `json:"last_seen"`
	LoginExpirationEnabled      bool           `json:"login_expiration_enabled"`
	LoginExpired                bool           `json:"login_expired"`
	Name                        string         `json:"name"`
	OS                          string         `json:"os"`
	SerialNumber                string         `json:"serial_number"`
	SSHEnabled                  bool           `json:"ssh_enabled"`
	UIVersion                   string         `json:"ui_version"`
	UserID                      string         `json:"user_id"`
	Version                     string         `json:"version"`
}

type ListNetbirdPeersParams struct{}

func listNetbirdPeers(ctx context.Context, args ListNetbirdPeersParams) ([]NetbirdPeer, error) {
	client := mcpnetbird.NewNetbirdClient()

	var peers []NetbirdPeer
	if err := client.Get(ctx, "/peers", &peers); err != nil {
		return nil, err
	}

	return peers, nil
}

var ListNetbirdPeers = mcpnetbird.MustTool(
	"list_netbird_peers",
	"List all Netbird peers",
	listNetbirdPeers,
)

// AddNetbirdTools registers all Netbird tools with the MCP server
func AddNetbirdPeerTools(mcp *server.MCPServer) {
	ListNetbirdPeers.Register(mcp)
}
