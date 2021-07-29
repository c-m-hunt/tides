package cmd

import (
	"fmt"
	"os"

	"github.com/c-m-hunt/tides/pkg/tides"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tides",
	Short: "tides shows tides tables for locations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dt := tides.GetTides()
		dt.Display()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
