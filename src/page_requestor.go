package src

import (
	"net/http"

	"github.com/antchfx/htmlquery"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

func Request_handler(requested_url string) *html.Node {

	// request the url
	status_code, html, err := get_url(requested_url)

	// error handling
	if err != nil {
		log.Error().Err(err).Msg("Error when requesting url")
	}

	if status_code == 200 {
		return html
	} else if status_code == 405 {
		log.Warn().Msg("We are being rate limited")
		// trigger backoff
		return nil

	} else if status_code == 504 {
		log.Warn().Msg("Got 504 (gateway timeout). We might be being rate limited, or MAL has an error. Waiting before retrying...")
		// trigger backoff
		return nil
	} else {
		log.Fatal().Int("code", status_code).Msg("HTTP error, no catch")
		return nil
	}
}

func get_url(requested_url string) (status_code int, html_node *html.Node, err error) {
	log.Info().Msg("Requesting url: " + requested_url)

	// request the url
	resp, err := http.Get(requested_url)

	// error handling
	if err != nil {
		log.Error().Err(err).Msg("Error when trying to get url")
		return 0, nil, err
	}
	defer resp.Body.Close()
	log.Debug().Int("code", resp.StatusCode).Msg("Response")

	// if response code is 200
	// return html no
	if resp.StatusCode == 200 {
		html_node, err = htmlquery.Parse(resp.Body)

		// error handling
		if err != nil {
			log.Error().Err(err).Msg("Error when parsing response body")
		}
	}

	return resp.StatusCode, html_node, err
}
