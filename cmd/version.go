/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get a version of the tool",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current version is alpha 1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
