/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type JsonData struct {
	Commands	[]Command `json:"commands"`
}

type Command struct {
	Name	string `json:"name"`
	Base	string `json:"base"`
	// flags	string `json:"flags"`

}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a command providing a name",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var commandName = "go run main.go get test-get"
        if len(args) >= 1 && args[0] != "" {
            commandName = args[0]
        }

		file, err := os.OpenFile(".pnotes", os.O_CREATE, 0664) // TODO better permissions and flags
		check(err)
		defer file.Close()
		
		byteValue, _ := io.ReadAll(file)


		var jsonData JsonData
		json.Unmarshal(byteValue, &jsonData)

		for i := 0; i < len(jsonData.Commands); i++ {
			fmt.Println("Base command: " + jsonData.Commands[i].Base)
			fmt.Println("Names: " + jsonData.Commands[i].Name)
			if jsonData.Commands[i].Name == commandName {
				fmt.Println("Found")

				cmdArgs := []string{"-c", jsonData.Commands[i].Base}
				cmd := exec.Command("bash", cmdArgs...) // TODO get this to happen inline
				output, err := cmd.CombinedOutput()
				check(err)
				fmt.Println("Output:", string(output))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
