package cmd

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completionCmd)
}

var completionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate completion script",
	Long:                  "To load completions",
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			if err := cmd.Root().GenBashCompletion(os.Stdout); err != nil {
				log.Error().Err(err).Send()
			}
		case "zsh":
			if err := cmd.Root().GenZshCompletion(os.Stdout); err != nil {
				log.Error().Err(err).Send()
			}
		case "fish":
			if err := cmd.
				Root().GenFishCompletion(os.Stdout, true); err != nil {
				log.Error().Err(err).Send()
			}
		case "powershell":
			if err := cmd.
				Root().GenPowerShellCompletionWithDesc(os.Stdout); err != nil {
				log.Error().Err(err).Send()
			}
		}
	},
}
