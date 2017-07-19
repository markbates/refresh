package cmd

import (
	"context"
	"log"
	"os"

	"github.com/markbates/refresh/refresh"
	"github.com/spf13/cobra"
)

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

func Run(cfgFile string) {
	ctx := context.Background()
	RunWithContext(cfgFile, ctx)
}

func RunWithContext(cfgFile string, ctx context.Context) {
	c := &refresh.Configuration{}
	err := c.Load(cfgFile)
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}

	if debug {
		c.Debug = true
	}
	r := refresh.NewWithContext(c, ctx)
	err = r.Start()
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
}
