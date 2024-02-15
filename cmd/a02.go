/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/vinceyoumans/cobraTest/libs/a02"

	"github.com/spf13/cobra"
)

// a02Cmd represents the a02 command
var a02Cmd = &cobra.Command{
	Use:   "a02",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("a02 called")

		myStr, _ := cmd.Flags().GetString("myString")
		ret := a02.A02(myStr)
		fmt.Println(ret)

	},
}

func init() {
	rootCmd.AddCommand(a02Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// a02Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// a02Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	a02Cmd.Flags().String("myString", "MyString", "Please enter a string... default: hi There")

}
