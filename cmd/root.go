/*
Copyright © 2022 Azat Nurzhanuly, Bekdaulet Shapigullin, Aslan Omirzak

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "EME",
	Short: "Allows user to identify encryption method used on a malware",
	Long: `This application - EME, is created to identify encryption methods
used to cipher a malware.

EME stands for executable method encryption
EME has several options to run.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
