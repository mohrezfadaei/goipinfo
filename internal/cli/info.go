package cli

import (
	"fmt"
	"os"

	"github.com/mohrezfadaei/goipinfo/internal/config"
	"github.com/mohrezfadaei/goipinfo/internal/ip"
	"github.com/spf13/cobra"
)

var ipAddr string

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information for a specific IP address",
	Run: func(cmd *cobra.Command, args []string) {
		if ipAddr == "" {
			fmt.Println("IP address is required")
			os.Exit(1)
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			os.Exit(1)
		}

		fetcher := ip.NewIPAPIFetcher(cfg)
		info, err := fetcher.Fetch(ipAddr)
		if err != nil {
			fmt.Printf("Error fetching IP information: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("IP Information: %+v\n", info)
	},
}

func init() {
	InfoCmd.Flags().StringVar(&ipAddr, "ip", "", "IP address to fetch information for")
}
