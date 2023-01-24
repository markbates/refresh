package cmd

import (
	"fmt"
	"os"

	"github.com/markbates/refresh/refresh"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "generates a default configuration file for you.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Do not report errors as wrong usage
		cmd.SilenceUsage = true

		c := refresh.Configuration{
			AppRoot:            ".",
			IgnoredFolders:     []string{"vendor", "log", "logs", "tmp", "node_modules", "bin", "templates"},
			IncludedExtensions: []string{".go"},
			BinaryName:         "refresh-build",
			BuildDelay:         200,
			BuildTargetPath:    "",
			BuildPath:          os.TempDir(),
			CommandFlags:       []string{},
			CommandEnv:         []string{},
			EnableColors:       true,
			Livereload: refresh.Livereload{
				Enable:          false,
				Port:            35729,
				IncludedFolders: []string{},
				Tasks:           []string{},
			},
		}

		if cfgFile == "" {
			cfgFile = "refresh.yml"
		}

		_, err := os.Stat(cfgFile)
		if !os.IsNotExist(err) {
			return fmt.Errorf("config file %q already exists, skipping init", cfgFile)
		}

		return c.Dump(cfgFile)
	},
}
