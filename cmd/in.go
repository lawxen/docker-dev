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

// // 定义结构体，用于解析 Docker Compose 配置
// type DockerComposeConfig struct {
// 	Name     string `yaml:"name"`
// 	Services map[string]struct {
// 		ContainerName string            `yaml:"container_name"`
// 		Environment   map[string]string `yaml:"environment"`
// 		Image         string            `yaml:"image"`
// 		// 添加其他你需要的字段
// 	} `yaml:"services"`
// 	Networks map[string]struct {
// 		Name     string `yaml:"name"`
// 		External bool   `yaml:"external"`
// 	} `yaml:"networks"`
// }

// inCmd represents the in command
var inCmd = &cobra.Command{
	Use:   "in",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var containerName string
		var containerPort string
		//If args is not empty, then get the container name from the first argument
		// if len(args) != 0 {
		// 	containerName := args[0];
		// }

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
			fmt.Println(containerName)
			fmt.Println(containerPort)
			break
		}


		finalCmd := exec.Command("docker", "exec", "-it", containerName, "bash")
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
