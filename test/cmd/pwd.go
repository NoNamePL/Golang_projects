/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// pwdCmd represents the pwd command
var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "print working directory",
	Long: `is a shell builtin command that prints the current location.
	The output shows an absolute directory path, starting with the root directory`,
	Run: func(cmd *cobra.Command, args []string) {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("current directory is", currentDir)

	},
}

func init() {
	rootCmd.AddCommand(pwdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pwdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pwdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
