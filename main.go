package main

import (
	"fmt"
	"time"

	"os"

	"github.com/kanchimoe/MAL_score_scraper_go/src"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const ROOTURL = "https://myanimelist.net"
const MAX_OFFSET = 7425
const SLEEP_PERIOD time.Duration = 1 * time.Second

func main() {
	project_config()
	var offset int = 0
	var stats_slug string = "/anime/6547/Angel_Beats/stats?show="

	for offset <= MAX_OFFSET {
		// create url
		var current_url string = fmt.Sprintf("%s%s%d", ROOTURL, stats_slug, offset)

		// request the page
		//var page_html = src.Request_handler(current_url)
		var _ = src.Request_handler(current_url)

		// increase offset by 75
		offset += 75

		// sleep
		log.Debug().Str("duration", SLEEP_PERIOD.String()).Msg("Sleeping")
		time.Sleep(SLEEP_PERIOD)
	}
}

func project_config() {
	// set log level
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	zerolog.DurationFieldUnit = time.Second
}
