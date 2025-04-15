package mcpnetbird

import (
	"context"
	"testing"
)

func TestNetbirdAPIKeyContext(t *testing.T) {
	ctx := context.Background()
	apiKey := "test-api-key"

	// Test adding API key to context
	ctxWithKey := WithNetbirdAPIKey(ctx, apiKey)
	if got := NetbirdAPIKeyFromContext(ctxWithKey); got != apiKey {
		t.Errorf("NetbirdAPIKeyFromContext() = %v, want %v", got, apiKey)
	}

	// Test getting API key from context without key
	if got := NetbirdAPIKeyFromContext(ctx); got != "" {
		t.Errorf("NetbirdAPIKeyFromContext() = %v, want empty string", got)
	}
}
