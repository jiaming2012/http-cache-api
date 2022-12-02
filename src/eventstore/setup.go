package eventstore

import (
	"fmt"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/jiaming2012/http-cache-api/src/constants"
)

func NewDB() (*esdb.Client, error) {
	settings, err := esdb.ParseConnectionString(constants.EsdbURL)
	if err != nil {
		return nil, fmt.Errorf("esdb.ParseConnectionString: %w", err)
	}

	db, err := esdb.NewClient(settings)
	if err != nil {
		return nil, fmt.Errorf("esdb.NewClient: %w", err)
	}

	return db, nil
}
