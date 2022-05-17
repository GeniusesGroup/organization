/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// GroupRole indicate the domain record data fields.
type GroupRole interface {
	ID() [16]byte                          // quiddity domain
	GroupID() [16]byte                     // group-status domain
	AccessControl() protocol.AccessControl //
	Time() protocol.Time                   // Save time
	RequestID() [16]byte                   // user-request domain
}

type GroupRole_StorageServices interface {
	Save(gr GroupRole) protocol.Error

	Count(id [16]byte) (numbers uint64, err protocol.Error)
	Get(id [16]byte, versionOffset uint64) (gr GroupRole, err protocol.Error)
	Last(id [16]byte) (gr GroupRole, err protocol.Error)

	FindByGroupID(groupID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)
}
