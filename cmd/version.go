package cmd

import (
	"fmt"
	"runtime"

	"github.com/joao-rrondon/wipeOs/ui"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  "Display detailed version information including build details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ui.RenderWelcomeBanner())
		fmt.Println()
		
		fmt.Printf("%s\n", ui.StyleHeader("Version Information"))
		fmt.Printf("Version:    %s\n", ui.StyleInfo(version))
		fmt.Printf("Commit:     %s\n", ui.StyleMuted(commit))
		fmt.Printf("Built:      %s\n", ui.StyleMuted(date))
		fmt.Printf("Go version: %s\n", ui.StyleMuted(runtime.Version()))
		fmt.Printf("Platform:   %s/%s\n", ui.StyleMuted(runtime.GOOS), ui.StyleMuted(runtime.GOARCH))
		
		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			fmt.Printf("\n%s\n", ui.StyleHeader("Build Details"))
			fmt.Printf("Compiler:   %s\n", ui.StyleMuted(runtime.Compiler))
			fmt.Printf("NumCPU:     %s\n", ui.StyleMuted(fmt.Sprintf("%d", runtime.NumCPU())))
			fmt.Printf("GOMAXPROCS: %s\n", ui.StyleMuted(fmt.Sprintf("%d", runtime.GOMAXPROCS(0))))
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolP("verbose", "v", false, "Show detailed build information")
} 