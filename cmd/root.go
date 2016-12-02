package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh is a command line tool that builds and (re)starts your Go application everytime you save a Go or template file.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Refresh (%s)\n\n", Version)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "refresh.yml", "path to configuration file")
}
