/*
Copyright 2023 The K8sGPT Authors.

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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	kubeconfig string
	namespace  string
	output     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k8sgpt",
	Short: "K8sGPT is a tool for scanning your Kubernetes clusters and diagnosing issues",
	Long: `K8sGPT gives Kubernetes superpowers to everyone. It scans your Kubernetes
clusters, diagnoses issues, and provides AI-powered explanations and
recommendations in plain English.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags available to all subcommands
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.k8sgpt.yaml)")
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "path to the kubeconfig file (default is $HOME/.kube/config)")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "namespace to analyze (default is all namespaces)")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "text", "output format (text, json)")

	// Bind flags to viper
	_ = viper.BindPFlag("kubeconfig", rootCmd.PersistentFlags().Lookup("kubeconfig"))
	_ = viper.BindPFlag("namespace", rootCmd.PersistentFlags().Lookup("namespace"))
	_ = viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Search config in home directory with name ".k8sgpt" (without extension)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".k8sgpt")
	}

	// Read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
