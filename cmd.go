package version

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

// NewCmd creates a new version command.
func NewCmd(info Info) *cobra.Command {
	return &cobra.Command{
		Use:          "version",
		Short:        "Displays version information.",
		Long:         "Displays version information.",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return json.NewEncoder(os.Stdout).Encode(info)
		},
	}
}
