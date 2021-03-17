package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "cat file",
	RunE: func(cmd *cobra.Command, args []string) error {
		objType, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}

		if objType == "" {
			objType = cmd.Flags().Arg(0) // TODO: implement required args
		}
		log.Println(objType)
		if err := catFile(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(catFileCmd)
}

func catFile() error {
	log.Println("cat-file!!!!")
	return nil
}
