package cmd

import (
	"fmt"
	"os"
	"text/template"

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
	// Add other flags and configurations as needed
}

// FileWriter2 is a custom type that implements the io.Writer interface
type FileWriter2 struct {
	FileName string
}

// Write implements the io.Writer interface for FileWriter2
func (fw FileWriter2) Write(p []byte) (n int, err error) {
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
	// Create an instance of FileWriter2 with a specific filename
	fileWriter2 := FileWriter2{FileName: "output.txt"}

	// Example data to be formatted
	data := struct {
		Message string
		Count   int
	}{
		Message: "Hello, World!",
		Count:   42,
	}

	// Define a template
	outputTemplate := `Message: {{.Message}}
Count: {{.Count}}
`

	// Create a new template and parse the template string
	tmpl, err := template.New("output").Parse(outputTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Execute the template and write the output to FileWriter2
	err = tmpl.Execute(fileWriter2, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}
