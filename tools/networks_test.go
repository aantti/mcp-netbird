package tools

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mcpnetbird "github.com/aantti/mcp-netbird"
)

func TestListNetbirdNetworks(t *testing.T) {
	// Mock response data
	mockResp := []NetbirdNetwork{
		{
			ID:   "net1",
			Name: "Test Network",
			// Add other fields as needed for your struct
		},
	}

	// Create mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/networks" {
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
	networks, err := listNetbirdNetworks(ctx, ListNetbirdNetworksParams{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(networks) != 1 || networks[0].ID != "net1" {
		t.Errorf("unexpected result: %+v", networks)
	}
}
