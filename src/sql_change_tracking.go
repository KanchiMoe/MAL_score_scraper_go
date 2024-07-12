package src

import (
	"fmt"
	"strings"
	"time"

	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func sql_change_tracking(
	// inputs
	db_connection *pgxpool.Pool,
	member_id int,
	action string,
	field string,
	old string,
	new string) {

	uuid := create_uuid()

	fmt.Println(uuid)

	timestamp, err := get_time()
	if err != nil {
		panic("no time")
	}

	// query
	const change_tracking_query string = `INSERT INTO scores.change_tracking
		(uuid, timestamp, member_id, action, field, old_value, new_value)
		VALUES ($1, $2, $3, $4, $5, $6, $7);`

	// query results
	_, err = db_connection.Exec(context.Background(), change_tracking_query,
		// args
		uuid,
		timestamp,
		member_id,
		strings.ToUpper(action),
		field,
		old,
		new,
	)

	if err != nil {
		panic("error writing change tracking")
	}

	// cursor.execute("""
	//     INSERT INTO scores.change_tracking
	//     (uuid, timestamp, member_id, action, field, old_value, new_value)
	//     VALUES (%s, %s, %s, %s, %s, %s, %s)
	//     """, (str(random_uuid), datetime.now(pytz.timezone('Europe/London')), member_id, action.upper(), field, old, new)
	//     )

	//panic("end of sql")
}

func create_uuid() (generated_uuid string) {
	generated_uuid = uuid.New().String()
	return generated_uuid
}

func get_time() (timestamp string, err error) {
	location, err := time.LoadLocation("Europe/London")
	if err != nil {
		panic("no location")
	}
	current_time := time.Now().In(location)

	const layout = "2006-01-02 15:04:05.999999-07"
	formatted := current_time.Format(layout)

	return formatted, err
}
