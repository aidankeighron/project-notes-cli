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
        var commandName string
		if len(args) == 1 && args[0] != "" {
            commandName = args[0]
        } else {
			fmt.Println("get [name]")
			return;
		}

		// TODO
		file, err := os.OpenFile(FILE_NAME, os.O_CREATE, 0664)
		check(err, "Error when opening/creating file")
		defer file.Close()
		
		byteValue, _ := io.ReadAll(file)
		var jsonData JsonData
		json.Unmarshal(byteValue, &jsonData)

		existingCommand, exists := jsonData.Commands[commandName]
		if exists {
			fmt.Println("Base command: " + existingCommand.Base)
			fmt.Println("Names: " + existingCommand.Name)
			if existingCommand.Name == commandName {
				fmt.Println("Found")
	
				cmdArgs := []string{"-c", existingCommand.Base}
				cmd := exec.Command("bash", cmdArgs...) // TODO get this to happen inline
				output, err := cmd.CombinedOutput()
				check(err, "Error when running command")
				fmt.Println("Output:", string(output))
			}
		} else {
			fmt.Println("Command not found")
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
