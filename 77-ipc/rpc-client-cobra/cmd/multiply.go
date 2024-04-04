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

var (
	A, B int
	RPC  string
)

// This init function can be used any number of times.
// This is a special function
// As soon as the package is used , all init functions are called.[No need to call them specially]
func init() {
	multiplyCmd.Flags().IntVar(&A, "A", 0, "--A=0 or -A=0 or --A 0")
	multiplyCmd.Flags().IntVar(&B, "B", 0, "--B=0 or -B=0 or --B 0")
	multiplyCmd.Flags().StringVar(&RPC, "server", "localhost:8089", "--server=localhost:8089 or -server=localhost:8089 or --server localhost:8089")
}

// multiplyCmd represents the multiply command
var multiplyCmd = &cobra.Command{
	Use:   "multiply",
	Short: "A brief description of your command",
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
		err = client.Call("Calc.Multiply", modelargs, &result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Multiply:", result)
	},
}

func init() {
	rootCmd.AddCommand(multiplyCmd)

}
