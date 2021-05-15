package cmd

import "github.com/spf13/cobra"

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "word format transformer",
	Long:  "Support different types of transformation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {}
