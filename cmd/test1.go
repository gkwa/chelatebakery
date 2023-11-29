/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// test1Cmd represents the test1 command
var test1Cmd = &cobra.Command{
	Use:   "test1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test1 called")
		runTest1()
	},
}

func init() {
	rootCmd.AddCommand(test1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// FileWriter is a custom type that implements the io.Writer interface
type FileWriter struct {
	FileName string
}

// Write implements the io.Writer interface for FileWriter
func (fw FileWriter) Write(p []byte) (n int, err error) {
	// Open the file with the specified filename
	file, err := os.OpenFile(fw.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Write the data to the file
	return file.Write(p)
}

func runTest1() {
	// Create an instance of FileWriter with a specific filename
	fileWriter := FileWriter{FileName: "output.txt"}

	// Use Fprintf to format and write data to FileWriter
	fmt.Fprintf(fileWriter, "This data will be written to the file.\n")
}
