package src

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type db_results struct {
	id       int
	username string
}

func sql_is_id_in_db(db_connection *pgxpool.Pool, member_id int) (
	db_username string,
	in_db bool) {
	// query
	var sql_query string = "SELECT * FROM public.users WHERE member_id = $1"

	// query results
	var query_results db_results
	err := db_connection.QueryRow(context.Background(), sql_query, member_id).Scan(
		&query_results.id, &query_results.username)

	// if no results
	if err != nil {
		// new user
		log.Info().Int("id", member_id).Msg("ID NOT in DB")
		return "", false
	}
	// rename needed
	log.Info().Int("id", member_id).Msg("ID IS in DB")
	return query_results.username, true
}
