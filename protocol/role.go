/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// Role indicate the domain record data fields.
type Role interface {
	ID() [16]byte                          // quiddity domain. use to get and show locale title.
	AccessControl() protocol.AccessControl //
	Time() protocol.Time                   // Save time
	RequestID() [16]byte                   // user-request domain
}

type Role_StorageServices interface {
	Save(r Role) protocol.Error

	Count(id [16]byte) (numbers uint64, err protocol.Error)
	Get(id [16]byte, versionOffset uint64) (r Role, err protocol.Error)
	Last(id [16]byte) (r Role, err protocol.Error)

	GetIDs(offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)
}
