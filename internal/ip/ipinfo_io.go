package ip

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohrezfadaei/goipinfo/internal/config"
)

type IPInfoIOFetcher struct {
	config *config.Config
}

func NewIPInfoIOFetcher(cfg *config.Config) IPInfoFetcher {
	return &IPInfoIOFetcher{config: cfg}
}

func (f *IPInfoIOFetcher) Fetch(ip string) (*IPInfo, error) {
	url := fmt.Sprintf("%s/%s", f.config.IPInfoAPIURL, ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch IP information: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	var info struct {
		IP       string `json:"ip"`
		City     string `json:"city"`
		Region   string `json:"region"`
		Country  string `json:"country"`
		Location string `json:"loc"`
		Org      string `json:"org"`
		Timezone string `json:"timezone"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &IPInfo{
		IP:       info.IP,
		City:     info.City,
		Region:   info.Region,
		Country:  info.Country,
		Location: info.Location,
		Org:      info.Org,
		Timezone: info.Timezone,
	}, nil
}
