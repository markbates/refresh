package cmd

import (
	"github.com/markbates/refresh/refresh"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "generates a default configuration file for you.",
	Run: func(cmd *cobra.Command, args []string) {
		c := refresh.Configuration{
			AppRoot:            ".",
			IgnoredFolders:     []string{"vendor", "log", "tmp"},
			IncludedExtensions: []string{".go"},
			BuildPath:          "/tmp",
			BuildDelay:         200,
			BinaryName:         "refresh-build",
			CommandFlags:       []string{},
			EnableColors:       true,
		}
		c.Dump(cfgFile)
	},
}
