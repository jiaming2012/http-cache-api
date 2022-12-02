package types

import (
	"context"
	"encoding/json"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/jiaming2012/http-cache-api/src/constants"
)

type DeleteCachedEntityRequest struct {
	URL string `json:"url"`
}

// DeleteCachedEntity deletes a stored cached response
type DeleteCachedEntity struct {
	Key     string
	Version uint
}

func (entity *DeleteCachedEntity) AppendToStream(db *esdb.Client) error {
	data, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	eventData := esdb.EventData{
		ContentType: esdb.JsonContentType,
		EventType:   constants.DeleteCachedEntityType,
		Data:        data,
	}

	_, err = db.AppendToStream(context.Background(), constants.HttpCacheStreamName, esdb.AppendToStreamOptions{}, eventData)
	if err != nil {
		return err
	}

	return nil
}
