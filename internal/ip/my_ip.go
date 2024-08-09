package ip

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohrezfadaei/goipinfo/internal/config"
)

type MyIPFetcher struct {
	config *config.Config
}

func NewMyIPFetcher(cfg *config.Config) IPInfoFetcher {
	return &MyIPFetcher{config: cfg}
}

func (f *MyIPFetcher) Fetch(ip string) (*IPInfo, error) {
	url := fmt.Sprintf("%s/%s", f.config.MyIPAPIURL, ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch IP information: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	var info struct {
		IP      string `json:"ip"`
		City    string `json:"city"`
		Region  string `json:"region"`
		Country struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"country"`
		Location struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"location"`
		Org struct {
			Number  string `json:"number"`
			Name    string `json:"name"`
			Network string `json:"network"`
		} `json:"asn"`
		Timezone string `json:"timeZone"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &IPInfo{
		IP:       info.IP,
		City:     info.City,
		Region:   info.Region,
		Country:  info.Country.Name,
		Location: fmt.Sprintf("%f,%f", info.Location.Lat, info.Location.Lon),
		Org:      info.Org.Name,
		Timezone: info.Timezone,
	}, nil
}