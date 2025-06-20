package main

import (
	"cbtg/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "cbtg"}
	rootCmd.AddCommand(cmd.InitCmd)
	rootCmd.AddCommand(cmd.TourCmd)
	rootCmd.AddCommand(cmd.VersionCmd)
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}
