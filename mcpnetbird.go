package mcpnetbird

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mark3labs/mcp-go/server"
)

const (
	defaultNetbirdHost = "api.netbird.io"
	defaultNetbirdURL  = "https://" + defaultNetbirdHost
	netbirdAPIPath     = "/api"

	netbirdHostEnvVar = "NETBIRD_HOST"
	netbirdAPIEnvVar  = "NETBIRD_API_TOKEN"
)

// NetbirdClient provides methods to interact with the Netbird API
type NetbirdClient struct {
	baseURL string
	client  *http.Client
}

// Single global variable for testing
var TestNetbirdClient *NetbirdClient

// NewNetbirdClient creates a new NetbirdClient with the given API key
func NewNetbirdClient() *NetbirdClient {
	host := os.Getenv(netbirdHostEnvVar)
	if host == "" {
		host = defaultNetbirdHost
	}

	baseURL := "https://" + host + netbirdAPIPath
	return &NetbirdClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func NewNetbirdClientWithBaseURL(baseURL string) *NetbirdClient {
	return &NetbirdClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// do performs an HTTP request to the Netbird API
func (c *NetbirdClient) do(ctx context.Context, method, path string, body, v any) error {
	token := NetbirdAPIKeyFromContext(ctx)
	if token == "" {
		return fmt.Errorf("netbird API token not found in context")
	}

	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("decoding response: %w", err)
		}
	}

	return nil
}

// Get performs a GET request to the Netbird API
func (c *NetbirdClient) Get(ctx context.Context, path string, v any) error {
	return c.do(ctx, http.MethodGet, path, nil, v)
}

// Put performs a PUT request to the Netbird API
func (c *NetbirdClient) Put(ctx context.Context, path string, body, v any) error {
	return c.do(ctx, http.MethodPut, path, body, v)
}

type netbirdAPIKeyKey struct{}

// ExtractNetbirdInfoFromEnv is a StdioContextFunc that extracts Netbird configuration
// from environment variables and injects it into the context.
var ExtractNetbirdInfoFromEnv server.StdioContextFunc = func(ctx context.Context) context.Context {
	apiKey := os.Getenv(netbirdAPIEnvVar)
	if apiKey == "" {
		log.Printf("Warning: %s environment variable not found", netbirdAPIEnvVar)
	} else {
		log.Printf("Found %s environment variable", netbirdAPIEnvVar)
	}
	return WithNetbirdAPIKey(ctx, apiKey)
}

// ExtractNetbirdInfoFromEnvSSE is an SSEContextFunc that extracts Netbird configuration
// from environment variables and injects it into the context.
var ExtractNetbirdInfoFromEnvSSE server.SSEContextFunc = func(ctx context.Context, req *http.Request) context.Context {
	apiKey := os.Getenv(netbirdAPIEnvVar)
	if apiKey == "" {
		log.Printf("SSE MODE - Warning: %s environment variable not found", netbirdAPIEnvVar)
	} else {
		log.Printf("SSE MODE - Found %s environment variable with length %d", netbirdAPIEnvVar, len(apiKey))
	}
	return WithNetbirdAPIKey(ctx, apiKey)
}

// WithNetbirdAPIKey adds the Netbird API key to the context.
func WithNetbirdAPIKey(ctx context.Context, apiKey string) context.Context {
	return context.WithValue(ctx, netbirdAPIKeyKey{}, apiKey)
}

// NetbirdAPIKeyFromContext extracts the Netbird API key from the context.
func NetbirdAPIKeyFromContext(ctx context.Context) string {
	if v := ctx.Value(netbirdAPIKeyKey{}); v != nil {
		return v.(string)
	}
	return ""
}

// ComposeStdioContextFuncs composes multiple StdioContextFuncs into a single one.
func ComposeStdioContextFuncs(funcs ...server.StdioContextFunc) server.StdioContextFunc {
	return func(ctx context.Context) context.Context {
		for _, f := range funcs {
			ctx = f(ctx)
		}
		return ctx
	}
}

// ComposeSSEContextFuncs composes multiple SSEContextFuncs into a single one.
func ComposeSSEContextFuncs(funcs ...server.SSEContextFunc) server.SSEContextFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		for _, f := range funcs {
			ctx = f(ctx, req)
		}
		return ctx
	}
}

// ComposedStdioContextFunc is a StdioContextFunc that comprises all predefined StdioContextFuncs.
var ComposedStdioContextFunc = ComposeStdioContextFuncs(
	ExtractNetbirdInfoFromEnv,
)

// ComposedSSEContextFunc is an SSEContextFunc that comprises all predefined SSEContextFuncs.
var ComposedSSEContextFunc = ComposeSSEContextFuncs(
	ExtractNetbirdInfoFromEnvSSE,
)
