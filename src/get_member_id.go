package src

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/antchfx/htmlquery"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

func get_member_id_from_mal(db_results mal_member) (user_id int, err error) {
	var username string = db_results.username

	// check if username is blank
	if username == "" {
		log.Panic().Msg("Username is blank")
		return 0, err
	}

	// create profile page url
	var profile_page string = fmt.Sprintf(ROOTURL + "/profile/" + username)
	log.Trace().Str("url", profile_page).Msg("Profile page")

	// request the url

	page_html := Request_handler(profile_page)
	user_id, err = scrape_member_page(page_html)
	if err != nil {
		log.Error().Err(err).Msg("Error scraping member's page")
		return 0, err
	}

	return user_id, nil

}

func scrape_member_page(page_html *html.Node) (user_id int, err error) {
	var report_link_xpath string = "/html/body/div[1]/div[2]/div[3]/div[1]/h1/a"

	// get member's id
	id_node, err := htmlquery.Query(page_html, report_link_xpath)
	if err != nil {
		log.Error().Err(err).Msg("Error getting member's id")
		return 0, err
	}
	if id_node == nil {
		log.Error().Err(err).Msg("ID node empty")
		return 0, err
	}

	// get ahref
	report_href := htmlquery.SelectAttr(id_node, "href")
	if report_href == "" {
		log.Error().Err(err).Msg("ahref node empty")
		return 0, nil
	}

	// parse the url
	parsed_url, err := url.Parse(report_href)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse URL")
		return 0, err
	}

	// get ID from the url
	user_id_string := parsed_url.Query().Get("id")
	if user_id_string == "" {
		log.Error().Err(err).Msg("No id parameter found in the URL")
		return 0, nil
	}

	user_id_int, err := strconv.Atoi(user_id_string)
	if err != nil {
		log.Error().Str("string", user_id_string).Err(err).Msg("Unable to convert id string to int")
		return 0, err
	}

	log.Trace().Int("id", user_id_int).Msg("Got member ID")
	return user_id_int, nil
}
