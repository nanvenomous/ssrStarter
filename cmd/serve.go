/*
Copyright © 2025 nanvenomous mrgarelli@gmail.com
*/
package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/nanvenomous/ssrStarter/handle"
	"github.com/spf13/cobra"
)

func runService() error {
	midMux, err := handle.Setup(http.NewServeMux(), buildFS)
	if err != nil {
		return err
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	log.Printf("App running on port %s", port)
	// log.Printf("App running in docker container on port %s", port)
	// log.Printf("caddyserver/caddy serving air-verse/air proxy to host machine at http://localhost:4001")

	return http.ListenAndServe(":"+port, midMux)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run ssr starter service",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runService()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
