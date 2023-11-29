/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// test2Cmd represents the test2 command
var test2Cmd = &cobra.Command{
	Use:   "test2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test2 called")
		runTest2()
	},
}

func init() {
	rootCmd.AddCommand(test2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// MyWriter is a custom type that implements the io.Writer interface
type MyWriter struct{}

// Write implements the io.Writer interface for MyWriter
func (mw MyWriter) Write(p []byte) (n int, err error) {
	// In this example, we simply print the data to the console
	fmt.Println("Writing:", string(p))
	return len(p), nil
}

func runTest2() {
	// Create an instance of MyWriter
	myWriter := MyWriter{}

	// Use Fprintf to format and write data to MyWriter
	fmt.Fprintf(myWriter, "Hello, %s!", "World")
}
