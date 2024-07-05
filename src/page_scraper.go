package src

import (
	"github.com/antchfx/htmlquery"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

type mal_member struct {
	username string
}

//  (status_code int, html_node *html.Node, err error)

func Stats_scrape(page_html *html.Node) error {
	var stats_table_xpath string = "/html/body/div[1]/div[2]/div[3]/div[2]/table/tbody/tr/td[2]/div[1]/table[2]/tbody/tr"
	var member_name_xpath string = "/td[1]/div[2]/a"
	var member_score_xpath string = "/td[2]"
	var member_status_xpath string = "/td[3]"

	member_row, err := htmlquery.QueryAll(page_html, stats_table_xpath)
	if err != nil {
		log.Error().Err(err).Msg("Error when getting stats table")
		return err
	}

	for _, element := range member_row[1:] {

		// get username
		name_node, err := htmlquery.Query(element, member_name_xpath)
		if err != nil {
			log.Error().Err(err).Msg("Error when getting member name from stats table")
			return err
		}

		MAL_username_string := htmlquery.InnerText(name_node)

		// get score
		score_node, err := htmlquery.Query(element, member_score_xpath)
		if err != nil {
			log.Error().Err(err).Msg("Error when getting score from stats table")
			return err
		}

		MAL_score_string := htmlquery.InnerText(score_node)

		status_node, err := htmlquery.Query(element, member_status_xpath)
		if err != nil {
			log.Error().Err(err).Msg("Error when getting member watch status from stats table")
		}

		MAL_status_string := htmlquery.InnerText(status_node)

		log.Debug().Str("username", MAL_username_string).
			Str("score", MAL_score_string).
			Str("status", MAL_status_string).
			Msg("member")

	}

	return nil

}
