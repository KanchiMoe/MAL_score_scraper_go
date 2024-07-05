package src

import (
	"github.com/antchfx/htmlquery"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"

	"regexp"
	"strings"
)

type mal_member struct {
	username string
	score    string
	status   string
	eps_seen string
}

func Stats_scrape(page_html *html.Node) (member mal_member, err error) {
	var stats_table_xpath string = "/html/body/div[1]/div[2]/div[3]/div[2]/table/tbody/tr/td[2]/div[1]/table[2]/tbody/tr"
	var member_name_xpath string = "/td[1]/div[2]/a"
	var member_score_xpath string = "/td[2]"
	var member_status_xpath string = "/td[3]"
	var member_eps_seen_xpath string = "/td[4]"

	member_row, err := htmlquery.QueryAll(page_html, stats_table_xpath)
	if err != nil {
		log.Error().Err(err).Msg("Error getting stats table")
		return member, err
	}

	for _, element := range member_row[1:] {

		// get username
		name_node, err := htmlquery.Query(element, member_name_xpath)
		if err != nil {
			log.Error().Err(err).Msg("Error getting member name from stats table")
			return member, err
		}
		MAL_username_string := htmlquery.InnerText(name_node)

		// get score
		score_node, err := htmlquery.Query(element, member_score_xpath)
		if err != nil {
			log.Error().Err(err).Msg("Error getting score from stats table")
			return member, err
		}
		MAL_score_string := htmlquery.InnerText(score_node)

		status_node, err := htmlquery.Query(element, member_status_xpath)
		if err != nil {
			log.Error().Err(err).Msg("Error getting watch status from stats table")
			return member, err
		}
		MAL_status_string := htmlquery.InnerText(status_node)

		// get eps seen
		eps_node, err := htmlquery.Query(element, member_eps_seen_xpath)
		if err != nil {
			log.Error().Err(err).Msg("Error getting eps seen from stats table")
			return member, err
		}
		raw_eps_seen_string := htmlquery.InnerText(eps_node)

		// remove white spaces
		eps_no_whitespace := strings.TrimSpace(raw_eps_seen_string)
		eps_regex := regexp.MustCompile(`^\d+`)
		MAL_eps_seen_string := eps_regex.FindString(eps_no_whitespace)

		// set as 0 if string empty
		if MAL_eps_seen_string == "" {
			MAL_eps_seen_string = "0"
		}

		// log.Debug().Str("username", MAL_username_string).
		// 	Str("score", MAL_score_string).
		// 	Str("status", MAL_status_string).
		// 	Str("eps seen", MAL_eps_seen_string).
		// 	Msg("Member")

		member.username = MAL_username_string
		member.score = MAL_score_string
		member.status = MAL_status_string
		member.eps_seen = MAL_eps_seen_string
	}

	return member, nil

}
