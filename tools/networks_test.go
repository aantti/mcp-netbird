package tools

import (
	"context"
	"testing"

	mcpnetbird "github.com/aantti/mcp-netbird"
)

func TestListNetbirdNetworks(t *testing.T) {
	// Create a context with API key
	ctx := mcpnetbird.WithNetbirdAPIKey(context.Background(), "test-token")

	// Call the function - we expect it to fail without a real API key
	networks, err := listNetbirdNetworks(ctx, ListNetbirdNetworksParams{})
	if err == nil {
		t.Error("expected error when calling API without valid token")
	}
	if networks != nil {
		t.Error("expected nil networks when error occurs")
	}

	// Test with empty context (no API key)
	networks, err = listNetbirdNetworks(context.Background(), ListNetbirdNetworksParams{})
	if err == nil {
		t.Error("expected error when calling API without token")
	}
	if networks != nil {
		t.Error("expected nil networks when error occurs")
	}
}
