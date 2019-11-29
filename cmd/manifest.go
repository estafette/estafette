package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(manifestCmd)
}

var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "Allows to interact with the estafette manifest",
}
