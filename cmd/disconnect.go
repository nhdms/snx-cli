package cmd

import (
	"fmt"
	"os"

	"github.com/nhdms/snx-cli/internal/snx"

	"github.com/spf13/cobra"
)

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect from SNX VPN",
	Run: func(cmd *cobra.Command, args []string) {
		if err := snx.Disconnect(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Disconnected.")
	},
}

func init() {
	rootCmd.AddCommand(disconnectCmd)
}
