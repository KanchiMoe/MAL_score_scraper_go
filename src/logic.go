package src

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

const ROOTURL = "https://myanimelist.net"
const MAX_OFFSET = 7425
const SLEEP_PERIOD time.Duration = 10 * time.Second

func Logic_main() {
	var offset int = 0
	var stats_slug string = "/anime/6547/Angel_Beats/stats?show="

	for offset <= MAX_OFFSET {
		// create url
		var current_url string = fmt.Sprintf("%s%s%d", ROOTURL, stats_slug, offset)

		// request the page
		var page_html = Request_handler(current_url)

		// scrape the page, get list of members
		list_of_members, err := Stats_scrape(page_html)
		if err != nil {
			log.Panic().Err(err).Msg("Error returned from stats page scrape")
		}
		count_list_of_members := len(list_of_members)
		log.Trace().Int("count", count_list_of_members).Msg("No. of members returned from scrape")
		if count_list_of_members != 75 {
			log.Error().Int("count", count_list_of_members).Msg("No of members not 75")
			return
		}
		//log.Trace().Interface("contents", list_of_members).Msg("List of members") - not working
		//fmt.Println(list_of_members) - placeholder

		// check to see if these are in the db already
		db_connection := DB()
		for _, member_object := range list_of_members {
			db_results, user_in_db := sql_is_in_db(db_connection, member_object)

			// username not in db
			// check to see if their ID is in the db
			if !user_in_db {

				// get the member id first
				member_id, err := get_member_id_from_mal(db_results)
				if err != nil {
					log.Error().Err(err).Msg("Error trying to get member ID")
				}

				// check if the id is in the db
				_, id_in_db := sql_is_id_in_db(db_connection, member_id)

				// the ID is in the db
				if id_in_db {
					// actions
					continue
				}

				// the ID is not in the db
				sql_new_to_db(db_connection, member_id, member_object)

				//panic("stop")
			}

			//panic("stop")
		}

		// increase offset by 75
		offset += 75

		// sleep
		log.Debug().Str("duration", SLEEP_PERIOD.String()).Msg("Sleeping")
		time.Sleep(SLEEP_PERIOD)

	}

}
