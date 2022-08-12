package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:     "welcomi",
	Version: version,
	Short:   "info - a simple CLI to install required sdks, tools and packages for new developers",
	Long:    `welcomi is a CLI tool that automatically installs predetermined sdks, tools and packages for new developers.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing welcomi '%s'", err)
		os.Exit(1)
	}
}
