/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// mvCmd represents the mv command
var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "moveing from one folder to another",
	Long:  `moveing from one folder to another`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Flags().GetString(cmd.ValidArgs[0])
	},
}

func changeDir(cmd *cobra.Command,args []string) {
	os.Chdir(args[0])
}


func init() {
	rootCmd.AddCommand(mvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mvCmd.PersistentFlags().String("foo", "", "A help for foo")
	
	mvCmd.PersistentFlags().String("change Directory","",)
	

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")


	mvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
