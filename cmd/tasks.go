package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/timwehrle/asars/pkg/api"
	"github.com/timwehrle/asars/pkg/asars"
	"github.com/timwehrle/asars/pkg/auth"
	"github.com/timwehrle/asars/utils"
)

var TasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List your tasks assigned in Asana",
	Long: `Retrieve a numbered list of tasks assigned to you in your Asana account.
Each task includes its due date and name. This command works in the context
of your default workspace.`,
	Aliases: []string{"ts"},
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
		tasks, err := client.GetTasks(workspace)
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}

		// Define the width of the number column for consistent alignment
		numberWidth := 3

		// Initialize the lines slice to store task output
		lines := []string{"Your Tasks:"}
		for index, task := range tasks {
			// Format the task number with a trailing dot for readability
			numberWithDot := fmt.Sprintf("%d.", index+1)

			// Format the task line with padded number, due date, and name
			line := fmt.Sprintf("%-*s [%s] %s", numberWidth, numberWithDot, utils.FormatDate(task.DueOn), task.Name)

			// Add the formatted line to the output slice
			lines = append(lines, line)
		}

		// Print all tasks as a single joined string
		fmt.Println(strings.Join(lines, "\n"))
	},
}
