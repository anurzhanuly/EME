/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/anurzhanuly/EME/app/methods"
	mROT "github.com/anurzhanuly/EME/app/methods/rot13"
	mXOR "github.com/anurzhanuly/EME/app/methods/xor"
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
		//xor, _ := cmd.Flags().GetBool("xor")
		filePath, _ := cmd.Flags().GetString("path")

		detectors := initDetectors(filePath)

		for _, detector := range detectors {
			if result, err := detector.Detect(); err == nil && result {
				detector.Present()

				return
			}
		}

		fmt.Println("Encryption method is not detected")
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

func initDetectors(filepath string) []methods.Detector {
	detectors := []methods.Detector{
		&mXOR.Detector{Filepath: filepath},
		&mROT.Detector{Filepath: filepath},
	}

	return detectors
}
