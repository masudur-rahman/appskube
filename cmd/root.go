package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var name string
var replicas int32
var host string

var rootCmd = &cobra.Command{
	Use:	"appskube",
	Short:	"Short description of appskube",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Welcome from AppsKube...!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
