package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/timwehrle/asars/pkg/api"
	"github.com/timwehrle/asars/pkg/asars"
	"github.com/timwehrle/asars/pkg/auth"
	"github.com/timwehrle/asars/utils"
)

var (
	showComments bool
	tasks        []api.Task
)

func init() {
	TaskCmd.Flags().BoolVarP(&showComments, "comments", "c", false, "Show comments for the task")
}

var TaskCmd = &cobra.Command{
	Use:     "task [index]",
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

		if len(args) == 0 && len(tasks) > 0 {
			task := tasks[0]
			fmt.Println("Task Details (First Task):")
			fmt.Printf("ID: %s\nName: %s\nDue: %s\n", task.GID, task.Name, task.DueOn)
			return
		}

		if len(args) > 0 {
			index, err := strconv.Atoi(args[0])
			if err != nil || index < 1 || index > len(tasks) {
				fmt.Printf("Invalid index. Please choose a number between 1 and %d\n", len(tasks))
				os.Exit(1)
			}

			task := tasks[index-1]
			detailedTask, err := client.GetTask(workspace, task.GID)
			if err != nil {
				fmt.Println("Error getting task details:", err)
				return
			}

			fmt.Printf("%s [%s], %s\n", utils.BoldUnderline.Sprint(detailedTask.Name), utils.FormatDate(detailedTask.DueOn), displayProjects(detailedTask.Projects))
			fmt.Println(displayTags(detailedTask.Tags))

			fmt.Print(displayNotes(detailedTask.Notes))

			if showComments {
				comments, err := client.GetStories(workspace, detailedTask.GID)
				if err != nil {
					fmt.Println("Error getting comments:", err)
					return
				}

				if len(comments) > 0 {
					fmt.Println("Comments:")
					for _, comment := range comments {
						fmt.Printf("%s: %s\n", comment.CreatedBy.Name, comment.Text)
					}
				} else {
					fmt.Println("Comments: None")
				}
			}
		}
	},
}

func displayProjects(projects []api.Project) string {
	if len(projects) > 0 {
		var projectNames []string
		for _, project := range projects {
			projectNames = append(projectNames, project.Name)
		}
		return "Projects: " + fmt.Sprintf("%s", projectNames)
	}

	return "Projects: None"
}

func displayTags(tags []api.Tag) string {
	if len(tags) > 0 {
		var tagNames []string
		for _, tag := range tags {
			tagNames = append(tagNames, tag.Name)
		}
		return "Tags: " + fmt.Sprintf("%s", tagNames)
	}
	return "Tags: None"
}

func displayNotes(notes string) string {
	if notes != "" {
		return fmt.Sprintf("\n%s\n", notes)
	}

	return ""
}
