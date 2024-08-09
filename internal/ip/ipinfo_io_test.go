package ip

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mohrezfadaei/goipinfo/internal/config"
)

func TestIPInfoIOFetcher(t *testing.T) {
	mockResponse := `{
  		"ip": "104.28.242.72",
  		"city": "Frankfurt am Main",
  		"region": "Hesse",
  		"country": "DE",
  		"loc": "50.1155,8.6842",
  		"org": "AS13335 Cloudflare, Inc.",
  		"postal": "60306",
  		"timezone": "Europe/Berlin",
  		"readme": "https://ipinfo.io/missingauth"
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		n, err := w.Write([]byte(mockResponse))
		if err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
		if n != len(mockResponse) {
			t.Fatalf("unexpected number of bytes written: expected %d, got %d", len(mockResponse), n)
		}
	}))
	defer server.Close()

	cfg := &config.Config{
		IPInfoAPIURL: server.URL,
	}
	fetcher := NewIPInfoIOFetcher(cfg)

	info, err := fetcher.Fetch("")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if info.IP != "104.28.242.72" {
		t.Errorf("expected IP to be '104.28.242.72', got '%s'", info.IP)
	}
	if info.City != "Frankfurt am Main" {
		t.Errorf("expected city to be 'Frankfurt am Main', got '%s'", info.City)
	}
	if info.Region != "Hesse" {
		t.Errorf("expected region to be 'Hesse', got '%s'", info.Region)
	}
	if info.Country != "DE" {
		t.Errorf("expected country to be 'DE', got '%s'", info.Country)
	}
	if info.Location != "50.1155,8.6842" {
		t.Errorf("expected location to be '50.1155,8.6842', got '%s'", info.Location)
	}
	if info.Org != "AS13335 Cloudflare, Inc." {
		t.Errorf("expected org to be 'AS13335 Cloudflare, Inc.', got '%s'", info.Org)
	}
	if info.Timezone != "Europe/Berlin" {
		t.Errorf("expected timezone to be 'Europe/Berlin', got '%s'", info.Timezone)
	}
}
