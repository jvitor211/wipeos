package cmd

import (
	"github.com/joao-rrondon/wipeOs/ui"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: ui.IconStart() + " Start WipeOs (same as interactive mode)",
	Long: `Start WipeOs in interactive mode. This is the main way to use WipeOs
where you get a persistent terminal session with command history,
autocompletion, and beautiful interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Just run the interactive command
		interactiveCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
} 