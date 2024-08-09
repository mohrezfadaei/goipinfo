package ip

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mohrezfadaei/goipinfo/internal/config"
)

func TestMyIPFetcher_Fetcher(t *testing.T) {
	mockResponse := `{
		"success": true,
		"ip": "2a09:bac5:281b:126e::1d6:fb",
		"type": "IPv6",
		"country": {
		  	"code": "DE",
		  	"name": "Germany"
		},
		"region": "North Rhine-Westphalia",
		"city": "Düsseldorf",
		"location": {
			"lat": 51.2184,
		  	"lon": 6.7734
		},
		"timeZone": "Europe/Berlin",
		"asn": {
		  	"number": 13335,
		  	"name": "CLOUDFLARENET",
		  	"network": "2a09:bac5::/34"
		}
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
		MyIPAPIURL: server.URL,
	}
	fetcher := NewMyIPFetcher(cfg)
	info, err := fetcher.Fetch("")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if info.IP != "2a09:bac5:281b:126e::1d6:fb" {
		t.Errorf("expected IP '2a09:bac5:281b:126e::1d6:fb', got '%s'", info.IP)
	}
	if info.City != "Düsseldorf" {
		t.Errorf("expected city 'Düsseldorf', got '%s'", info.City)
	}
	if info.Region != "North Rhine-Westphalia" {
		t.Errorf("expected region 'North Rhine-Westphalia', got '%s'", info.Region)
	}
	if info.Country != "Germany" {
		t.Errorf("expected country 'Germany', got '%s'", info.Country)
	}
	if info.Location != "51.2184,6.7734" {
		t.Errorf("expected location '51.2184,6.7734', got %s", info.Location)
	}
	if info.Timezone != "Europe/Berlin" {
		t.Errorf("expected timezone 'Europe/Berlin', got '%s'", info.Timezone)
	}
	if info.Org != "CLOUDFLARENET" {
		t.Errorf("expected org 'CLOUDFLARENET', got '%s'", info.Org)
	}
}
