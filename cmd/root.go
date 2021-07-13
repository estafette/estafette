package cmd

import (
	"context"
	"os"
	"runtime"

	foundation "github.com/estafette/estafette-foundation"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	appgroup        string
	app             string
	version         string
	branch          string
	revision        string
	buildDate       string
	goVersion       = runtime.Version()
	applicationInfo = foundation.NewApplicationInfo(appgroup, app, version, branch, revision, buildDate)
)

// rootCmd represents the base command when called without any subcommands
var (
	verbose bool

	rootCmd = &cobra.Command{
		Use:          "estafette",
		Short:        "The command-line interface for Estafette",
		Long:         `Estafette is the resilient and cloud-native CI/CD platform. Read more at https://estafette.io`,
		SilenceUsage: true,
	}
)

// Execute executes the root command.
func Execute(ctx context.Context) {
	// initialize logging
	foundation.InitLoggingByFormatSilent(applicationInfo, foundation.LogFormatConsole)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(versionCmd)
}
