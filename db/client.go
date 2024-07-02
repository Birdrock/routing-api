package db

import (
	"database/sql"

	"gorm.io/gorm"
)

//go:generate counterfeiter -o fakes/fake_client.go . Client
type Client interface {
	// Close() error
	Where(query interface{}, args ...interface{}) Client
	Create(value interface{}) (int64, error)
	Delete(value interface{}, where ...interface{}) (int64, error)
	Save(value interface{}) (int64, error)
	Update(column string, value interface{}) (tx *DB)
	First(out interface{}, where ...interface{}) error
	Find(out interface{}, where ...interface{}) error
	AutoMigrate(values ...interface{}) error
	Begin() Client
	Rollback() error
	Commit() error
	HasTable(value interface{}) bool
	CreateIndex(dst interface{}, name string) error
	DropIndex(dst interface{}, name string) error
	Model(value interface{}) (tx *DB)
	// AddUniqueIndex(indexName string, columns ...string) (Client, error)
	// RemoveIndex(indexName string) (Client, error)
	// Model(value interface{}) Client
	Exec(query string, args ...interface{}) int64
	Rows(tableName string) (*sql.Rows, error)
	DropColumn(dst interface{}, field string) error
}

type gormClient struct {
	db *gorm.DB
}

func NewGormClient(db *gorm.DB) Client {
	return &gormClient{db: db}
}
func (c *gormClient) DropColumn(dst interface{}, name string) error {
	return c.db.Migrator().DropColumn(dst, name)
}

func (c *gormClient) CreateIndex(dst interface{}, name string) error {
	return c.db.Migrator().CreateIndex(dst, name)
}

func (c *gormClient) DropIndex(dst interface{}, name string) error {
	return c.db.Migrator().DropIndex(dst, name)
}

func (c *gormClient) Model(value interface{}) (tx *DB) {
	return c.Begin().Model(value)
}

// func (c *gormClient) Close() error {
// 	return c.db.Close()
// }

// func (c *gormClient) AddUniqueIndex(indexName string, columns ...string) (Client, error) {
// 	var newClient gormClient
// 	newClient.db = c.db.AddUniqueIndex(indexName, columns...)
// 	return &newClient, newClient.db.Error
// }

// func (c *gormClient) RemoveIndex(indexName string) (Client, error) {
// 	var newClient gormClient
// 	newClient.db = c.db.RemoveIndex(indexName)
// 	return &newClient, newClient.db.Error
// }

//	func (c *gormClient) Model(value interface{}) Client {
//		var newClient gormClient
//		newClient.db = c.db.Model(value)
//		return &newClient
//	}
func (c *gormClient) Where(query interface{}, args ...interface{}) (tx *DB) {
	return c.Begin().Where(query, args...)
}

func (c *gormClient) Create(value interface{}) (int64, error) {
	newDb := c.db.Create(value)
	return newDb.RowsAffected, newDb.Error
}

func (c *gormClient) Delete(value interface{}, where ...interface{}) (int64, error) {
	newDb := c.db.Delete(value, where...)
	return newDb.RowsAffected, newDb.Error
}

func (c *gormClient) Save(value interface{}) (int64, error) {
	newDb := c.db.Save(value)
	return newDb.RowsAffected, newDb.Error
}

func (c *gormClient) Update(column string, value interface{}) (tx *DB) {
	ntx := c.Update(column, value)
	return ntx
}

func (c *gormClient) First(out interface{}, where ...interface{}) error {
	return c.db.First(out, where...).Error
}

func (c *gormClient) Find(out interface{}, where ...interface{}) error {
	return c.db.Find(out, where...).Error
}

func (c *gormClient) AutoMigrate(values ...interface{}) error {
	return c.db.AutoMigrate(values...)
}

func (c *gormClient) Begin() Client {
	var newClient gormClient
	newClient.db = c.db.Begin()
	return &newClient
}

func (c *gormClient) Rollback() error {
	return c.db.Rollback().Error
}

func (c *gormClient) Commit() error {
	return c.db.Commit().Error
}

func (c *gormClient) HasTable(dst interface{}) bool {
	return c.db.Migrator().HasTable(dst)
}

func (c *gormClient) Exec(query string, args ...interface{}) int64 {
	dbClient := c.db.Exec(query, args)
	return dbClient.RowsAffected
}

func (c *gormClient) Rows(tablename string) (*sql.Rows, error) {
	tableDb := c.db.Table(tablename)
	return tableDb.Rows()
}
