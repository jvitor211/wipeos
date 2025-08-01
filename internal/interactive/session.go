package interactive

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joao-rrondon/wipeOs/internal/forensic"
	"github.com/joao-rrondon/wipeOs/internal/shredder"
	"github.com/joao-rrondon/wipeOs/ui"
)

type Model struct {
	textInput    textinput.Model
	history      []string
	output       []string
	shredder     *shredder.Shredder
	historyIndex int
	quitting     bool
}

type sessionEndMsg struct{}

func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Type a command... (help for commands, exit to quit)"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 80

	return Model{
		textInput: ti,
		history:   []string{},
		output:    []string{},
		shredder:  shredder.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit

		case tea.KeyEnter:
			input := strings.TrimSpace(m.textInput.Value())
			if input != "" {
				m.history = append(m.history, input)
				m.historyIndex = len(m.history)
				
				// Process command
				output := m.processCommand(input)
				m.output = append(m.output, ui.StyleInfo("> "+input))
				m.output = append(m.output, output...)
				
				// Keep output manageable
				if len(m.output) > 50 {
					m.output = m.output[len(m.output)-50:]
				}
				
				if input == "exit" || input == "quit" {
					m.quitting = true
					return m, tea.Quit
				}
			}
			m.textInput.SetValue("")

		case tea.KeyUp:
			if len(m.history) > 0 && m.historyIndex > 0 {
				m.historyIndex--
				m.textInput.SetValue(m.history[m.historyIndex])
				m.textInput.CursorEnd()
			}

		case tea.KeyDown:
			if m.historyIndex < len(m.history)-1 {
				m.historyIndex++
				m.textInput.SetValue(m.history[m.historyIndex])
				m.textInput.CursorEnd()
			} else if m.historyIndex == len(m.history)-1 {
				m.historyIndex = len(m.history)
				m.textInput.SetValue("")
			}
		}

	case sessionEndMsg:
		m.quitting = true
		return m, tea.Quit
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return ui.StyleSuccess("Thanks for using WipeOs! Stay secure! 🛡️\n")
	}

	var s strings.Builder

	// Header
	s.WriteString(ui.RenderWelcomeBanner())
	s.WriteString("\n")
	s.WriteString(ui.StyleHeader(ui.IconInteractive() + " Interactive Mode - Type commands below"))
	s.WriteString("\n\n")

	// Output history
	if len(m.output) > 0 {
		outputStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#4ECDC4")).
			Padding(1).
			Width(80).
			Height(15)

		// Show last lines that fit
		start := 0
		if len(m.output) > 12 {
			start = len(m.output) - 12
		}
		
		outputContent := strings.Join(m.output[start:], "\n")
		s.WriteString(outputStyle.Render(outputContent))
		s.WriteString("\n\n")
	}

	// Input prompt
	promptStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF6B6B")).
		Bold(true)

	s.WriteString(promptStyle.Render("WipeOs" + ui.Icon().Prompt + " "))
	s.WriteString(m.textInput.View())
	s.WriteString("\n\n")

	// Help
	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#747D8C")).
		Italic(true)

	s.WriteString(helpStyle.Render("💡 Commands: help, wipe, clean, forensic, version, status, clear, exit | ↑↓ for history | Ctrl+C to quit"))

	return s.String()
}

func (m *Model) processCommand(input string) []string {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return []string{ui.StyleError("No command provided")}
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "help", "h":
		return m.showHelp()
	
	case "clear", "cls":
		m.output = []string{}
		return []string{ui.StyleSuccess("Screen cleared!")}
	
	case "status":
		return m.showStatus()
	
	case "wipe":
		return m.handleWipe(args)
	
	case "clean":
		return m.handleClean(args)
	
	case "version", "v":
		return m.showVersion()
	
	case "forensic":
		return m.handleForensic(args)
	
	case "exit", "quit", "q":
		return []string{ui.StyleSuccess("Goodbye! 👋")}
	
	default:
		return []string{
			ui.StyleError(fmt.Sprintf("Unknown command: %s", command)),
			ui.StyleInfo("Type 'help' to see available commands"),
		}
	}
}

