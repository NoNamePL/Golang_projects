/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "displays the contents of a file in the terminal ",
	Long: `The cat command (concatenate) displays the contents of a file in the terminal (standard output or stdout).
	To use the command, provide a file name from the current directory:`,
	Run: func(cmd *cobra.Command, args []string) {
		mess, err := os.ReadFile(args[0])

		if err != nil {
			file, _ := os.Create(args[0])
			defer file.Close()

			return
		}

		fmt.Println(string(mess))

	},
}

func init() {
	rootCmd.AddCommand(catCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
