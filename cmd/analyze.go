/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	mXOR "ABA/EME/app/methods/xor"
	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "GOes through every possible algorithm to detect the encryption type",
	Long: `Will iterate over every possible detection method
to find the one that is suitable. It will start with simple ciphers
like XOR or ROT13 and then continue with DES, AES - like ones.`,

	Run: func(cmd *cobra.Command, args []string) {
		xor, _ := cmd.Flags().GetBool("xor")
		filePath, _ := cmd.Flags().GetString("path")

		if xor {
			mXOR.Detect(filePath)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().String("path", "", "Gets filepath for an executable")
	err := analyzeCmd.MarkFlagRequired("path")
	if err != nil {
		return
	}
	// Here you will define your flags and configuration settings.
	analyzeCmd.Flags().BoolP("xor", "x", false, "uses only XOR detection")
}
