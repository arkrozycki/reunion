package datastore

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	conf "github.com/arkrozycki/reunion/config"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

// LevelDB struct
type LevelDB struct {
	File string
	db   *leveldb.DB
	open bool
}

func init() {
	conf.Config()

}

// Open method
func (l *LevelDB) Open() *leveldb.DB {
	var err error
	// return singleton if already connected
	if l.db != nil && l.open {
		return l.db
	}

	log.Info().Msgf("open/create leveldb %s", l.File)

	if l.File == "mem" {
		stor := storage.NewMemStorage()
		l.db, _ = leveldb.Open(stor, nil)
		l.open = true
		return l.db
	}

	// if we are using a local filename
	abs, err := filepath.Abs(l.File)
	if err != nil {
		log.Fatal().Err(err)
		return nil
	}

	// open database at absolute filepath
	l.db, err = leveldb.OpenFile(abs, nil)
	if err != nil {
		log.Fatal().Err(err)
		return nil
	}
	l.open = true
	// return singleton
	return l.db

}

// GetUser method
func (l *LevelDB) GetUser(key string) (string, error) {
	log.Debug().Msg("LevelDb.Get()")
	d := l.Open()
	record, err := d.Get([]byte("u:"+key), nil)
	if err != nil && err != leveldb.ErrNotFound {
		fmt.Println(err)
		return "", err
	}
	return string(record), nil
}

// GetCode method
func (l *LevelDB) GetCode(key string) (string, error) {
	log.Debug().Msg("LevelDb.Get()")
	d := l.Open()
	record, err := d.Get([]byte("c:"+key), nil)
	if err == leveldb.ErrNotFound {
		return "", ErrNotFound
	}
	if err != nil {
		return "", err
	}
	return string(record), nil
}

// SaveUser method
func (l *LevelDB) SaveUser(key string, user interface{}) error {
	log.Debug().Msg("LevelDb.SaveUser()")

	tmp, err := json.Marshal(user)
	if err != nil {
		return err
	}
	l.Open()
	err = l.db.Put([]byte("u:"+key), []byte(string(tmp)), nil)
	if err != nil {
		return err
	}
	return nil
}

// SaveCode method
func (l *LevelDB) SaveCode(key string, v interface{}) error {
	log.Debug().Msg("LevelDb.SaveCode()")
	tmp, err := json.Marshal(v)
	if err != nil {
		return err
	}
	l.Open()
	err = l.db.Put([]byte("c:"+key), []byte(string(tmp)), nil)
	if err != nil {
		return err
	}
	return nil

}
