/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "prints a list of the current directory's contents",
	Long:  `prints a list of the current directory's contents`,
	Run: func(cmd *cobra.Command, args []string) {
		//found current directory
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		//found all files in currensy directory
		allDir, err := os.ReadDir(currentDir)
		if err != nil {
			fmt.Println(err)
			return
		}
		//print files to consol
		for idx := 0; idx < len(allDir); idx++ {
			fmt.Println(allDir[idx].Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
