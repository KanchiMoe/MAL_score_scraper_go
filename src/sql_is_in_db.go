package src

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func sql_is_in_db(db_connection *pgxpool.Pool, member_object mal_member) (db_results mal_member, user_in_db bool) {
	var username string = member_object.username
	log.Debug().Msg("")
	log.Trace().Str("member", username).Msg("In DB check")

	// query
	var sql_query string = "SELECT * FROM scores.scores WHERE member_username = $1;"

	// query results
	err := db_connection.QueryRow(context.Background(), sql_query, username).Scan(&db_results.id,
		&db_results.username,
		&db_results.score,
		&db_results.status,
		&db_results.eps_seen)

	// If no results
	if err != nil {
		log.Info().Str("username", username).Msg("Username NOT in database")
		user_in_db = false
		return mal_member{username: username}, user_in_db
	}

	// debug logging
	log.Debug().Str("username", username).Msg("Username IS in database")
	log.Debug().Int("id", db_results.id).
		Str("username", db_results.username).
		Str("score", db_results.score).
		Str("status", db_results.status).
		Str("eps seen", db_results.eps_seen).
		Msg("Data from DB")

	user_in_db = true
	return db_results, user_in_db
}
