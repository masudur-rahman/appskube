package cmd

import (
	"github.com/masudur-rahman/appskube/appsclient"
	"github.com/spf13/cobra"
)

var createDeploy = &cobra.Command{
	Use:	"create",
	Short:	"Creates deployment for AppsCodeServer",
	Run: 	func(cmd *cobra.Command, args []string) {
		appsclient.CreateDeployment(name, replicas)
	},
}

func init()  {
	createDeploy.Flags().StringVarP(&name, "name", "n", "appscode", "Name of the deployment")
	createDeploy.Flags().Int32VarP(&replicas, "replicas", "r", 1, "Number of replicas")
	rootCmd.AddCommand(createDeploy)
}
