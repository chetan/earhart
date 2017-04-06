package cmd

import (
	"github.com/chetan/earhart"
	"github.com/spf13/cobra"
)

var waitForService string

// wait_forCmd represents the wait_for command
var wait_forCmd = &cobra.Command{
	Use:     "wait_for",
	Aliases: []string{"waitfor", "wait", "w"},
	Short:   "A brief description of your command",
	Long:    `wait for cmd`,
	Run: func(cmd *cobra.Command, args []string) {
		earhart.GetService(dnsUrl, waitForService)
	},
}

func init() {
	RootCmd.AddCommand(wait_forCmd)
	wait_forCmd.Flags().StringVarP(&waitForService, "svc", "s", "", "Service to wait for")
}
