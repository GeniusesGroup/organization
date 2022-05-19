/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// RoleLocation indicate the domain record data fields.
// Let staffs register enter and exit time by their mobiles in specific location.
type RoleLocation interface {
	RoleID() [16]byte     // role domain
	LocationID() [16]byte // location domain
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}

type RoleLocation_StorageServices interface {
	Save(rl RoleLocation) protocol.Error

	Count(roleID [16]byte) (numbers uint64, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (rl RoleLocation, err protocol.Error)
	Last(roleID [16]byte) (rl RoleLocation, err protocol.Error)
}
