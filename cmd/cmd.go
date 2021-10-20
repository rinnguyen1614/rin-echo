package cmd

import (
	"errors"
	"os"
	"rin-echo/cmd/databases/migrate"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "rin-echo",
		Short:        "rin-echo",
		SilenceUsage: true,
		Long:         `rin-echo`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				tip()
				return errors.New("requires at least one arg, " +
					"you can view the available parameters through `--help`",
				)
			}
			return nil
		},
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
		Run: func(cmd *cobra.Command, args []string) {
			tip()
		},
	}
)

func tip() {
	// add tip for usage
}

func init() {
	rootCmd.AddCommand(migrate.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
