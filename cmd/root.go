package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	baseUrl string
	apiKey  string
)

var rootCmd = &cobra.Command{
	Use:   "rolgo-cli",
	Short: "Manage device rental in Racks of Labs",
	Long:  `Manage device rental in Racks of Labs`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
