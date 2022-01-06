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
			Use:   "add",
			Short: "Add a new task to a list",
			Long:  "Adds a new task to your list. The list stores the task itself and a deadline.",
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) == 0 {
					fmt.Fprintln(os.Stderr, "Specify the tasks you need to add")
					os.Exit(1)
				}
				if len(args)%2 != 0 {
					fmt.Fprintln(os.Stderr, "The number of arguments cannot be odd. The arguments come in pairs taskName and deadline")
					os.Exit(1)
				}
				for i := 0; i < len(args)-1; i += 2 {
					if err := tasklistmanager.AddTask(args[i], args[i+1]); err != nil {
						fmt.Fprintln(os.Stderr, "An error occured: "+err.Error()+", skipping the task "+args[i])
					}
				}
				fmt.Println("The tasks were successfully added to your manager")
			},
		},
	)
}
