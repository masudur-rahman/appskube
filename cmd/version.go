package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of AppsKube",
	Long:  "The version of the AppsKube app is",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AppsKube - v0.0 -- HEAD")
	},
}
