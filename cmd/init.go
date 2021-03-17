package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"yoshiking/got/schema"
	"yoshiking/got/util"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new, empty repository.",
	RunE: func(cmd *cobra.Command, argv []string) error {
		dir, err := cmd.Flags().GetString("directory")
		if err != nil {
			return err
		}

		if arg := cmd.Flags().Arg(0); dir == "." && arg != "" {
			dir = arg
		}
		log.Println(dir)

		if err := initialize(dir); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initialize(dir string) error {
	if !util.IsExist(dir) {
		if err := os.MkdirAll(dir, 0775); err != nil {
			return err
		}
	}
	current, err := os.Getwd()
	if err != nil {
		return err
	}
	repo := schema.GitRepository{}
	if err = repo.New(filepath.Join(current, dir)); err != nil {
		return err
	}
	return nil
}
