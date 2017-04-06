package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/chetan/earhart"
	"github.com/spf13/cobra"
)

var axfrCmd = &cobra.Command{
	Use:     "axfr",
	Aliases: []string{"list"},
	Short:   "AXFR query",
	Long:    `axfr query`,
	Run: func(cmd *cobra.Command, args []string) {
		axfr, err := earhart.GetAxfr(dnsUrl)
		if err != nil {
			fmt.Println("error: %s", err.Error())
			return
		}
		b, err := json.MarshalIndent(axfr, "", "  ")
		println(string(b))
	},
}

func init() {
	RootCmd.AddCommand(axfrCmd)
}
