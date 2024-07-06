package main

import (
	"time"

	"os"

	"github.com/kanchimoe/MAL_score_scraper_go/src"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	project_config()
	src.Logic_main()
}

func project_config() {
	// set log level
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	zerolog.DurationFieldUnit = time.Second
}
