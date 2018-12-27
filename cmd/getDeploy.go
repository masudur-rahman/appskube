package cmd

import (
	"github.com/masudur-rahman/appskube/appsclient"
	"github.com/spf13/cobra"
)

var getDeploy = &cobra.Command{
	Use:	"get",
	Run: func(cmd *cobra.Command, args []string) {
		appsclient.GetDeployment()
	},
}
func init(){
	rootCmd.AddCommand(getDeploy)
}
