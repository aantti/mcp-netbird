package tools

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mcpnetbird "github.com/aantti/mcp-netbird"
)

func TestListNetbirdPolicies(t *testing.T) {
	// Mock response data
	mockResp := []NetbirdPolicy{
		{
			ID:   "policy1",
			Name: "Test Policy",
			// Add other fields as needed for your struct
		},
	}

	// Create mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/policies" {
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
	policies, err := listNetbirdPolicies(ctx, ListNetbirdPoliciesParams{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(policies) != 1 || policies[0].ID != "policy1" {
		t.Errorf("unexpected result: %+v", policies)
	}
}