func (m *Model) showHelp() []string {
	return []string{
		ui.StyleHeader(ui.IconBanner() + " WipeOs Commands"),
		"",
		ui.StyleInfo(ui.IconWipe()+"  wipe <file>          - Securely delete files"),
		ui.StyleInfo(ui.IconClean()+"  clean <target>       - Quick predefined cleaning"),
		ui.StyleInfo(ui.IconForensic()+"  forensic <options>   - Advanced anti-forensic operations"),
		ui.StyleInfo(ui.Icon().Version+"   version              - Show version information"),
		ui.StyleInfo(ui.Icon().Status+"  status               - Show current session status"),
		ui.StyleInfo(ui.Icon().Clear+"  clear                - Clear screen output"),
		ui.StyleInfo(ui.Icon().Help+"  help                 - Show this help"),
		ui.StyleInfo(ui.Icon().Exit+"  exit                 - Quit WipeOs"),
		"",
		ui.StyleHeader(ui.Icon().Info + " Examples:"),
		ui.StyleMuted("  wipe test.txt --dry-run"),
		ui.StyleMuted("  clean browser"),
		ui.StyleMuted("  forensic --quick --dry-run"),
		ui.StyleMuted("  forensic --all --dry-run"),
		"",
		ui.StyleWarning(ui.IconWarning() + "  Always use --dry-run first to test!"),
	}
}

func (m *Model) showStatus() []string {
	return []string{
		ui.StyleHeader("📊 WipeOs Session Status"),
		"",
		ui.StyleInfo(fmt.Sprintf("Commands executed: %d", len(m.history))),
		ui.StyleInfo(fmt.Sprintf("Session active: %s", "Yes")),
		ui.StyleInfo(fmt.Sprintf("Safe mode: %s", "Enabled")),
		ui.StyleSuccess("🛡️ All systems operational"),
	}
}

func (m *Model) handleWipe(args []string) []string {
	if len(args) == 0 {
		return []string{
			ui.StyleError("Usage: wipe <file> [--dry-run] [--force] [--passes N]"),
			ui.StyleInfo("Example: wipe test.txt --dry-run"),
		}
	}

	// Parse basic flags
	files := []string{}
	dryRun := false
	force := false
	passes := 3

	for i, arg := range args {
		switch arg {
		case "--dry-run":
			dryRun = true
		case "--force":
			force = true
		case "--passes":
			if i+1 < len(args) {
				// Simple passes parsing (could be improved)
				passes = 3 // Default for now
			}
		default:
			if !strings.HasPrefix(arg, "--") {
				files = append(files, arg)
			}
		}
	}

	if len(files) == 0 {
		return []string{ui.StyleError("No files specified to wipe")}
	}

	options := shredder.WipeOptions{
		Recursive: false,
		Passes:    passes,
		Force:     force,
		DryRun:    dryRun,
	}

	if dryRun {
		return []string{
			ui.StyleWarning("🧪 DRY RUN MODE - No files will be deleted"),
			ui.StyleInfo(fmt.Sprintf("Would wipe %d file(s) with %d passes", len(files), passes)),
			ui.StyleMuted("Files: " + strings.Join(files, ", ")),
		}
	}

	results := m.shredder.WipeFiles(files, options)
	
	output := []string{}
	successCount := 0
	
	for _, result := range results {
		if result.Success {
			successCount++
			output = append(output, ui.StyleSuccess(fmt.Sprintf("✓ Wiped: %s", result.Path)))
		} else {
			output = append(output, ui.StyleError(fmt.Sprintf("✗ Failed: %s - %v", result.Path, result.Error)))
		}
	}
	
	output = append(output, ui.StyleHeader(fmt.Sprintf("📊 Summary: %d/%d files wiped successfully", successCount, len(results))))
	
	return output
}

