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
	Use:     "tasks",
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

		lines := []string{"Your Tasks:"}
		for index, task := range tasks {
			line := fmt.Sprintf("%d. [%s] %s", index+1, utils.FormatDate(task.DueOn), task.Name)
			lines = append(lines, line)
		}
		fmt.Println(strings.Join(lines, "\n"))
	},
}
