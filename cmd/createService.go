package cmd

import (
	"github.com/masudur-rahman/appskube/appsclient"
	"github.com/spf13/cobra"
)

var createService = &cobra.Command{
	Use:	"expose",
	Short:	"Creates a Service",
	Run:	func(cmd *cobra.Command, args []string) {
		appsclient.CreateService(name)
	},
}

func init() {
	createService.Flags().StringVarP(&name, "name", "n", "appscode", "Name of the service")
	rootCmd.AddCommand(createService)
}
