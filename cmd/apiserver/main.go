package main

import (
	"os"

	"github.com/spf13/cobra"

	"go-skeleton/cmd/apiserver/app"
	"go-skeleton/config"
)

var (
	rootCMD = &cobra.Command{
		Short: "skeleton-go",
	}

	configCMD = &cobra.Command{
		Use:   "config",
		Short: "Show settings",
		Run: func(*cobra.Command, []string) {
			config.Show()
		},
	}

	serverCMD = &cobra.Command{
		Use:   "server",
		Short: "Run application server",
		Run: func(*cobra.Command, []string) {
			app.Run()
		},
	}
)

func main() {
	cobra.OnInitialize(config.Init)

	// Regist
	rootCMD.AddCommand(configCMD)
	rootCMD.AddCommand(serverCMD)
	if err := rootCMD.Execute(); err != nil {
		os.Exit(1)
	}
}
