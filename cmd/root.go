/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/graydovee/clipboardshare/pkg/logger"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "share clipboard",
	Long:  `share clipboard from different machine`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&logger.DefaultOptions.EncoderJson, "log-encode-json", false, "log output with json format")
	rootCmd.Flags().StringVar(&logger.DefaultOptions.LogFile, "log-file", "", "log output file, if empty, only output to std")
	rootCmd.Flags().BoolVar(&logger.DefaultOptions.Development, "log-development", false, "development mod")
}
