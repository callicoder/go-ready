package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/callicoder/go-ready/internal/app"
	"github.com/spf13/cobra"
)

func newServerCommand() *cobra.Command {
	var serverCmd = &cobra.Command{
		Use:          "server",
		Short:        "Run the go-ready server",
		RunE:         serverCmdF,
		SilenceUsage: true,
	}
	return serverCmd
}

func serverCmdF(command *cobra.Command, args []string) error {
	config, err := command.Flags().GetString("config")

	if err != nil {
		return err
	}

	interruptChan := make(chan os.Signal, 1)
	return runServer(config, interruptChan)
}

func runServer(configFileLocation string, interruptChan chan os.Signal) error {
	app, err := app.New(configFileLocation)
	if err != nil {
		return err
	}
	app.Start()

	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan

	app.Shutdown()
	return nil
}
