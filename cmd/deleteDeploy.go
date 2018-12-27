package cmd

import (
	"github.com/masudur-rahman/appskube/appsclient"
	"github.com/spf13/cobra"
)

var nameDelete string

var deleteDeploy = &cobra.Command{
	Use:	"delete",
	Short:	"Deleting everything related to this Deployment",
	Run: func(cmd *cobra.Command, args []string) {
		appsclient.DeleteDeployment(nameDelete)
	},
}

func init()  {
	deleteDeploy.Flags().StringVarP(&nameDelete, "name", "n", "appscode", "Name of the deployment")
	rootCmd.AddCommand(deleteDeploy)
}
