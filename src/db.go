package src

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func DB() (ret *pgxpool.Pool) {
	var err error
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
