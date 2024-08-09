package utils

import (
	"os"

	"github.com/mohrezfadaei/goipinfo/internal/ip"
	"github.com/olekukonko/tablewriter"
)

func DisplayTable(info *ip.IPInfo) {
	data := [][]string{
		{"IP", info.IP},
		{"City", info.City},
		{"Region", info.Region},
		{"Country", info.Country},
		{"Location", info.Location},
		{"Org", info.Org},
		{"Timezone", info.Timezone},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