func (m *Model) handleClean(args []string) []string {
	if len(args) == 0 {
		return []string{
			ui.StyleError("Usage: clean <target> [--dry-run]"),
			ui.StyleInfo("Targets: browser, temp, cache, all"),
			ui.StyleInfo("Example: clean browser --dry-run"),
		}
	}

	target := args[0]
	dryRun := false
	
	for _, arg := range args[1:] {
		if arg == "--dry-run" {
			dryRun = true
		}
	}

	options := shredder.WipeOptions{
		Recursive: true,
		Passes:    3,
		Force:     true,
		DryRun:    dryRun,
	}

	switch target {
	case "browser":
		if dryRun {
			return []string{
				ui.StyleWarning("🧪 DRY RUN: Would clean browser data"),
				ui.StyleMuted("• Cache files"),
				ui.StyleMuted("• History databases"),
				ui.StyleMuted("• Temporary files"),
			}
		}
		
		err := m.shredder.WipeBrowserData(options)
		if err != nil {
			return []string{ui.StyleError(fmt.Sprintf("Failed to clean browser data: %v", err))}
		}
		return []string{ui.StyleSuccess("🌐 Browser data cleaned successfully!")}

	case "temp":
		if dryRun {
			return []string{
				ui.StyleWarning("🧪 DRY RUN: Would clean temporary files"),
				ui.StyleMuted("• System temp directories"),
				ui.StyleMuted("• User cache folders"),
			}
		}
		
		err := m.shredder.WipeSystemTemp(options)
		if err != nil {
			return []string{ui.StyleError(fmt.Sprintf("Failed to clean temp files: %v", err))}
		}
		return []string{ui.StyleSuccess("📂 Temporary files cleaned successfully!")}

	case "all":
		if dryRun {
			return []string{
				ui.StyleWarning("🧪 DRY RUN: Would perform complete cleanup"),
				ui.StyleMuted("• Browser data"),
				ui.StyleMuted("• Temporary files"),
				ui.StyleMuted("• Cache directories"),
			}
		}
		
		output := []string{ui.StyleWarning("🧹 Performing complete cleanup...")}
		
		if err := m.shredder.WipeBrowserData(options); err != nil {
			output = append(output, ui.StyleError("Browser: Failed"))
		} else {
			output = append(output, ui.StyleSuccess("Browser: ✓"))
		}
		
		if err := m.shredder.WipeSystemTemp(options); err != nil {
			output = append(output, ui.StyleError("Temp files: Failed"))
		} else {
			output = append(output, ui.StyleSuccess("Temp files: ✓"))
		}
		
		output = append(output, ui.StyleSuccess("✨ Complete cleanup finished!"))
		return output

	default:
		return []string{
			ui.StyleError(fmt.Sprintf("Unknown clean target: %s", target)),
			ui.StyleInfo("Available targets: browser, temp, all"),
		}
	}
}

func (m *Model) showVersion() []string {
	return []string{
		ui.StyleHeader("WipeOs v1.0.0"),
		ui.StyleInfo("🔒 Military-grade secure file deletion"),
		ui.StyleInfo("🎨 Beautiful terminal interface"),
		ui.StyleMuted("Built with Go + Bubble Tea + Lip Gloss"),
	}
}

