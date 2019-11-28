package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Estafette CLI version=%v branch=%v revision=%v goversion=%v builddate=%v", version, branch, revision, goVersion, buildDate)
	},
}
