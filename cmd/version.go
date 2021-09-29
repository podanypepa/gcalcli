package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/podanypepa/gcalcli/pkg/version"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gcalcli",
	Long:  `All software has versions. This is gcalcli`,
	Run:   showVersion,
}

func showVersion(cmd *cobra.Command, args []string) {
	fmt.Println("gcalcli version:")
	fmt.Printf("%s\n", version.GetBuildInfo())
}
