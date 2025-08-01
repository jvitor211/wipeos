package cmd

import (
	"fmt"

	"github.com/joao-rrondon/wipeOs/internal/shredder"
	"github.com/joao-rrondon/wipeOs/ui"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: ui.IconClean() + " Quick clean with predefined targets",
	Long: ui.StyleHeader("Quick Clean Operations") + `

This command provides quick access to common cleaning operations
without needing to specify individual files or complex patterns.

Available clean operations:
  all        - Clean everything (browser data + system temp + common junk)
  browser    - Clean all browser data (same as 'wipe --browser-data')
  temp       - Clean system temporary files
  logs       - Clean application and system logs
  cache      - Clean user cache directories
  downloads  - Clean Downloads folder (with confirmation)

Examples:
  wipeOs clean all                # Clean everything
  wipeOs clean browser temp       # Clean browser data and temp files
  wipeOs clean logs --dry-run     # Preview log cleaning`,
	ValidArgs: []string{"all", "browser", "temp", "logs", "cache", "downloads"},
	Args:      cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		force, _ := cmd.Flags().GetBool("force")
		passes, _ := cmd.Flags().GetInt("passes")

		options := shredder.WipeOptions{
			Recursive: true,
			Passes:    passes,
			Force:     force,
			DryRun:    dryRun,
		}

		s := shredder.New()

		for _, target := range args {
			switch target {
			case "all":
				fmt.Println(ui.StyleWarning("🧹 Performing comprehensive cleanup..."))
				cleanBrowser(s, options)
				cleanTemp(s, options)
				cleanLogs(s, options)
				cleanCache(s, options)

			case "browser":
				cleanBrowser(s, options)

			case "temp":
				cleanTemp(s, options)

			case "logs":
				cleanLogs(s, options)

			case "cache":
				cleanCache(s, options)

			case "downloads":
				cleanDownloads(s, options)

			default:
				fmt.Printf(ui.StyleError("Unknown clean target: %s\n"), target)
				fmt.Println(ui.StyleInfo("Available targets: all, browser, temp, logs, cache, downloads"))
			}
		}

		fmt.Println(ui.StyleSuccess("✨ Cleanup completed!"))
	},
}

func cleanBrowser(s *shredder.Shredder, options shredder.WipeOptions) {
	fmt.Println(ui.StyleInfo("🌐 Cleaning browser data..."))
	if err := s.WipeBrowserData(options); err != nil {
		fmt.Printf(ui.StyleError("Failed to clean browser data: %v\n"), err)
	}
}

func cleanTemp(s *shredder.Shredder, options shredder.WipeOptions) {
	fmt.Println(ui.StyleInfo("📂 Cleaning temporary files..."))
	if err := s.WipeSystemTemp(options); err != nil {
		fmt.Printf(ui.StyleError("Failed to clean temp files: %v\n"), err)
	}
}

func cleanLogs(s *shredder.Shredder, options shredder.WipeOptions) {
	fmt.Println(ui.StyleInfo("📝 Cleaning log files..."))
	// Implementation would go here for log cleaning
	fmt.Println(ui.StyleMuted("Log cleaning not yet implemented"))
}

func cleanCache(s *shredder.Shredder, options shredder.WipeOptions) {
	fmt.Println(ui.StyleInfo("💾 Cleaning cache directories..."))
	// Implementation would go here for cache cleaning
	fmt.Println(ui.StyleMuted("Cache cleaning not yet implemented"))
}

func cleanDownloads(s *shredder.Shredder, options shredder.WipeOptions) {
	fmt.Println(ui.StyleWarning("⬇️ Cleaning Downloads folder..."))
	if !options.Force {
		if !ui.ConfirmDangerous("clean your Downloads folder") {
			fmt.Println(ui.StyleInfo("Downloads cleaning cancelled"))
			return
		}
	}
	fmt.Println(ui.StyleMuted("Downloads cleaning not yet implemented"))
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().Bool("dry-run", false, "Show what would be cleaned without actually doing it")
	cleanCmd.Flags().BoolP("force", "f", false, "Skip confirmation prompts")
	cleanCmd.Flags().IntP("passes", "p", 3, "Number of overwrite passes (1-35)")
} 