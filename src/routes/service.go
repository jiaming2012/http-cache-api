package routes

import (
	"encoding/json"
	"fmt"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/gin-gonic/gin"
	"github.com/jiaming2012/http-cache-api/src/constants"
	"github.com/jiaming2012/http-cache-api/src/eventstore"
	"github.com/jiaming2012/http-cache-api/src/types"
	"io"
	"net/http"
)

var db *esdb.Client

func init() {
	_db, err := eventstore.NewDB()
	if err != nil {
		panic(err)
	}
	db = _db
}

func CreateDeleteCachedEntityEvent(c *gin.Context) {
	var delCachedEntityRequest types.DeleteCachedEntityRequest

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("failed to read request"))
	}

	if err = json.Unmarshal(jsonData, &delCachedEntityRequest); err != nil {
		c.AbortWithError(400, fmt.Errorf("failed to read request"))
	}

	fmt.Println(delCachedEntityRequest)
	deleteCachedEntity := types.DeleteCachedEntity{
		Key:     delCachedEntityRequest.URL,
		Version: constants.HttpCacheStreamVersion,
	}

	if err = deleteCachedEntity.AppendToStream(db); err != nil {
		c.AbortWithError(500, fmt.Errorf("failed to process request"))
	}

	c.JSON(http.StatusOK, nil)
}
