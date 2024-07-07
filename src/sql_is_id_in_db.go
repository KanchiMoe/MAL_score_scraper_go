package src

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

func sql_is_id_in_db(_ *pgxpool.Pool, db_results mal_member) {
	var username string = db_results.username

	// check if username is blank
	if username == "" {
		log.Panic().Msg("Username is blank")
	}

	// create profile page url
	var profile_page string = fmt.Sprintf(ROOTURL + "/profile/" + username)
	log.Trace().Str("url", profile_page).Msg("Profile page")

	// request the url

	page_html := Request_handler(profile_page)
	_, err := scrape_member_page(page_html)
	if err != nil {
		log.Error().Err(err).Msg("Error scraping member's page")
		return
	}

}

func scrape_member_page(page_html *html.Node) (user_id int, err error) {
	panic("foo111")

}
