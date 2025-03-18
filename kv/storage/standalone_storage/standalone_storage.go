package standalone_storage

import (
	"io/ioutil"

	"github.com/Connor1996/badger"
	"github.com/villanel/tinykv-scheduler/kv/config"
	"github.com/villanel/tinykv-scheduler/kv/storage"
	"github.com/villanel/tinykv-scheduler/kv/util/engine_util"
	"github.com/villanel/tinykv-scheduler/proto/pkg/kvrpcpb"
)

// StandAloneStorage is an implementation of `Storage` for a single-node TinyKV instance. It does not
// communicate with other nodes and all data is stored locally.
type StandAloneStorage struct {
	db *badger.DB
}

func NewStandAloneStorage(conf *config.Config) *StandAloneStorage {

	return &StandAloneStorage{
		nil,
	}
}

func (s *StandAloneStorage) Start() error {
	dir, _ := ioutil.TempDir("", "engine_util")
	opts := badger.DefaultOptions
	opts.Dir = dir
	opts.ValueDir = dir
	db, _ := badger.Open(opts)
	s.db = db
	return nil
}

func (s *StandAloneStorage) Stop() error {
	err := s.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *StandAloneStorage) Reader(ctx *kvrpcpb.Context) (storage.StorageReader, error) {
	return &StandaloneReader{
		s.db.NewTransaction(false),
	}, nil

}

func (s *StandAloneStorage) Write(ctx *kvrpcpb.Context, batch []storage.Modify) error {
	e := new(engine_util.WriteBatch)
	for _, modify := range batch {
		if modify.Value() != nil {
			e.SetCF(modify.Cf(), modify.Key(), modify.Value())
		} else {
			e.DeleteCF(modify.Cf(), modify.Key())
		}
	}
	return e.WriteToDB(s.db)
}
