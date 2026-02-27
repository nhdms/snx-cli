package cmd

import (
	"fmt"
	"os"

	"github.com/nhdms/snx-cli/internal/config"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a default config file",
	Run: func(cmd *cobra.Command, args []string) {
		path := cfgFile
		if path == "" {
			path = config.DefaultPath()
		}

		if _, err := os.Stat(path); err == nil {
			fmt.Fprintf(os.Stderr, "Config already exists at %s\n", path)
			os.Exit(1)
		}

		if err := config.CreateDefault(path); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating config: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Config created at %s (permissions: 0600)\n", path)
		fmt.Println("Edit it with your VPN credentials and TOTP secret.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
