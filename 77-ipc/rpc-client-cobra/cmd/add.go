/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"demo/models"
	"flag"
	"fmt"
	"net/rpc"

	"github.com/spf13/cobra"
)

// This init function can be used any number of times.
// This is a special function
// As soon as the package is used , all init functions are called.[No need to call them specially]
func init() {
	addCmd.Flags().IntVar(&A, "A", 0, "--A=0 or -A=0 or --A 0")
	addCmd.Flags().IntVar(&B, "B", 0, "--B=0 or -B=0 or --B 0")
	addCmd.Flags().StringVar(&RPC, "server", "localhost:8089", "--server=localhost:8089 or -server=localhost:8089 or --server localhost:8089")
}

// multiplyCmd represents the multiply command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add command add two numbers",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		client, err := rpc.Dial("tcp", RPC)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer client.Close()
		modelargs := new(models.Args)
		modelargs.A = A
		modelargs.B = B
		var result int
		err = client.Call("Calc.Add", modelargs, &result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Add:", result)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
