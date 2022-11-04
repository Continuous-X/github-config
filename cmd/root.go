/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github-config/pkg/output"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	Config    *GHCConfig
	LogOutput *output.Output
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-config",
	Short: "github config synchronizer",
	Long:  `.......`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		LogOutput.AddLoggingLine(output.LogTypeError, "init", err.Error())
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.github-config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			LogOutput.AddLoggingLine(output.LogTypeError, "init", err.Error())
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".github-config.yaml".
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".github-config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		LogOutput.AddLoggingLine(output.LogTypeInfo, "init", fmt.Sprintf("Using config file:", viper.ConfigFileUsed()))
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		Config = &GHCConfig{}
		configReadErr := viper.Unmarshal(Config)
		if configReadErr != nil {
			LogOutput.AddLoggingLine(output.LogTypeError, "init", fmt.Sprintf("unable to decode into config struct, %v", configReadErr))
			fmt.Printf("unable to decode into config struct, %v", configReadErr)
		}
		LogOutput.AddLoggingLine(output.LogTypeInfo, "init", fmt.Sprintf("readed config: \n%v", Config))
		fmt.Printf("readed config: \n%v", Config)
	}

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		LogOutput.AddLoggingLine(output.LogTypeInfo, "init", fmt.Sprintf("'%s':'%s'", k, value))
		fmt.Printf("\n\"%s\":\"%s\"", k, value)
	}
}
