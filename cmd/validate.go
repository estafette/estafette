package cmd

import (
	"fmt"

	manifest "github.com/estafette/estafette-ci-manifest"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates the .estafette.yaml manifest in the current directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := manifest.ReadManifestFromFile(nil, ".estafette.yaml", true)
		if err != nil {
			return fmt.Errorf("The .estafette.yaml file is not valid: %w", err)
		}
		log.Info().Msg("The .estafette.yaml file is valid!")

		return nil
	},
}
