package sqlite

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"{{ .App.ModuleName }}/internal/store"
	v1 "{{ .App.ModuleName }}/pkg/api/apiserver/v1"
	"{{ .App.ModuleName }}/pkg/xerr"
)

var (
	_factory store.Factory
	once     sync.Once
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}

func (ds *datastore) Close() error {
	instance, err := ds.db.DB()
	if err != nil {
		return xerr.Wrap(err, "get gorm db instance failed")
	}
	return instance.Close()
}

// GetSQLiteFactoryOr returns a SQLite-backed store, creating it on first call.
// dsn uses SQLite DSN syntax, e.g. "file::memory:?cache=shared" or a file path.
func GetSQLiteFactoryOr(dsn string) (store.Factory, error) {
	var err error
	once.Do(func() {
		var instance *gorm.DB
		instance, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			return
		}
		if migrateErr := instance.AutoMigrate(&v1.User{}); migrateErr != nil {
			err = migrateErr
			return
		}
		_factory = &datastore{instance}
	})
	if err != nil {
		return nil, err
	}
	return _factory, nil
}
