package cmd

import (
	"runtime"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().
			Str("branch", branch).
			Str("revision", revision).
			Str("buildDate", buildDate).
			Str("goVersion", goVersion).
			Str("os", runtime.GOOS).
			Msgf("Estafette command-line interface v%v", version)
	},
}
