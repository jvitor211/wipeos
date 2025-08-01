package cmd

import (
	"fmt"
	"os"

	"github.com/joao-rrondon/wipeOs/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wipeOs",
	Short: ui.IconBanner() + " Professional secure file wiping tool",
	Long: ui.RenderWelcomeBanner() + `

WipeOs is a modern, secure file wiping utility that permanently removes files 
and sensitive data from your system using military-grade overwriting techniques.

Features:
• 🔒 Secure multi-pass overwriting (DoD 5220.22-M standard)
• 🎯 Target specific files, directories, or system artifacts
• 🚀 Fast and efficient with progress indicators
• 🛡️ Safe mode with confirmation prompts
• 📊 Detailed logging and reporting
• 🎨 Beautiful terminal interface`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no arguments provided, launch interactive mode by default
		if len(args) == 0 {
			fmt.Println(ui.StyleInfo("🚀 Launching WipeOs in interactive mode..."))
			fmt.Println(ui.StyleMuted("Use 'wipeOs --help' for CLI mode or 'exit' to quit interactive mode"))
			fmt.Println()
			
			// Launch interactive mode
			interactiveCmd.Run(cmd, args)
		} else {
			fmt.Println(ui.RenderWelcomeBanner())
			fmt.Println(ui.StyleSuccess("Welcome to WipeOs! Use --help to see available commands."))
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, ui.StyleError("Error: %v\n"), err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().BoolP("version", "", false, "Show version information")
} 