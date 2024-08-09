package ip

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mohrezfadaei/goipinfo/internal/config"
)

func TestIPAPIFetcher_Fetch(t *testing.T) {
	mockResponse := `{
		"status": "success",
		"country": "Germany",
		"countryCode": "DE",
		"region": "NW",
		"regionName": "North Rhine-Westphalia",
		"city": "Düsseldorf",
		"zip": "40213",
		"lat": 51.2184,
		"lon": 6.7734,
		"timezone": "Europe/Berlin",
		"isp": "Cloudflare, Inc.",
		"org": "Cloudflare WARP",
		"as": "AS13335 Cloudflare, Inc.",
		"query": "104.28.242.72"
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	cfg := &config.Config{
		IPAPIURL: server.URL,
	}
	fetcher := NewIPAPIFetcher(cfg)

	info, err := fetcher.Fetch("104.28.242.72")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if info.IP != "104.28.242.72" {
		t.Errorf("expected IP '104.28.242.72', got '%s'", info.IP)
	}
	if info.City != "Düsseldorf" {
		t.Errorf("expected city 'Düsseldorf', got '%s'", info.City)
	}
	if info.Region != "North Rhine-Westphalia" {
		t.Errorf("expected region 'North Rhine-Westphalia', got '%s'", info.Region)
	}
	if info.Country != "Germany" {
		t.Errorf("expected country 'DE', got '%s'", info.Country)
	}
	if info.Location != "51.2184,6.7734" {
		t.Errorf("expected location '51.2184,6.7734', got '%s'", info.Location)
	}
	if info.Org != "Cloudflare WARP" {
		t.Errorf("expected org 'Cloudflare WARP', got '%s'", info.Org)
	}
	if info.Timezone != "Europe/Berlin" {
		t.Errorf("expected timezone 'Europe/Berlin', got '%s'", info.Timezone)
	}

}
