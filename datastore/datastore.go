package datastore

import (
	"errors"
	"os"

	"github.com/arkrozycki/reunion/logger"
	"github.com/syndtr/goleveldb/leveldb"
)

var log = logger.Get()
var (
	ErrNotFound = errors.New("not found")
)

type Datastore interface {
	GetUser(key string) (string, error)
	GetCode(key string) (string, error)
	SaveUser(key string, user interface{}) error
	SaveCode(key string, v interface{}) error
}

func GetDatastore() Datastore {
	log.Debug().Msgf("GetDatastore() %s", os.Getenv("Datastore"))
	switch os.Getenv("Datastore") {
	case "leveldb":
		db := &LevelDB{
			File: os.Getenv("LocalDbFilename"),
			db:   &leveldb.DB{},
		}
		return db
	default:
		return nil
	}
}
