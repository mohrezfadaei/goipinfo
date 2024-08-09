package ip

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohrezfadaei/goipinfo/internal/config"
)

type IPAPIFetcher struct {
	config *config.Config
}

func NewIPAPIFetcher(cfg *config.Config) IPInfoFetcher {
	return &IPAPIFetcher{config: cfg}
}

func (f *IPAPIFetcher) Fetch(ip string) (*IPInfo, error) {
	url := fmt.Sprintf("%s/%s", f.config.IPAPIURL, ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch IP information: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	var info struct {
		Status      string  `json:"status"`
		Country     string  `json:"country"`
		CountryCode string  `json:"countryCode"`
		Region      string  `json:"region"`
		RegionName  string  `json:"regionName"`
		City        string  `json:"city"`
		Zip         string  `json:"zip"`
		Lat         float64 `json:"lat"`
		Lon         float64 `json:"lon"`
		Timezone    string  `json:"timezone"`
		Org         string  `json:"org"`
		As          string  `json:"as"`
		Query       string  `json:"query"` // IP
	}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &IPInfo{
		IP:       info.Query,
		City:     info.City,
		Region:   info.RegionName,
		Country:  info.Country,
		Location: fmt.Sprintf("%.4f,%.4f", info.Lat, info.Lon),
		Org:      info.Org,
		Timezone: info.Timezone,
	}, nil
}
