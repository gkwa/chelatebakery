package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

// test3Cmd represents the test1 command
var test3Cmd = &cobra.Command{
	Use:   "test3",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test1 called")

		// Check if a style argument is provided
		if len(args) == 0 {
			fmt.Println("Please provide a template style as an argument (e.g., 'style1')")
			return
		}

		style := args[0]
		switch style {
		case "style1":
			runTest3(style1Template, "output.txt")
		case "style2":
			runTest3(style2Template, "output.txt")
		default:
			fmt.Println("Unknown template style:", style)
		}
	},
}

func init() {
	rootCmd.AddCommand(test3Cmd)
	// Add other flags and configurations as needed
}

// Define multiple template styles
const (
	style1Template = `Style 1:
Message: {{.Message}}
Count: {{.Count}}
`

	style2Template = `Style 2:
Message: {{.Message}} (Count: {{.Count}})
`
)

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

// runTest1 runs the test with the specified template and filename
func runTest3(templateString, fileName string) {
	// Example data to be formatted
	data := struct {
		Message string
		Count   int
	}{
		Message: "Hello, World!",
		Count:   42,
	}

	// Create a new template and parse the template string
	tmpl, err := template.New("output").Parse(templateString)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Create an instance of FileWriter with the specified filename
	fileWriter := FileWriter{FileName: fileName}

	// Execute the template and write the output to FileWriter
	err = tmpl.Execute(fileWriter, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}
