package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/timwehrle/asars/pkg/api"
	"github.com/timwehrle/asars/pkg/asars"
	"github.com/timwehrle/asars/pkg/auth"
)

func init() {
	ConfigCmd.AddCommand(DeleteCmd)
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure your Asana Personal Access Token and default workspace",
	Run: func(cmd *cobra.Command, args []string) {
		if auth.HasToken() {
			fmt.Println("Token is already set.")
			return
		}

		token := promptForToken()
		if token == "" {
			fmt.Println("No token entered. Aborting.")
			return
		}

		err := auth.SetToken(token)
		if err != nil {
			fmt.Println("Error setting token:", err)
			os.Exit(1)
		}

		client := api.NewClient(token)
		workspaces, err := client.GetWorkspaces()
		if err != nil {
			fmt.Println("Error getting workspaces:", err)
			return
		}

		fmt.Println("Available workspaces:")
		for index, workspace := range workspaces {
			fmt.Printf("%d. %s (GID: %s)\n", index+1, workspace.Name, workspace.GID)
		}

		selected := promptForWorkspaceSelection(len(workspaces))
		if selected == -1 {
			fmt.Println("Invalid selection. Aborting.")
			return
		}

		selectedWorkspace := workspaces[selected]
		if err := asars.SaveDefaultWorkspace(selectedWorkspace.GID); err != nil {
			fmt.Println("Error saving default workspace:", err)
		} else {
			fmt.Printf("Default workspace set to: %s (GID: %s)\n", selectedWorkspace.Name, selectedWorkspace.GID)
		}
	},
}

func promptForToken() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Asana Personal Access Token: ")
	token, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	return strings.TrimSpace(token)
}

func promptForWorkspaceSelection(count int) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select the number of the default workspace: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return -1
	}

	selection, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || selection < 1 || selection > count {
		fmt.Println("Invalid selection. Please enter a number between 1 and", count)
		return -1
	}

	return selection - 1
}
