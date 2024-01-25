/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/skratchdot/open-golang/open"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		containerInfo := getFirstContainer()
		containerName := containerInfo["container_name"].(string)
		virtual_host := containerInfo["environment"].(map[string]interface{})["VIRTUAL_HOST"].(string)
		host := "http://" + virtual_host
		finalCmd := exec.Command("docker", "exec", containerName, "drush", "uli")
		output, err := finalCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error execute drush uli: %v\n", err)
			return
		}
		outputString := string(output)

		loginUrl := strings.Replace(outputString, "http://default", host, 1)

		// Drop the \n in the end
		loginUrl = strings.TrimRight(loginUrl, "\n")

		// Open the login url in default browser
		err = open.Run(loginUrl)

		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
