/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type Quiddity interface {
	ID() [16]byte        // Unique content ID
	UserID() [16]byte    // user-status domain. Owner of the quiddity
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

type Quiddity_StorageServices interface {
	Save(q Quiddity) (err protocol.Error)

	Count(id [16]byte) (numbers uint64, err protocol.Error)
	Get(id [16]byte, versionOffset uint64) (q Quiddity, err protocol.Error)
	Last(id [16]byte) (q Quiddity, numbers uint64, err protocol.Error)

	ListUserQuiddities(userID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)
}
