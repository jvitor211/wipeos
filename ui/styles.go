package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Color palette
	colorPrimary   = lipgloss.Color("#FF6B6B")
	colorSecondary = lipgloss.Color("#4ECDC4")
	colorSuccess   = lipgloss.Color("#45B7D1")
	colorWarning   = lipgloss.Color("#FFA07A")
	colorError     = lipgloss.Color("#FF4757")
	colorInfo      = lipgloss.Color("#5F27CD")
	colorMuted     = lipgloss.Color("#747D8C")

	// Base styles
	baseStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA"))

	// Component styles
	headerStyle = baseStyle.Copy().
			Foreground(colorPrimary).
			Bold(true).
			MarginBottom(1)

	successStyle = baseStyle.Copy().
			Foreground(colorSuccess).
			Bold(true)

	errorStyle = baseStyle.Copy().
			Foreground(colorError).
			Bold(true)

	warningStyle = baseStyle.Copy().
			Foreground(colorWarning).
			Bold(true)

	infoStyle = baseStyle.Copy().
		      Foreground(colorInfo)

	mutedStyle = baseStyle.Copy().
		     Foreground(colorMuted)

	// Banner style
	bannerStyle = lipgloss.NewStyle().
			Foreground(colorPrimary).
			Bold(true).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorSecondary).
			Padding(1, 2).
			MarginBottom(2)
)

// RenderWelcomeBanner creates an ASCII art banner
func RenderWelcomeBanner() string {
	banner := `
 ██╗    ██╗██╗██████╗ ███████╗ ██████╗ ███████╗
 ██║    ██║██║██╔══██╗██╔════╝██╔═══██╗██╔════╝
 ██║ █╗ ██║██║██████╔╝█████╗  ██║   ██║███████╗
 ██║███╗██║██║██╔═══╝ ██╔══╝  ██║   ██║╚════██║
 ╚███╔███╔╝██║██║     ███████╗╚██████╔╝███████║
  ╚══╝╚══╝ ╚═╝╚═╝     ╚══════╝ ╚═════╝ ╚══════╝
`
	
	subtitle := fmt.Sprintf("%s Professional Secure File Wiping Tool v1.0.0", IconBanner())
	
	return bannerStyle.Render(banner + "\n" + subtitle)
}

// Style functions
func StyleHeader(text string) string {
	return headerStyle.Render(text)
}

func StyleSuccess(text string) string {
	return successStyle.Render(text)
}

func StyleError(text string) string {
	return errorStyle.Render(text)
}

func StyleWarning(text string) string {
	return warningStyle.Render(text)
}

func StyleInfo(text string) string {
	return infoStyle.Render(text)
}

func StyleMuted(text string) string {
	return mutedStyle.Render(text)
}

// ConfirmDangerous asks for user confirmation for dangerous operations
func ConfirmDangerous(operation string) bool {
	fmt.Printf(StyleWarning("⚠️  You are about to %s. This action is IRREVERSIBLE!\n"), operation)
	fmt.Print(StyleInfo("Type 'yes' to continue: "))
	
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	
	input = strings.TrimSpace(strings.ToLower(input))
	return input == "yes"
}

// ProgressBar creates a simple progress indicator
func ProgressBar(current, total int, label string) string {
	if total == 0 {
		return ""
	}
	
	percentage := float64(current) / float64(total) * 100
	barWidth := 30
	filled := int(float64(barWidth) * percentage / 100)
	
	bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)
	
	style := lipgloss.NewStyle().
		Foreground(colorSecondary).
		Bold(true)
	
	return style.Render(fmt.Sprintf("[%s] %.1f%% %s (%d/%d)", bar, percentage, label, current, total))
} 