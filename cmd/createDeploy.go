package cmd

import (
	"github.com/masudur-rahman/appskube/appsclient"
	"github.com/spf13/cobra"
)

var nameCreate string
var replicaCreate int32



var createDeploy = &cobra.Command{
	Use:	"create",
	Short:	"Creates deployment for AppsCodeServer",
	Run: 	func(cmd *cobra.Command, args []string) {
		appsclient.CreateDeploymentKutil(nameCreate, replicaCreate)
	},
}

func init()  {
	createDeploy.Flags().StringVarP(&nameCreate, "name", "n", "appscode", "Name of the deployment")
	createDeploy.Flags().Int32VarP(&replicaCreate, "replicas", "r", 1, "Number of replicas")
	rootCmd.AddCommand(createDeploy)
}
