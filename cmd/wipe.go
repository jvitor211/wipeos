package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/joao-rrondon/wipeOs/internal/shredder"
	"github.com/joao-rrondon/wipeOs/ui"
	"github.com/spf13/cobra"
)

var wipeCmd = &cobra.Command{
	Use:   "wipe [files...]",
	Short: ui.IconWipe() + "  Securely wipe files and directories",
	Long: ui.StyleHeader("Secure File Wiping") + `

This command securely overwrites and deletes the specified files using
military-grade overwriting patterns to ensure data cannot be recovered.

Examples:
  wipeOs wipe secret.txt                    # Wipe a single file
  wipeOs wipe *.log --recursive             # Wipe all .log files recursively
  wipeOs wipe /tmp/sensitive/ --recursive   # Wipe entire directory
  wipeOs wipe --browser-data                # Wipe browser cache/history
  wipeOs wipe --system-temp                 # Clean system temporary files

⚠️  WARNING: This operation is IRREVERSIBLE!`,
	Args: func(cmd *cobra.Command, args []string) error {
		browserData, _ := cmd.Flags().GetBool("browser-data")
		systemTemp, _ := cmd.Flags().GetBool("system-temp")
		
		if len(args) == 0 && !browserData && !systemTemp {
			return fmt.Errorf("specify files to wipe or use predefined options (--browser-data, --system-temp)")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		recursive, _ := cmd.Flags().GetBool("recursive")
		passes, _ := cmd.Flags().GetInt("passes")
		force, _ := cmd.Flags().GetBool("force")
		browserData, _ := cmd.Flags().GetBool("browser-data")
		systemTemp, _ := cmd.Flags().GetBool("system-temp")
		dryRun, _ := cmd.Flags().GetBool("dry-run")

		options := shredder.WipeOptions{
			Recursive: recursive,
			Passes:    passes,
			Force:     force,
			DryRun:    dryRun,
		}

		s := shredder.New()

		if browserData {
			fmt.Println(ui.StyleWarning("🌐 Wiping browser data..."))
			if err := s.WipeBrowserData(options); err != nil {
				fmt.Printf(ui.StyleError("Failed to wipe browser data: %v\n"), err)
				return
			}
		}

		if systemTemp {
			fmt.Println(ui.StyleWarning("🗂️  Wiping system temporary files..."))
			if err := s.WipeSystemTemp(options); err != nil {
				fmt.Printf(ui.StyleError("Failed to wipe system temp: %v\n"), err)
				return
			}
		}

		if len(args) > 0 {
			var targets []string
			for _, arg := range args {
				if strings.Contains(arg, "*") {
					matches, err := filepath.Glob(arg)
					if err != nil {
						fmt.Printf(ui.StyleError("Invalid pattern '%s': %v\n"), arg, err)
						continue
					}
					targets = append(targets, matches...)
				} else {
					targets = append(targets, arg)
				}
			}

			if len(targets) == 0 {
				fmt.Println(ui.StyleWarning("No files matched the specified patterns"))
				return
			}

			if !force && !ui.ConfirmDangerous(fmt.Sprintf("wipe %d file(s)", len(targets))) {
				fmt.Println(ui.StyleInfo("Operation cancelled"))
				return
			}

			fmt.Printf(ui.StyleInfo("🧹 Wiping %d file(s) with %d passes...\n"), len(targets), passes)
			
			results := s.WipeFiles(targets, options)
			
			successCount := 0
			for _, result := range results {
				if result.Success {
					successCount++
					fmt.Printf(ui.StyleSuccess("✓ %s\n"), result.Path)
				} else {
					fmt.Printf(ui.StyleError("✗ %s: %v\n"), result.Path, result.Error)
				}
			}

			fmt.Printf(ui.StyleHeader("\n📊 Summary: %d/%d files wiped successfully\n"), successCount, len(results))
		}
	},
}

func init() {
	rootCmd.AddCommand(wipeCmd)

	wipeCmd.Flags().BoolP("recursive", "r", false, "Wipe directories recursively")
	wipeCmd.Flags().IntP("passes", "p", 3, "Number of overwrite passes (1-35)")
	wipeCmd.Flags().BoolP("force", "f", false, "Skip confirmation prompts")
	wipeCmd.Flags().Bool("browser-data", false, "Wipe browser cache, history, and temp files")
	wipeCmd.Flags().Bool("system-temp", false, "Wipe system temporary files")
	wipeCmd.Flags().Bool("dry-run", false, "Show what would be wiped without actually doing it")
} 