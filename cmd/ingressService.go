package cmd

import (
	"github.com/masudur-rahman/appskube/appsclient"
	"github.com/spf13/cobra"
)

var ingressService = &cobra.Command{
	Use:	"ingress",
	Short:	"Ingress to a host",
	Run: func(cmd *cobra.Command, args []string) {
		appsclient.IngressService(host, name)
	},
}

func init() {
	ingressService.Flags().StringVarP(&host, "host", "H", "software.farm", "The hostname of the service")
	ingressService.Flags().StringVarP(&name, "name", "n", "appscode", "Name of the service")
	rootCmd.AddCommand(ingressService)
}