package cmd

import (
	"fmt"
	"runtime"

	"github.com/joao-rrondon/wipeOs/internal/forensic"
	"github.com/joao-rrondon/wipeOs/ui"
	"github.com/spf13/cobra"
)

var forensicCmd = &cobra.Command{
	Use:   "forensic",
	Short: ui.IconForensic() + " Advanced anti-forensic operations",
	Long: ui.StyleHeader("🕵️‍♂️ Anti-Forensic Operations") + `

This command performs comprehensive anti-forensic cleanup to remove traces
that could be recovered by digital forensic analysis. Use with extreme caution.

⚠️  WARNING: These operations are IRREVERSIBLE and will permanently remove
system traces, logs, and metadata. Only use if you understand the consequences.

Operations performed:
• 🗂️ Clean system and application logs
• 📋 Clean Windows Registry traces  
• ⚡ Clean Prefetch execution traces
• 🖼️ Clean thumbnails and recent files
• 📋 Clear Windows Event Logs
• 🗃️ Clean Master File Table records
• 👥 Delete Volume Shadow Copies
• 🧠 Remove memory dump files
• 💾 Clean swap/page files
• 🗂️ Wipe free disk space

Examples:
  wipeOs forensic --dry-run           # Preview operations
  wipeOs forensic --all               # Full cleanup
  wipeOs forensic --logs --registry   # Selective cleanup
  wipeOs forensic --quick             # Quick essential cleanup`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get flags
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		verbose, _ := cmd.Flags().GetBool("verbose")
		all, _ := cmd.Flags().GetBool("all")
		quick, _ := cmd.Flags().GetBool("quick")
		logs, _ := cmd.Flags().GetBool("logs")
		registry, _ := cmd.Flags().GetBool("registry")
		prefetch, _ := cmd.Flags().GetBool("prefetch")
		thumbnails, _ := cmd.Flags().GetBool("thumbnails")
		eventlogs, _ := cmd.Flags().GetBool("eventlogs")
		mft, _ := cmd.Flags().GetBool("mft")
		shadows, _ := cmd.Flags().GetBool("shadows")
		memory, _ := cmd.Flags().GetBool("memory")
		swap, _ := cmd.Flags().GetBool("swap")
		freespace, _ := cmd.Flags().GetBool("freespace")
		passes, _ := cmd.Flags().GetInt("passes")

		// Show warning for non-dry runs
		if !dryRun {
			fmt.Println(ui.StyleError("🚨 DANGER: Anti-Forensic Operations"))
			fmt.Println(ui.StyleWarning("This will permanently remove system traces and cannot be undone!"))
			fmt.Println()
			
			if !ui.ConfirmDangerous("proceed with IRREVERSIBLE anti-forensic cleanup") {
				fmt.Println(ui.StyleInfo("Operation cancelled"))
				return
			}
			fmt.Println()
		}

		// Set options based on flags
		options := forensic.ForensicCleanOptions{
			DryRun:  dryRun,
			Verbose: verbose,
			Passes:  passes,
		}

		// Determine what to clean
		if all {
			// Enable everything for complete cleanup
			options.CleanLogs = true
			options.CleanRegistry = true
			options.CleanPrefetch = true
			options.CleanThumbnails = true
			options.CleanEventLogs = true
			options.CleanMFT = true
			options.CleanShadowCopies = true
			options.CleanMemory = true
			options.CleanSwap = true
			options.WipeFreespace = true
		} else if quick {
			// Quick essential cleanup
			options.CleanLogs = true
			options.CleanRegistry = true
			options.CleanThumbnails = true
			options.CleanEventLogs = true
		} else {
			// Individual flags
			options.CleanLogs = logs
			options.CleanRegistry = registry
			options.CleanPrefetch = prefetch
			options.CleanThumbnails = thumbnails
			options.CleanEventLogs = eventlogs
			options.CleanMFT = mft
			options.CleanShadowCopies = shadows
			options.CleanMemory = memory
			options.CleanSwap = swap
			options.WipeFreespace = freespace
		}

		// Check if no operations selected
		if !options.CleanLogs && !options.CleanRegistry && !options.CleanPrefetch &&
		   !options.CleanThumbnails && !options.CleanEventLogs && !options.CleanMFT &&
		   !options.CleanShadowCopies && !options.CleanMemory && !options.CleanSwap &&
		   !options.WipeFreespace {
			fmt.Println(ui.StyleError("No operations selected. Use --all, --quick, or specific flags."))
			fmt.Println(ui.StyleInfo("Run 'wipeOs forensic --help' for available options."))
			return
		}

		// Show platform warning
		if runtime.GOOS != "windows" {
			fmt.Println(ui.StyleWarning("⚠️ Some operations are Windows-specific and will be skipped"))
			fmt.Println()
		}

		// Show dry run info
		if dryRun {
			fmt.Println(ui.StyleInfo("🧪 DRY RUN MODE - No actual operations will be performed"))
			fmt.Println()
		}

		// Perform anti-forensic operations
		antiForensic := forensic.New(dryRun, verbose)
		results := antiForensic.PerformForensicCleanup(options)

		// Display results
		fmt.Println(ui.StyleHeader("📊 Anti-Forensic Operation Results:"))
		fmt.Println()

		successCount := 0
		for _, result := range results {
			if result.Success {
				successCount++
				fmt.Printf(ui.StyleSuccess("✓ %s: %s\n"), result.Operation, result.Details)
			} else {
				fmt.Printf(ui.StyleError("✗ %s: %v\n"), result.Operation, result.Error)
			}
		}

		fmt.Println()
		fmt.Printf(ui.StyleHeader("🎯 Summary: %d/%d operations completed successfully\n"), successCount, len(results))
		
		if !dryRun && successCount > 0 {
			fmt.Println()
			fmt.Println(ui.StyleSuccess("🛡️ Anti-forensic cleanup completed!"))
			fmt.Println(ui.StyleMuted("System traces have been minimized."))
		}
	},
}

func init() {
	rootCmd.AddCommand(forensicCmd)

	// Operation flags
	forensicCmd.Flags().Bool("all", false, "Perform complete anti-forensic cleanup")
	forensicCmd.Flags().Bool("quick", false, "Quick essential cleanup (logs, registry, recent files)")
	
	// Individual operation flags
	forensicCmd.Flags().Bool("logs", false, "Clean system and application logs")
	forensicCmd.Flags().Bool("registry", false, "Clean Windows Registry traces")
	forensicCmd.Flags().Bool("prefetch", false, "Clean Windows Prefetch files")
	forensicCmd.Flags().Bool("thumbnails", false, "Clean thumbnails and recent files")
	forensicCmd.Flags().Bool("eventlogs", false, "Clear Windows Event Logs")
	forensicCmd.Flags().Bool("mft", false, "Clean Master File Table records")
	forensicCmd.Flags().Bool("shadows", false, "Delete Volume Shadow Copies")
	forensicCmd.Flags().Bool("memory", false, "Remove memory dump files")
	forensicCmd.Flags().Bool("swap", false, "Clean swap/page files")
	forensicCmd.Flags().Bool("freespace", false, "Wipe free disk space")
	
	// Configuration flags
	forensicCmd.Flags().Bool("dry-run", false, "Show what would be cleaned without doing it")
	forensicCmd.Flags().BoolP("verbose", "v", false, "Show detailed operation progress")
	forensicCmd.Flags().IntP("passes", "p", 3, "Number of overwrite passes for free space wiping")
} 