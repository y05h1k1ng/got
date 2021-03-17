package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "got",
	Short: "got = git + go!!!!",
}

func Execute() {
	initCmd.Flags().String("directory", ".", "initialize directory")
	catFileCmd.Flags().String("type", "", "can be one of: blob, tree, commit, tag")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
