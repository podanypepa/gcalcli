package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gcalcli",
		Short: "gcalcli is the best app for manage Google Calendar",
	}
)

// Execute executes the root command.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("%s", rootCmd.Execute())
	}
	return nil
}
