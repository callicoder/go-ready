package cmd

import (
	"github.com/spf13/cobra"
)

func NewCLI() *cobra.Command {
	cli := &cobra.Command{
		Use:   "go-ready",
		Short: "GoDockerCompose Example App",
		Long:  "GoDockerCompose is an example app to demonstrate docker compose, docker machine, and docker swarm functionalities",
	}

	cli.AddCommand(newServerCommand())

	cli.PersistentFlags().StringP("config", "c", "config/application.yml", "Configuration file to use (default is `config/application.yml`).")
	return cli
}
