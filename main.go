package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/estafette/estafette/cmd"
)

func main() {

	// https://pace.dev/blog/2020/02/17/repond-to-ctrl-c-interrupt-signals-gracefully-with-context-in-golang-by-mat-ryer.html
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		// <-signalChan // second signal, hard exit
		// os.Exit(2)
	}()

	if err := cmd.Execute(ctx); err != nil {
		log.Fatal(err)
	}
}
