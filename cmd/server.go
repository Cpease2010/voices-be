package cmd

import (
	"log"
	"net/http"
	"voices/db"
	"voices/routes"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the Voices backend server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.Connect(); err != nil {
			log.Fatalf("DB connection failed: %v", err)
		}

		routes.RegisterRoutes()

		log.Println("Server running on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
