package config

type Config struct {
	IPInfoAPIURL string // ipinfo.io
	IPAPIURL     string // ip-api.com
	MyIPAPIURL   string // api.my-ip.io
}

func LoadConfig() (*Config, error) {
	return &Config{
		IPInfoAPIURL: "http://ipinfo.io/json",
		IPAPIURL:     "http://ip-api.com/json",
		MyIPAPIURL:   "https://api.my-ip.io/v2/ip.json",
	}, nil
}
