package cmd

import (
	"os"

	"github.com/mohrezfadaei/goipinfo/internal/cli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goipinfo",
	Short: "A CLI tool to get IP information",
	Long:  "goipinfo is a CLI tool to fetch IP information from multiple providers and display it.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cli.MyipCmd)
	rootCmd.AddCommand(cli.InfoCmd)
}
