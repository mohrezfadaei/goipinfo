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
		IP       string `json:"query"`
		City     string `json:"city"`
		Region   string `json:"regionName"`
		Country  string `json:"country"`
		Location struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		}
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
		Location: fmt.Sprintf("%f,%f", info.Location.Lat, info.Location.Lon),
		Org:      info.Org,
		Timezone: info.Timezone,
	}, nil
}
