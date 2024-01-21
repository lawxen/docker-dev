/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dev",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.docker-dev.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getFirstContainer() (containerName, containerPort string) {
	// get the config info from the execute result of "docker compose config"
	composeCmd := exec.Command("docker", "compose", "config")
	config, err := composeCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Sth wrong:", err)
		return
	}
	// 解析 Docker Compose 配置
	var dockerComposeConfig map[string]interface{}
	err = yaml.Unmarshal(config, &dockerComposeConfig)
	if err != nil {
		fmt.Println("Sth wrong:", err)
		return
	}
	services, ok := dockerComposeConfig["services"].(map[string]interface{})
	if !ok {
		fmt.Println("Sth wrong: services not found")
		return
	}
	for _, service := range services {
		serviceMap := service.(map[string]interface{})
		containerName = serviceMap["container_name"].(string)
		containerPort = serviceMap["ports"].([]interface{})[0].(map[string]interface{})["published"].(string)
		break
	}
	return

}
