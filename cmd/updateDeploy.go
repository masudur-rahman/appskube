package cmd

import (
	"github.com/masudur-rahman/appskube/appsclient"
	"github.com/spf13/cobra"
)

var updateDeploy = &cobra.Command{
	Use: 	"scale",
	Short: 	"Scale to a specific number",
	Run: 	func(cmd *cobra.Command, args []string) {
		appsclient.UpdateDeploymentKutil(name, replicas)
	},
}

func init() {
	updateDeploy.Flags().StringVarP(&name, "name", "n", "appscode", "Name of the Deployment")
	updateDeploy.Flags().Int32VarP(&replicas, "replicas", "r", 5, "Number of replicas")
	rootCmd.AddCommand(updateDeploy)
}