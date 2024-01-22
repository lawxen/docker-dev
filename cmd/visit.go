/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

// visitCmd represents the visit command
var visitCmd = &cobra.Command{
	Use:   "visit",
	Short: "Use default browser to visit the web app",
	Long:  `Use default browser to visit the web app`,
	Run: func(cmd *cobra.Command, args []string) {
		containerName, containerPort := getFirstContainer()
		host := "http://" + containerName + ".docker:" + containerPort

		err := open.Run(host)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(visitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// visitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// visitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
