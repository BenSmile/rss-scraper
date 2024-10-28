package config

import "github.com/bensmile/rssaggregator/internal/database"

type ApiConfig struct {
	DB *database.Queries
}
