/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/vinceyoumans/cobraTest/libs/a03"

	"github.com/spf13/cobra"
)

// a03Cmd represents the a03 command
var a03Cmd = &cobra.Command{
	Use:   "a03",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("a03 called")

		myStr, _ := cmd.Flags().GetString("myString")
		ret := a03.A03(myStr)
		fmt.Println(ret)

	},
}

func init() {
	rootCmd.AddCommand(a03Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// a03Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// a03Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	a03Cmd.Flags().String("myString", "MyString", "Please enter a string... default: hi There")

}
