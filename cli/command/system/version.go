package system

import (
	"github.com/kassisol/hbm/version"
	"github.com/spf13/cobra"
)

//NewVersionCommand new version command
func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the HBM version information",
		Long:  versionDescription,
		Run: func(cmd *cobra.Command, args []string) {
			info := version.New()
			info.ShowVersion()
		},
	}

	return cmd
}

var versionDescription = `
All software has versions. This is HBM's

`
