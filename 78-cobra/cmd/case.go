/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	upperCmd.Flags().StringVar(&str, "string", "HELLO WORLD", `-str="hello world"`)
	lowerCmd.Flags().StringVar(&str, "string", "hello world", `-str="Hello World"`)

}

// upperCmd represents the upper command
var upperCmd = &cobra.Command{
	Use:   "upper",
	Short: "upper writes in UPPER case",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.ToUpper(str))
	},
}
var lowerCmd = &cobra.Command{
	Use:   "lower",
	Short: "lower writes in lower case",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.ToLower(str))
	},
}

func init() {
	echoCmd.AddCommand(upperCmd)
	echoCmd.AddCommand(lowerCmd)
}
