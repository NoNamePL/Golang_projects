/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "testiki",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

func Execute() error {
	return rootCmd.Execute()
}

/*
	func Execute() {
		err := rootCmd.Execute()
		if err != nil {
			os.Exit(1)
		}
	}
*/
func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "cmd/config.go")
	rootCmd.PersistentFlags().StringP("author", "a", "Roman", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "LICENSE")
	rootCmd.PersistentFlags().Bool("viper", true, "user Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("userViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Roman <rdymshakov@inbox.ru>")
	viper.SetDefault("license", "apache")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(pwdCmd)
	rootCmd.AddCommand(mvCmd)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.testiki.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		//user config file from tha flag
		//currentDir, err := os.Getwd()

		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}

	} else {
		//find home directory
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		//search config at home directory with name ".cobra" (without extension)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()
	/*
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	*/
}
