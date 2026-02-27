package cmd

import (
	"fmt"
	"os"

	"github.com/nhdms/snx-cli/internal/snx"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check VPN connection status",
	Run: func(cmd *cobra.Command, args []string) {
		connected, err := snx.Status()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if connected {
			fmt.Println("VPN: Connected")
		} else {
			fmt.Println("VPN: Disconnected")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
