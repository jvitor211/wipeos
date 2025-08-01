package cmd

import (
	"fmt"

	"github.com/joao-rrondon/wipeOs/ui"
	"github.com/spf13/cobra"
)

var iconsCmd = &cobra.Command{
	Use:   "icons",
	Short: "🎨 Manage icon packs and themes",
	Long: ui.StyleHeader("🎨 Icon Pack Management") + `

Customize the visual appearance of WipeOs with different icon themes.
Choose from various aesthetic styles to match your preference.

Available commands:
• list    - Show all available icon packs
• set     - Change to a specific icon pack  
• current - Show currently active icon pack
• preview - Preview an icon pack

Examples:
  wipeOs icons list                 # List all packs
  wipeOs icons set cyber            # Switch to cyber theme
  wipeOs icons preview military     # Preview military theme
  wipeOs icons current              # Show current pack`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// Default to list command
			listIconPacks()
			return
		}

		switch args[0] {
		case "list", "ls":
			listIconPacks()
		case "set":
			if len(args) < 2 {
				fmt.Println(ui.StyleError("Usage: wipeOs icons set <pack-name>"))
				fmt.Println(ui.StyleInfo("Run 'wipeOs icons list' to see available packs"))
				return
			}
			setIconPack(args[1])
		case "current":
			showCurrentPack()
		case "preview":
			if len(args) < 2 {
				fmt.Println(ui.StyleError("Usage: wipeOs icons preview <pack-name>"))
				return
			}
			previewIconPack(args[1])
		default:
			fmt.Printf(ui.StyleError("Unknown subcommand: %s\n"), args[0])
			fmt.Println(ui.StyleInfo("Available: list, set, current, preview"))
		}
	},
}

func listIconPacks() {
	fmt.Println(ui.StyleHeader("🎨 Available Icon Packs:"))
	fmt.Println()

	current := ui.GetCurrentIconPack()
	packs := ui.GetAllIconPacks()

	for _, pack := range packs {
		marker := " "
		if pack.Name == current.Name {
			marker = ui.StyleSuccess("▶")
		}

		fmt.Printf("%s %s - %s\n",
			marker,
			ui.StyleInfo(pack.Name),
			ui.StyleMuted(pack.Description))

		// Show a few sample icons
		fmt.Printf("    %s %s %s %s %s\n",
			pack.Icons.Wipe,
			pack.Icons.Clean,
			pack.Icons.Forensic,
			pack.Icons.Success,
			pack.Icons.Warning)
		fmt.Println()
	}

	fmt.Println(ui.StyleInfo("Use 'wipeOs icons set <pack-name>' to change theme"))
}

func setIconPack(packName string) {
	if err := ui.SetIconPack(packName); err != nil {
		fmt.Printf(ui.StyleError("Error: %v\n"), err)
		fmt.Println(ui.StyleInfo("Run 'wipeOs icons list' to see available packs"))
		return
	}

	fmt.Printf(ui.StyleSuccess("✓ Icon pack changed to: %s\n"), packName)
	fmt.Println(ui.StyleMuted("Restart WipeOs or start a new session to see changes"))
}

func showCurrentPack() {
	current := ui.GetCurrentIconPack()
	
	fmt.Printf(ui.StyleHeader("Current Icon Pack: %s\n"), current.Name)
	fmt.Printf(ui.StyleMuted("Description: %s\n"), current.Description)
	fmt.Println()

	fmt.Println(ui.StyleInfo("Sample icons:"))
	fmt.Printf("Wipe: %s  Clean: %s  Forensic: %s  Success: %s  Error: %s\n",
		current.Icons.Wipe,
		current.Icons.Clean,
		current.Icons.Forensic,
		current.Icons.Success,
		current.Icons.Error)
}

func previewIconPack(packName string) {
	// Temporarily set the pack to preview
	originalPack := ui.GetCurrentIconPack()
	
	if err := ui.SetIconPack(packName); err != nil {
		fmt.Printf(ui.StyleError("Error: %v\n"), err)
		return
	}
	
	pack := ui.GetCurrentIconPack()
	
	fmt.Printf(ui.StyleHeader("Preview: %s\n"), pack.Name)
	fmt.Printf(ui.StyleMuted("%s\n"), pack.Description)
	fmt.Println()

	// Show comprehensive preview
	fmt.Println(ui.StyleInfo("Operations:"))
	fmt.Printf("  %s Wipe    %s Clean    %s Forensic    %s Start\n",
		pack.Icons.Wipe, pack.Icons.Clean, pack.Icons.Forensic, pack.Icons.Start)
	
	fmt.Println(ui.StyleInfo("Results:"))
	fmt.Printf("  %s Success    %s Error    %s Warning    %s Info\n",
		pack.Icons.Success, pack.Icons.Error, pack.Icons.Warning, pack.Icons.Info)
	
	fmt.Println(ui.StyleInfo("Security:"))
	fmt.Printf("  %s Shield    %s Lock    %s Danger    %s Safe\n",
		pack.Icons.Shield, pack.Icons.Lock, pack.Icons.Danger, pack.Icons.Safe)

	fmt.Println(ui.StyleInfo("Forensic Operations:"))
	fmt.Printf("  %s Logs    %s Registry    %s Memory    %s Shadows\n",
		pack.Icons.SystemLogs, pack.Icons.Registry, pack.Icons.Memory, pack.Icons.Shadows)
	
	fmt.Println()
	fmt.Printf(ui.StyleInfo("Use 'wipeOs icons set %s' to apply this theme\n"), packName)
	
	// Restore original pack
	ui.SetIconPack(originalPack.Name)
}

func init() {
	rootCmd.AddCommand(iconsCmd)
} 