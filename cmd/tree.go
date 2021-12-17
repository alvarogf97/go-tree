package cmd

import (
	"fmt"

	"github.com/alvarogf97/go-tree/pkg/shortcuts"
	"github.com/spf13/cobra"
)

// Tree command
func TreeCommand() *cobra.Command {
	var directory string
	var showHidden bool

	command := &cobra.Command{
		Use:   "go-tree",
		Short: "tree system filepath",
		Long:  "tree system filepath",
		Run: func(cmd *cobra.Command, args []string) {
			tree, err := shortcuts.Tree(directory, showHidden)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(tree)
		},
	}

	command.Flags().StringVarP(&directory, "directory", "d", "./", "directory (required)")
	command.Flags().BoolVarP(&showHidden, "show hidden", "s", false, "shows hidden files & folders")
	command.MarkFlagRequired("region")

	return command
}
