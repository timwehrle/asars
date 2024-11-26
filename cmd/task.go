package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/timwehrle/asars/pkg/api"
	"github.com/timwehrle/asars/pkg/asars"
	"github.com/timwehrle/asars/pkg/auth"
)

var tasks []api.Task

var TaskCmd = &cobra.Command{
	Use:     "task",
	Aliases: []string{"t"},
	Short:   "Get a single Asana task with details by index",
	Long:    "Get a single Asana task with details by index. If no index is provided, the first task will be shown.",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token, err := auth.GetToken()
		if err != nil {
			fmt.Println("Error getting token:", err)
			return
		}

		workspace, err := asars.LoadDefaultWorkspace()
		if err != nil {
			fmt.Println("Error getting default workspace:", err)
			return
		}

		client := api.NewClient(token)

		if len(tasks) == 0 {
			var errGetTasks error
			tasks, errGetTasks = client.GetTasks(workspace)
			if errGetTasks != nil {
				fmt.Println("Error getting tasks:", errGetTasks)
				return
			}
		}

		if len(args) == 0 {
			if len(tasks) > 0 {
				task := tasks[0]
				fmt.Println("Task Details (First Task):")
				fmt.Printf("ID: %s\nName: %s\nDue: %s\n", task.GID, task.Name, task.DueOn)
			}
		} else {
			index, err := strconv.Atoi(args[0])
			if err != nil || index < 1 || index > len(tasks) {
				fmt.Printf("Invalid index. Please choose a number between 1 and %d\n", len(tasks))
				os.Exit(1)
			}

			task := tasks[index-1]
			fmt.Println("Task Details:")
			fmt.Printf("ID: %s\nName: %s\nDue: %s\n", task.GID, task.Name, task.DueOn)
		}
	},
}
