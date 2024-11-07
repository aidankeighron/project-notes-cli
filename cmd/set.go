/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Inserts command into .pnotes file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var commandName string
		var commandBase string
        if len(args) == 2 && args[0] != "" {
            commandName = args[0]
            commandBase = args[1]
        } else {
			fmt.Println("set [name] [command]")
			return;
		}

		fmt.Println("Setting", commandName, "->", commandBase)

		// TODO Make sure fix exists

		// Read the JSON file
		data, err := os.ReadFile(FILE_NAME)
		check(err, "Error while reading file")

		// Unmarshal the JSON data into a map
		var jsonData JsonData
		err = json.Unmarshal(data, &jsonData)
		check(err, "Error un-marshaling json file")

		// Add the new element
		var command Command
		command.Name = commandName
		command.Base = commandBase

		_, exists := jsonData.Commands[commandName]
		if exists {
			fmt.Println("Overwriting existing command")
		}
		jsonData.Commands[commandName] = command

		// Marshal the updated data back to JSON
		updatedData, err := json.MarshalIndent(jsonData, "", "  ")
		check(err, "Error marshaling file into json")

		// Write the updated JSON back to the file
		err = os.WriteFile(FILE_NAME, updatedData, 0644)
		check(err, "Error writing to file");
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
