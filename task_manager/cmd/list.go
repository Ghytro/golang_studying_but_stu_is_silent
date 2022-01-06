package cmd

import (
	"fmt"
	tasklistmanager "gotutorial/task_manager/tasklist_manager"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "list",
			Short: "Show a list of your tasks",
			Long:  "Show an enumerated list of your tasks with info about deadlines",
			Run: func(cmd *cobra.Command, args []string) {
				if l, err := tasklistmanager.GetTaskList(); err != nil {
					fmt.Fprintf(os.Stderr, "An error occured: %s", err.Error())
				} else {
					if len(l) == 0 {
						fmt.Println("Your task list is empty")
						return
					}
					for i, t := range l {
						fmt.Printf("%d) %s, deadline: %s\n", i+1, t.Name, t.Deadline)
					}
				}
			},
		},
	)
}
