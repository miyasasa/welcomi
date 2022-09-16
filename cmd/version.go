package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var version = "1.0.0"

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of welcomi",
		Long:  `All software has versions. This is Welcomi's.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			printVersion()
			return nil
		},
	}
}

func printVersion() {
	jww.FEEDBACK.Println(fmt.Sprintf("welcomi version %s", version))
}
