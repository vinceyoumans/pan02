package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vinceyoumans/pan02/libs/mon01"
	"github.com/vinceyoumans/pan02/structs"
)

// mon01Cmd represents the a01 command
var mon01Cmd = &cobra.Command{
	Use:   "mon01",
	Short: "Pan challange to monitor Subdir",
	Long: `a cobra CLI to...

1. Monitor a subdir
2. Maintain a running JSON file of Files and File Sizes
----------------------------------------------------------
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mon01 called")

		mon := structs.Mon{}
		mon.DirToMonitor, _ = cmd.Flags().GetString("dirToMonitor")
		mon.DirTorecordJSON, _ = cmd.Flags().GetString("dirTorecordJSON")
		mon.ConcMan, _ = cmd.Flags().GetInt("concMan")

		fmt.Println("======================================")
		fmt.Println("=====   Begining FIle Monitor ========")
		fmt.Println("=====   Dir to Monitor =", mon.DirToMonitor)
		fmt.Println("=====   dir To recordJSON =", mon.DirTorecordJSON)
		fmt.Println("=====   concMan =", mon.ConcMan)
		fmt.Println("======================================")

		ret := mon01.Mon01(mon)
		fmt.Println(ret)
	},
}

func init() {
	rootCmd.AddCommand(mon01Cmd)
	mon01Cmd.Flags().String("dirToMonitor", "./tmp", "What is the Directory to be Monitored")
	mon01Cmd.Flags().String("dirTorecordJSON", "./json", "What is Directory to save JSON results")

	//  not certain how ti implement
	mon01Cmd.Flags().Int("concMan", 10, "Concurrency Management: Process files concurrently. The degree of concurrency (i.e., the number of files processed in parallel) must be configurable.")

}
