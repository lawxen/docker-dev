/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// crCmd represents the cr command
var crCmd = &cobra.Command{
	Use:   "cr",
	Short: "Clear the drupal cache",
	Long: `Clear the drupal cache of the first docker compose container or the specified container. For example:

	docker-dev cr
	docker-dev cr <container_name>`,
	Run: func(cmd *cobra.Command, args []string) {
		var containerName string

		// If args is not empty, then get the container name from the first argument
		if len(args) != 0 {
			containerName = args[0]
		} else {
			containerName, _ = getFirstContainer()
		}
		fmt.Println("Clear cache for container:", containerName)
		finalCmd := exec.Command("docker", "exec", containerName, "drush", "cr")
		finalCmd.Stdin = os.Stdin
		finalCmd.Stdout = os.Stdout
		finalCmd.Stderr = os.Stderr

		finalCmd.Run()
	},
}

func init() {
	rootCmd.AddCommand(crCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
