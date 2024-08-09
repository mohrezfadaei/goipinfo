package cli

import (
	"fmt"
	"os"

	"github.com/mohrezfadaei/goipinfo/internal/config"
	"github.com/mohrezfadaei/goipinfo/internal/ip"
	"github.com/spf13/cobra"
)

var provider string

var MyipCmd = &cobra.Command{
	Use:   "myip",
	Short: "Get your public IP address",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			os.Exit(1)
		}

		var fetcher ip.IPInfoFetcher

		switch provider {
		case "ipinfo":
			fetcher = ip.NewIPInfoIOFetcher(cfg)
		case "ip-api":
			fetcher = ip.NewIPAPIFetcher(cfg)
		case "my-ip":
			fetcher = ip.NewMyIPFetcher(cfg)
		default:
			fmt.Println("Invalid provider. Supported providers are: ipinfo, ip-api, my-ip")
			os.Exit(1)
		}

		info, err := fetcher.Fetch("")
		if err != nil {
			fmt.Println("Error fetching IP information: ", err)
			os.Exit(1)
		}

		fmt.Printf("IP information: %+v\n", info)
	},
}

func init() {
	MyipCmd.Flags().StringVarP(&provider, "provider", "p", "ip-api", "API provider (ipinfo, ip-api, my-ip)")
}
