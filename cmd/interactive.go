package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joao-rrondon/wipeOs/internal/interactive"
	"github.com/joao-rrondon/wipeOs/ui"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: ui.IconInteractive() + " Launch WipeOs in interactive mode",
	Long: ui.StyleHeader("Interactive Terminal Mode") + `

Launch WipeOs in an interactive terminal session where you can execute
commands in a persistent environment, similar to a shell or REPL.

Features:
• 🎯 Persistent session with command history
• 🔄 Command autocompletion and navigation
• 📊 Real-time output and status updates
• 🎨 Beautiful terminal interface
• ⚡ Quick access to all WipeOs commands

Commands available in interactive mode:
• wipe <file> [flags]    - Secure file deletion
• clean <target>         - Predefined cleaning tasks
• status                 - Session information
• help                   - Show available commands
• clear                  - Clear screen output
• exit                   - Quit interactive mode

Navigation:
• ↑/↓ arrows            - Browse command history
• Ctrl+C or ESC         - Exit interactive mode
• Enter                 - Execute command`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the interactive model
		model := interactive.NewModel()
		
		// Create the Bubble Tea program
		p := tea.NewProgram(
			model,
			tea.WithAltScreen(),       // Use alternate screen buffer
			tea.WithMouseCellMotion(), // Enable mouse support
		)
		
		// Run the program
		finalModel, err := p.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, ui.StyleError("Error starting interactive mode: %v\n"), err)
			os.Exit(1)
		}
		
		// Check if we need to exit with an error
		if m, ok := finalModel.(interactive.Model); ok {
			_ = m // Could check for any final state if needed
		}
	},
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
} 