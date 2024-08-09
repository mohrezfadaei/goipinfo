package ip

import "testing"

func TestIPInfo(t *testing.T) {
	ipInfo := &IPInfo{
		IP:       "104.28.242.73",
		City:     "Düsseldorf",
		Region:   "North Rhine-Westphalia",
		Country:  "Germany",
		Location: "51.2217,6.7762",
		Org:      "AS13335 Cloudflare, Inc.",
		Timezone: "Europe/Berlin",
	}

	if ipInfo.IP != "104.28.242.73" {
		t.Errorf("Expected IP to be '104.28.242.73', got '%s'", ipInfo.IP)
	}
	if ipInfo.City != "Düsseldorf" {
		t.Errorf("Expected city to be 'Düsseldorf', got '%s'", ipInfo.City)
	}
	if ipInfo.Region != "North Rhine-Westphalia" {
		t.Errorf("Expected region to be 'North Rhine-Westphalia', got '%s'", ipInfo.Region)
	}
	if ipInfo.Country != "Germany" {
		t.Errorf("Expected country to be 'Germany', got '%s'", ipInfo.Country)
	}
	if ipInfo.Location != "51.2217,6.7762" {
		t.Errorf("Expected location to be '51.2217,6.7762', got '%s'", ipInfo.Location)
	}
	if ipInfo.Org != "AS13335 Cloudflare, Inc." {
		t.Errorf("Expected org to be 'AS13335 Cloudflare, Inc.', got '%s'", ipInfo.Org)
	}
	if ipInfo.Timezone != "Europe/Berlin" {
		t.Errorf("Expected timezone to be 'Europe/Berlin', got '%s'", ipInfo.Timezone)
	}
}
