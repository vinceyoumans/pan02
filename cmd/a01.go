/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vinceyoumans/cobraTest/libs/a01"
)

// a01Cmd represents the a01 command
var a01Cmd = &cobra.Command{
	Use:   "a01",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("a01 called")
		myStr, _ := cmd.Flags().GetString("myString")
		ret := a01.A01(myStr)
		fmt.Println(ret)
	},
}

func init() {
	rootCmd.AddCommand(a01Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// a01Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// a01Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	a01Cmd.Flags().String("myString", "MyString", "Please enter a string... default: hi There")

}
