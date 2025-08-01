package cmd

import (
	"fmt"

	"github.com/joao-rrondon/wipeOs/ui"
	"github.com/spf13/cobra"
)

var retroCmd = &cobra.Command{
	Use:   "retro",
	Short: "⚡ Quick switch to retro theme",
	Long: `Quickly switch to the awesome retro gaming theme and start WipeOs.
	
This combines setting the retro icon pack and launching interactive mode
for the ultimate 80s/90s gaming experience.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set retro theme
		if err := ui.SetIconPack("retro"); err != nil {
			fmt.Printf(ui.StyleError("Error setting retro theme: %v\n"), err)
			return
		}

		fmt.Println(ui.StyleSuccess("⚡ Retro theme activated!"))
		fmt.Println(ui.StyleInfo("★ Welcome to the 80s/90s gaming vibe! ★"))
		fmt.Println()

		// Launch interactive mode with retro theme
		interactiveCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(retroCmd)
} 