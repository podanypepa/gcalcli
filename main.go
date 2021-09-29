package main

import (
	"github.com/podanypepa/gcalcli/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("cannot start gcalcli")
	}
}