// handleForensic processes forensic anti-forensic commands
func (m *Model) handleForensic(args []string) []string {
	if len(args) == 0 {
		return []string{
			ui.StyleError("Usage: forensic <options>"),
			ui.StyleInfo("Options: --quick, --all, --dry-run, or specific flags"),
			ui.StyleInfo("Examples:"),
			ui.StyleMuted("  forensic --quick --dry-run"),
			ui.StyleMuted("  forensic --all --dry-run"),
			ui.StyleMuted("  forensic --logs --registry --dry-run"),
			"",
			ui.StyleWarning("⚠️  Use --dry-run first to preview operations!"),
		}
	}

	// Parse basic flags from args
	dryRun := false
	verbose := false
	all := false
	quick := false
	logs := false
	registry := false
	prefetch := false
	thumbnails := false
	eventlogs := false
	shadows := false
	memory := false
	swap := false
	freespace := false
	passes := 3

	for _, arg := range args {
		switch arg {
		case "--dry-run":
			dryRun = true
		case "--verbose", "-v":
			verbose = true
		case "--all":
			all = true
		case "--quick":
			quick = true
		case "--logs":
			logs = true
		case "--registry":
			registry = true
		case "--prefetch":
			prefetch = true
		case "--thumbnails":
			thumbnails = true
		case "--eventlogs":
			eventlogs = true
		case "--shadows":
			shadows = true
		case "--memory":
			memory = true
		case "--swap":
			swap = true
		case "--freespace":
			freespace = true
		case "--passes":
			// Simple implementation - could be improved
			passes = 3
		}
	}

	// Show danger warning for non-dry runs
	if !dryRun {
		return []string{
			ui.StyleError("🚨 DANGER: Anti-Forensic Operations"),
			ui.StyleWarning("This will permanently remove system traces!"),
			ui.StyleError("Use --dry-run first to preview operations"),
			ui.StyleInfo("Example: forensic --quick --dry-run"),
		}
	}

	// Set options based on flags
	options := forensic.ForensicCleanOptions{
		DryRun:  dryRun,
		Verbose: verbose,
		Passes:  passes,
	}

	// Determine what to clean
	if all {
		options.CleanLogs = true
		options.CleanRegistry = true
		options.CleanPrefetch = true
		options.CleanThumbnails = true
		options.CleanEventLogs = true
		options.CleanShadowCopies = true
		options.CleanMemory = true
		options.CleanSwap = true
		options.WipeFreespace = true
	} else if quick {
		options.CleanLogs = true
		options.CleanRegistry = true
		options.CleanThumbnails = true
		options.CleanEventLogs = true
	} else {
		options.CleanLogs = logs
		options.CleanRegistry = registry
		options.CleanPrefetch = prefetch
		options.CleanThumbnails = thumbnails
		options.CleanEventLogs = eventlogs
		options.CleanShadowCopies = shadows
		options.CleanMemory = memory
		options.CleanSwap = swap
		options.WipeFreespace = freespace
	}

	// Check if no operations selected
	if !options.CleanLogs && !options.CleanRegistry && !options.CleanPrefetch &&
		!options.CleanThumbnails && !options.CleanEventLogs &&
		!options.CleanShadowCopies && !options.CleanMemory && !options.CleanSwap &&
		!options.WipeFreespace {
		return []string{
			ui.StyleError("No operations selected"),
			ui.StyleInfo("Use --all, --quick, or specific flags"),
			ui.StyleInfo("Available flags: --logs, --registry, --prefetch, --thumbnails, etc."),
		}
	}

	// Perform anti-forensic operations
	antiForensic := forensic.New(dryRun, verbose)
	results := antiForensic.PerformForensicCleanup(options)

	// Format results for display
	output := []string{
		ui.StyleHeader("🔍 Anti-Forensic Operation Results:"),
		"",
	}

	successCount := 0
	for _, result := range results {
		if result.Success {
			successCount++
			output = append(output, ui.StyleSuccess(fmt.Sprintf("✓ %s: %s", result.Operation, result.Details)))
		} else {
			output = append(output, ui.StyleError(fmt.Sprintf("✗ %s: %v", result.Operation, result.Error)))
		}
	}

	output = append(output, "")
	output = append(output, ui.StyleHeader(fmt.Sprintf("📊 Summary: %d/%d operations completed", successCount, len(results))))

	if dryRun {
		output = append(output, ui.StyleInfo("🧪 This was a dry run - no actual operations performed"))
	}

	return output
} 