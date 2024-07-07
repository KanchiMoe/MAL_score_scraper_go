package src

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func DB() (ret *pgxpool.Pool) {
	var err error

	// load dotenv
	err = godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not load .env file")
	} else {
		log.Trace().Msg("Was able to load .env file")
	}

	// make db connection
	ctx := context.Background()
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not create database pool")
	}
	ret, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not create database pool")
	}
	err = ret.Ping(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not create database pool")
	}
	return
}
