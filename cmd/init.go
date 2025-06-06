package cmd

import (
	"cbtg/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var InitCmd = &cobra.Command{
	Use:   "init [targetPath]",
	Short: "Scan and analyze the codebase at the given path",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetPath := args[0]

		absPath, err := utils.ValidatePath(targetPath)
		if err != nil {
			fmt.Println("❌", err)
			return
		}

		fmt.Printf("📁 Initializing CBTG at: %s\n", absPath)

		// Define the script and command
		script := "../../PycharmProjects/cbtg-core/run_summary.py"
		pythonCmd := exec.Command("python", script, absPath)

		// Print the exact command being executed
		fmt.Printf("🛠️  Running: python %s \"%s\"\n", script, absPath)

		// Execute the command
		out, err := pythonCmd.CombinedOutput()
		if err != nil {
			fmt.Println("❌ Error running Python script:", err)
			fmt.Println("📤 Output:", string(out)) // also print Python stderr
			return
		}

		// Print the result
		fmt.Println(string(out))
	},
}
