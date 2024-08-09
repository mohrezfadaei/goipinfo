package ip

type IPInfo struct {
	IP       string
	City     string
	Region   string
	Country  string
	Location string
	Org      string
	Timezone string
}

type IPInfoFetcher interface {
	Fetch(ip string) (*IPInfo, error)
}
