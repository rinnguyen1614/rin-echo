package migrate

import (
	"github.com/spf13/cobra"
)

var (
	configYaml string
	StartCmd   = &cobra.Command{
		Use:     "migrate",
		Short:   "Migrate database",
		Example: "rin-echo migrate -c config/config.yaml",
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYaml, "config", "c", "config/config.yaml", "Start server with provided configuration file")
}
