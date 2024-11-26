package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/timwehrle/asars/pkg/auth"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the Asana Personal Access Token",
	Run: func(cmd *cobra.Command, args []string) {
		if !auth.HasToken() {
			fmt.Println("There is no token to delete.")
			return
		}

		err := auth.DeleteToken()
		if err != nil {
			fmt.Println("Error deleting token:", err)
			os.Exit(1)
		}
	},
}
