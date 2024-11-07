/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Command struct {
	Name	string `json:"name"`
	Base	string `json:"base"`
	// flags	string `json:"flags"`
}

type JsonData struct {
	Commands	map[string]Command `json:"commands"`
}

const FILE_NAME = ".pnotes"

func check(e error, message ...string) {
    if e != nil {
		if len(message) >= 1 && message[0] != "" {
			fmt.Println("##", message[0])
		}
        panic(e)
    }
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "project-notes-cli",
	Short: "CLI notes taking app for project commands",
	Long: `Project Notes is a CLI app to collect and reference notes.
This application stores run, build, or test commands to make
it easier to remember how to tun your program.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.project-notes-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


