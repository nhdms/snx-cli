package cmd

import (
	"fmt"
	"os"

	"github.com/nhdms/snx-cli/internal/config"
	"github.com/nhdms/snx-cli/internal/snx"
	"github.com/nhdms/snx-cli/internal/totp"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to SNX VPN",
	Run: func(cmd *cobra.Command, args []string) {
		path := cfgFile
		if path == "" {
			path = config.DefaultPath()
		}

		cfg, err := config.Load(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Generate TOTP
		code, err := totp.Generate(cfg.TOTPSecret)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating TOTP: %v\n", err)
			os.Exit(1)
		}

		password := cfg.FixedPassword + code
		fmt.Printf("Connecting to %s as %s...\n", cfg.Server, cfg.Username)

		if err := snx.Connect(cfg.Server, cfg.Username, password); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Connected successfully.")
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
