//package cmd
//
//import (
//	"fmt"
//	"github.com/spf13/cobra"
//)
//
//const Version = "v0.1.0"
//
//var VersionCmd = &cobra.Command{
//	Use:   "version",
//	Short: "Print the version of CBTG",
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Printf("CBTG CLI version %s\n", Version)
//	},
//}

package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var CliVersion = "0.1.0"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display versions of CBTG CLI and Core",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("🧭 CBTG CLI version: %s\n", CliVersion)

		// Adjust this to match the actual path of get_version.py
		script := "../../PycharmProjects/cbtg-core/get_version.py"
		out, err := exec.Command("python", script).CombinedOutput()
		if err != nil {
			fmt.Println("❌ Error getting CBTG Core version:", err)
			return
		}
		fmt.Printf("🧠 CBTG Core version: %s", out)
	},
}
