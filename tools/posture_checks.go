package tools

import (
	"context"

	mcpnetbird "github.com/aantti/mcp-netbird"
	"github.com/mark3labs/mcp-go/server"
)

type VersionCheck struct {
	MinVersion string `json:"min_version,omitempty"`
}

type OSVersions struct {
	MinVersion       string `json:"min_version,omitempty"`
	MinKernelVersion string `json:"min_kernel_version,omitempty"`
}

type OSVersionCheck struct {
	Android *OSVersions `json:"android,omitempty"`
	IOS     *OSVersions `json:"ios,omitempty"`
	Darwin  *OSVersions `json:"darwin,omitempty"`
	Linux   *OSVersions `json:"linux,omitempty"`
	Windows *OSVersions `json:"windows,omitempty"`
}

type Location struct {
	CountryCode string `json:"country_code"`
	CityName    string `json:"city_name"`
}

type GeoLocationCheck struct {
	Locations []Location `json:"locations"`
	Action    string     `json:"action"`
}

type NetworkRangeCheck struct {
	Ranges []string `json:"ranges"`
	Action string   `json:"action"`
}

type ProcessPath struct {
	LinuxPath   string `json:"linux_path,omitempty"`
	MacPath     string `json:"mac_path,omitempty"`
	WindowsPath string `json:"windows_path,omitempty"`
}

type ProcessCheck struct {
	Processes []ProcessPath `json:"processes"`
}

type CheckConfig struct {
	NBVersionCheck    *VersionCheck      `json:"nb_version_check,omitempty"`
	OSVersionCheck    *OSVersionCheck    `json:"os_version_check,omitempty"`
	GeoLocationCheck  *GeoLocationCheck  `json:"geo_location_check,omitempty"`
	NetworkRangeCheck *NetworkRangeCheck `json:"peer_network_range_check,omitempty"`
	ProcessCheck      *ProcessCheck      `json:"process_check,omitempty"`
}

type NetbirdPostureCheck struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Checks      CheckConfig `json:"checks"`
}

type ListNetbirdPostureChecksParams struct{}

func listNetbirdPostureChecks(ctx context.Context, args ListNetbirdPostureChecksParams) ([]NetbirdPostureCheck, error) {
	client := mcpnetbird.NewNetbirdClient()

	var checks []NetbirdPostureCheck
	if err := client.Get(ctx, "/posture-checks", &checks); err != nil {
		return nil, err
	}

	return checks, nil
}

var ListNetbirdPostureChecks = mcpnetbird.MustTool(
	"list_netbird_posture_checks",
	"List all Netbird posture checks",
	listNetbirdPostureChecks,
)

func AddNetbirdPostureCheckTools(mcp *server.MCPServer) {
	ListNetbirdPostureChecks.Register(mcp)
}
