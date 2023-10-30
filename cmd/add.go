/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var taskId string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task.",
	Long: `Add a task to the lists in the project. You can specify the id of the task
	using the -t flag.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("name", "t", "The name of the task to add")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//  addCmd.Flags().BoolP("name", "t", false, "The name of the task to add")
	addCmd.PersistentFlags().StringVar(&taskId, "task", "t", "The name of the task to add")
	addCmd.Flags().BoolP("description", "d", false, "A short description of the task to add.")
}
