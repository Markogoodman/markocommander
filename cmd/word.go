package cmd

import (
	"log"
	"strings"

	"github.com/Markogoodman/markocommander/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var desc = strings.Join([]string{
	"1: To Upper",
	"2: To Lower",
	"3: Underscore to upper camel case",
	"4: Underscore to lower camel case",
	"5: Camel case to underscore",
}, "\n")

func NewWordCmd() *cobra.Command {
	var (
		str  string
		mode int8
	)
	wordCmd := &cobra.Command{
		Use:   "word",
		Short: "word format transformer",
		Long:  desc,
		Run: func(cmd *cobra.Command, args []string) {
			var content string
			switch mode {
			case ModeUpper:
				content = word.ToUpper(str)
			case ModeLower:
				content = word.ToLower(str)
			case ModeUnderscoreToUpperCamelCase:
				content = word.UnderscoreToUpperCamelCase(str)
			case ModeUnderscoreToLowerCamelCase:
				content = word.UnderscoreToLowerCamelCase(str)
			case ModeCamelCaseToUnderscore:
				content = word.CamelCaseToUnderscore(str)
			default:
				log.Fatal("Unsupported mode")
			}
			log.Printf("Output: %s", content)
		},
	}
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "Input string")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "Input mode")
	return wordCmd
}
