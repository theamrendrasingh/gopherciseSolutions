package cmd

import (
	"github.com/spf13/cobra"
)

// var cfgFile string

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A command line task manager",
}
