package migration

import (
	"code.cloudfoundry.org/routing-api/db"
	v5 "code.cloudfoundry.org/routing-api/migration/v5"
)

type V5SniHostnameMigration struct{}

var _ Migration = new(V5SniHostnameMigration)

func NewV5SniHostnameMigration() *V5SniHostnameMigration {
	return &V5SniHostnameMigration{}
}

func (v *V5SniHostnameMigration) Version() int {
	return 5
}

func (v *V5SniHostnameMigration) Run(sqlDB *db.SqlDB) error {
	_ = sqlDB.Client.DropIndex(&v5.TcpRouteMapping{}, "idx_tcp_route")
	return sqlDB.Client.AutoMigrate(&v5.TcpRouteMapping{})
}
