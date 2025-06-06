package cmd

import (
	"cbtg/utils"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var format string

var TourCmd = &cobra.Command{
	Use:   "tour [path]",
	Short: "Generate a codebase tour",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		absPath, err := utils.ValidatePath(path)
		if err != nil {
			fmt.Println("❌", err)
			return
		}

		fmt.Printf("📚 Generating tour for: %s (format: %s)\n\n", absPath, format)

		script := "../../PycharmProjects/cbtg-core/run_tour.py"
		pythonCmd := exec.Command("python", script, absPath, "--format", format)

		fmt.Printf("🛠️  Running: %v\n", pythonCmd.Args)

		out, err := pythonCmd.CombinedOutput()
		fmt.Printf("📤 Output:\n%s\n", string(out))

		if err != nil {
			fmt.Println("❌ Error running tour script:", err)
			return
		}

		fmt.Println("✅ Code tour complete!")
	},
}

func init() {
	TourCmd.Flags().StringVarP(&format, "format", "f", "text", "Output format: text | md")
}
