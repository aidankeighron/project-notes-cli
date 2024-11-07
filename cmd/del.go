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

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete saved command",
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
			fmt.Println("del [name]")
			return;
		}

		fmt.Println("Removing", commandName)

		// TODO Make sure fix exists

		// Read the JSON file
		data, err := os.ReadFile(FILE_NAME)
		check(err, "Error while reading file")

		// Unmarshal the JSON data into a map
		var jsonData JsonData
		err = json.Unmarshal(data, &jsonData)
		check(err, "Error un-marshaling json file")

		// Remove the new element
		_, exists := jsonData.Commands[commandName]
		if !exists {
			fmt.Println("Command does not exist")
			return;
		}
		delete(jsonData.Commands, commandName)

		// Marshal the updated data back to JSON
		updatedData, err := json.MarshalIndent(jsonData, "", "  ")
		check(err, "Error marshaling file into json")

		// Write the updated JSON back to the file
		err = os.WriteFile(FILE_NAME, updatedData, 0644)
		check(err, "Error writing to file");
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
