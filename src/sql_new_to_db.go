package src

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func sql_new_to_db(db_connection *pgxpool.Pool, member_id int, member_object mal_member) {
	fmt.Println("member id")
	fmt.Println(member_id)

	fmt.Println("member username")
	fmt.Println(member_object.username)

	////////////////

	// query
	const sql_inster_into_scores string = "INSERT INTO scores.scores (member_id, member_username, score, status, eps_seen) VALUES ($1, $2, $3, $4, $5);"

	// query results
	_, err := db_connection.Exec(context.Background(), sql_inster_into_scores, member_id,
		member_object.username,
		member_object.score,
		member_object.status,
		member_object.eps_seen)

	if err != nil {
		panic("could not enter into scores")
	}

	// query
	const sql_insert_into_users string = "INSERT INTO users (member_id, member_username) VALUES ($1, $2);"

	// query results
	_, err = db_connection.Exec(context.Background(), sql_insert_into_users, member_id, member_object.username)
	if err != nil {
		panic("could not enter into users")
	}

	sql_change_tracking(db_connection, member_id, "create", "ALL", "NULL", "n/a")

	//panic("dsgfgd")

}

// cursor.execute("""
// INSERT INTO users
// (member_id, member_username)
// VALUES (%s, %s)
// """, (member["id"], member["name"])
// )
