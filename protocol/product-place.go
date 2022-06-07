/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductPlace indicate the domain record data fields.
// ProductPlace indicate service provides in specific place e.g. stations(bus, train, airport, ...), ...
type ProductPlace interface {
	ProductID() [16]byte          // product domain.
	BuildingLocationID() [16]byte // building-location domain. use also for Coordinate domain
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type ProductPlace_StorageServices interface {
	Save(pp ProductPlace) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pp ProductPlace, err protocol.Error)
	Last(productID [16]byte) (pp ProductPlace, numbers uint64, err protocol.Error)

	FindByBuildingLocation(buildingLocationID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}
