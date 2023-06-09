package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"mgw/mgw-resi/cmd/apiserver/app"
	"mgw/mgw-resi/config"
	"os"
)

var (
	rootCMD = &cobra.Command{
		Short: "mgw-backend",
	}

	configCMD = &cobra.Command{
		Use:   "config",
		Short: "Show settings",
		Run: func(*cobra.Command, []string) {
			fmt.Printf("cobra command: ")
			config.Show()
		},
	}

	serverCMD = &cobra.Command{
		Use:   "server",
		Short: "Run application server",
		Run: func(*cobra.Command, []string) {
			fmt.Println("Running application")
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
