/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// inCmd represents the in command
var inCmd = &cobra.Command{
	Use:   "in",
	Short: "Docker exec into the first container of docker compose or the specified container",
	Long: `Docker exec into the first container of docker compose or the specified container. For example:

docker-dev-drupal in
docker-dev-drupal in <container_name>`,
	Run: func(cmd *cobra.Command, args []string) {
		var containerName string

		// If args is not empty, then get the container name from the first argument
		if len(args) != 0 {
			containerName = args[0]
		} else {
			containerInfo := getFirstContainer()
			containerName = containerInfo["container_name"].(string)
		}

		finalCmd := exec.Command("docker", "exec", "-it", containerName, "sh")
		finalCmd.Stdin = os.Stdin
		finalCmd.Stdout = os.Stdout
		finalCmd.Stderr = os.Stderr

		finalCmd.Run()
	},
}

func init() {
	rootCmd.AddCommand(inCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
