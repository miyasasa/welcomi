package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type commandBuilder struct {
	commands []*cobra.Command
}

func newCommandBuilder() *commandBuilder {
	return &commandBuilder{}
}

func (b *commandBuilder) addCommands(commands ...*cobra.Command) *commandBuilder {
	b.commands = append(b.commands, commands...)
	return b
}

func (b *commandBuilder) addAll() *commandBuilder {
	b.addCommands(newVersionCmd())
	return b
}

func (b *commandBuilder) build() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "welcomi",
		Short: "info - a simple CLI to install required sdks, tools and packages for new developers",
		Long:  `welcomi is a CLI tool that automatically installs predetermined sdks, tools and packages for new developers.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rootCmd.AddCommand(b.commands...)
	return rootCmd
}

func Execute() {
	welcomiCmd := newCommandBuilder().addAll().build()

	if err := welcomiCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing welcomi '%s'", err)
		os.Exit(1)
	}
}
