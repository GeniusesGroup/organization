/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// DiscountLocation indicate the domain record data fields.
type DiscountLocation interface {
	DiscountID() [16]byte         // discount domain
	BuildingLocationID() [16]byte // building-location domain
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type DiscountLocation_StorageServices interface {
	Save(dl DiscountLocation) (numbers uint64, err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (dl DiscountLocation, numbers uint64, err protocol.Error)

	FindByBuildingLocation(buildingLocationID [16]byte, offset, limit uint64) (discountIDs [][16]byte, numbers uint64, err protocol.Error)
}
