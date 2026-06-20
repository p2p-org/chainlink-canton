package authentication

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// writeMetadata writes a minimal RFC 8414 metadata document for the given issuer.
func writeMetadata(w http.ResponseWriter, issuer string) {
	payload, err := json.Marshal(AuthorizationServerMetadata{
		Issuer:                        issuer,
		AuthorizationEndpoint:         issuer + "authorize",
		TokenEndpoint:                 issuer + "oauth/token",
		CodeChallengeMethodsSupported: []string{"S256"},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(payload)
}

// TestGetAuthorizationServerMetadata_TrailingSlash verifies that the issuer check tolerates a trailing
// slash mismatch between the configured base URL and the issuer reported by the server. Auth0 always
// reports the issuer with a trailing slash regardless of how the base URL was supplied.
func TestGetAuthorizationServerMetadata_TrailingSlash(t *testing.T) {
	t.Parallel()

	// The server reports its own address as the issuer, always with a trailing slash (like Auth0).
	mux := http.NewServeMux()
	mux.HandleFunc("/.well-known/oauth-authorization-server", func(w http.ResponseWriter, r *http.Request) {
		writeMetadata(w, "http://"+r.Host+"/")
	})
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	issuerWithSlash := server.URL + "/"

	tests := []struct {
		name    string
		baseURL string
	}{
		{name: "base url with trailing slash", baseURL: server.URL + "/"},
		{name: "base url without trailing slash", baseURL: server.URL},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			metadata, err := GetAuthorizationServerMetadata(t.Context(), tt.baseURL)
			require.NoError(t, err)
			require.Equal(t, issuerWithSlash, metadata.Issuer)
			require.Equal(t, issuerWithSlash+"oauth/token", metadata.TokenEndpoint)
		})
	}
}

// TestGetAuthorizationServerMetadata_IssuerMismatch verifies that a genuinely different issuer is still rejected.
func TestGetAuthorizationServerMetadata_IssuerMismatch(t *testing.T) {
	t.Parallel()

	mux := http.NewServeMux()
	mux.HandleFunc("/.well-known/oauth-authorization-server", func(w http.ResponseWriter, _ *http.Request) {
		writeMetadata(w, "https://attacker.example.com/")
	})
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	_, err := GetAuthorizationServerMetadata(t.Context(), server.URL)
	require.Error(t, err)
}
