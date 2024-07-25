package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:     "version",
	Short:   "Get version info",
	Example: "server version",
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

func run() error {
	fmt.Println("0.0.1")
	return nil
}
