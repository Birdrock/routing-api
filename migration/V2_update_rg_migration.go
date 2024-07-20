package migration

import (
	"code.cloudfoundry.org/routing-api/db"
	v2 "code.cloudfoundry.org/routing-api/migration/v2"
)

type V2UpdateRgMigration struct{}

var _ Migration = new(V2UpdateRgMigration)

func NewV2UpdateRgMigration() *V2UpdateRgMigration {
	return &V2UpdateRgMigration{}
}

func (v *V2UpdateRgMigration) Version() int {
	return 2
}

func (v *V2UpdateRgMigration) Run(sqlDB *db.SqlDB) error {
	return sqlDB.Client.AutoMigrate(&v2.RouterGroup{})
}
