// Copyright © 2018 James Rivett-Carnac <james.rivettcarnac@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	gamePath   string
	workingDir string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "list",
	Short: "list available versions to download",
	Long: `Launcher manages and launches version of CDDA, worlds and mods

It has both a CLI, and a wishfull GUI (its all in your head)`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.launcher.yaml)")
	rootCmd.Flags().StringVar(&gamePath, "gameExecutable", "", "Executable for game")
	rootCmd.Flags().StringVar(&workingDir, "workingDir", "", "working dir")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	tf := &logrus.TextFormatter{}
	logrus.SetFormatter(tf)
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".launcher" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".launcher")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
