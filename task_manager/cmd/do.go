package cmd

import (
	"fmt"
	tasklistmanager "gotutorial/task_manager/tasklist_manager"

	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "do",
			Short: "Perform a task from your manager",
			Long:  "The commands removes the task from the manager",
			Run: func(cmd *cobra.Command, args []string) {
				for _, a := range args {
					if intArg, err := strconv.Atoi(a); err != nil { // the argument is a number, so we need to remove task by index
						tasklistmanager.DoTaskByIndex(intArg)
					} else { // else we need to do task by name
						tasklistmanager.DoTaskByName(a)
					}
				}
				fmt.Println("All the listed tasks were performed. If some tasks were not in the list, they were ignored.")
			},
		},
	)
}
