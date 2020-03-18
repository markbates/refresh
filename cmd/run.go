package cmd

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/markbates/refresh/refresh"
	"github.com/spf13/cobra"
)

// ErrConfigNotExist is returned when a configuration file cannot be found.
var ErrConfigNotExist = errors.New("no config file was found")

func init() {
	RootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r", "start", "build", "watch"},
	Short:   "(default) watches your files and rebuilds/restarts your app accordingly.",
	Run: func(cmd *cobra.Command, args []string) {
		Run(cfgFile)
	},
}

func Run(cfgFile string) error {
	ctx := context.Background()
	return RunWithContext(cfgFile, ctx)
}

func RunWithContext(cfgFile string, ctx context.Context) error {
	c := &refresh.Configuration{}

	if err := loadConfig(c, cfgFile); err != nil {
		if err != ErrConfigNotExist {
			return err
		}

		log.Println("No configuration loaded, proceeding with defaults")
	}

	if len(c.Path) > 0 {
		log.Printf("Configuration loaded from %s\n", c.Path)
	}

	if debug {
		c.Debug = true
	}

	r := refresh.NewWithContext(c, ctx)
	return r.Start()
}

func loadConfig(c *refresh.Configuration, path string) error {
	if len(path) > 0 {
		return c.Load(path)
	}

	for _, f := range [4]string{
		".refresh.yml",
		".refresh.yaml",
		"refresh.yml",
		"refresh.yaml",
	} {
		err := c.Load(f)
		if err != nil && os.IsNotExist(err) {
			continue
		}
		return err
	}

	return ErrConfigNotExist
}
