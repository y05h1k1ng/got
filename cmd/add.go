package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

func add() error {
	fmt.Println("add!!!!")
	return nil
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add any files.",
	RunE: func(cmd *cobra.Command, argv []string) error {
		if err := add(); err != nil {
			return err
		}
		return nil
	},
}
