package cmd

import (
	manifest "github.com/estafette/estafette-ci-manifest"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	manifestCmd.AddCommand(validateValidateCmd)
}

var validateValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates the .estafette.yaml manifest in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := manifest.ReadManifestFromFile(".estafette.yaml")
		if err != nil {
			log.Fatal().Err(err).Msg("The .estafette.yaml file is not valid")
		}
		log.Info().Msg("The .estafette.yaml file is valid!")
	},
}
