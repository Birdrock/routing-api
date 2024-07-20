package migration

import (
	"code.cloudfoundry.org/routing-api/db"
	v4 "code.cloudfoundry.org/routing-api/migration/v4"
)

type V4AddRgUniqIdxTCPRoute struct{}

var _ Migration = new(V4AddRgUniqIdxTCPRoute)

func NewV4AddRgUniqIdxTCPRouteMigration() *V4AddRgUniqIdxTCPRoute {
	return &V4AddRgUniqIdxTCPRoute{}
}

func (v *V4AddRgUniqIdxTCPRoute) Version() int {
	return 4
}

func (v *V4AddRgUniqIdxTCPRoute) Run(sqlDB *db.SqlDB) error {
	return sqlDB.Client.AutoMigrate(&v4.TcpRouteMapping{})
}
