package tools

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mcpnetbird "github.com/aantti/mcp-netbird"
)

func TestListNetbirdNameservers(t *testing.T) {
	// Mock response data
	mockResp := []NetbirdNameservers{
		{
			ID:   "ns1",
			Name: "Test Nameserver",
			Nameservers: []Nameserver{
				{IP: "1.2.3.4", NSType: "A", Port: 53},
			},
			Enabled: true,
		},
	}

	// Create mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/dns/nameservers" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(mockResp)
	}))
	defer server.Close()

	// Set the test client
	mcpnetbird.TestNetbirdClient = mcpnetbird.NewNetbirdClientWithBaseURL(server.URL)
	defer func() { mcpnetbird.TestNetbirdClient = nil }()

	ctx := mcpnetbird.WithNetbirdAPIKey(context.Background(), "test-token")
	nameservers, err := listNetbirdNameservers(ctx, ListNetbirdNameserversParams{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(nameservers) != 1 || nameservers[0].ID != "ns1" {
		t.Errorf("unexpected result: %+v", nameservers)
	}
}
