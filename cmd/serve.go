package cmd

import (
	"github.com/spf13/cobra"
	"github.com/werberson/dejavu/web"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		return web.Initialize()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
