package cmd

import (
	"context"
	"strings"

	"github.com/estafette/estafette-ci-builder/pkg/builder"
	contracts "github.com/estafette/estafette-ci-contracts"
	crypt "github.com/estafette/estafette-ci-crypt"
	foundation "github.com/estafette/estafette-foundation"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build [comma separated list of stages to run; defaults to first stage]",
	Short: "Build .estafette.yaml manifest locally",
	RunE: func(cmd *cobra.Command, args []string) error {

		// handle cancellation
		ctx := foundation.InitCancellationContext(context.Background())

		// init secret helper
		decryptionKey := ""
		secretHelper := crypt.NewSecretHelper(decryptionKey, false)

		// bootstrap
		ciBuilder := builder.NewCIBuilder(applicationInfo)
		tailLogsChannel := make(chan contracts.TailLogLine, 10000)
		obfuscator := builder.NewObfuscator(secretHelper)
		envvarHelper := builder.NewEnvvarHelper("ESTAFETTE_", secretHelper, obfuscator)
		whenEvaluator := builder.NewWhenEvaluator(envvarHelper)
		var builderConfig contracts.BuilderConfig

		// ensure GetPipelineName does not fail below
		err := envvarHelper.SetPipelineName(builderConfig)
		if err != nil {
			return err
		}

		containerRunner := builder.NewDockerRunner(envvarHelper, obfuscator, builderConfig, tailLogsChannel, false)
		pipelineRunner := builder.NewPipelineRunner(envvarHelper, whenEvaluator, containerRunner, false, tailLogsChannel, applicationInfo)

		// get stages to run from arguments
		stagesToRun := []string{}
		if len(args) > 0 {
			stagesToRun = strings.Split(args[0], ",")
		}

		err = ciBuilder.RunLocalBuild(ctx, pipelineRunner, containerRunner, envvarHelper, contracts.BuilderConfig{}, stagesToRun)
		if err != nil {
			return err
		}

		return nil
	},
}
