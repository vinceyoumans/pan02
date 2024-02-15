/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	cobraTest "github.com/vinceyoumans/cobraTest/libs/t01"

	"github.com/spf13/cobra"
)

// t01Cmd represents the t01 command
var t01Cmd = &cobra.Command{
	Use:   "t01",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("t01 called")
		fmt.Println("and the responce is", cobraTest.T01("hi there"))
	},
}

func init() {
	rootCmd.AddCommand(t01Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// t01Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// t01Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
