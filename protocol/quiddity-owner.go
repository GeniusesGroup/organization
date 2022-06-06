/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// QuiddityOwner indicate the domain record data fields.
type QuiddityOwner interface {
	QuiddityID() [16]byte // Unique content ID
	UserID() [16]byte     // user domain. Owner of the quiddity
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type QuiddityOwner_StorageServices interface {
	Save(qo QuiddityOwner) (err protocol.Error)

	Count(quiddityID [16]byte) (numbers uint64, err protocol.Error)
	Get(quiddityID [16]byte, versionOffset uint64) (qo QuiddityOwner, err protocol.Error)
	Last(quiddityID [16]byte) (qo QuiddityOwner, numbers uint64, err protocol.Error)

	FindByOwnerID(userID [16]byte, offset, limit uint64) (quiddityIDs [][16]byte, numbers uint64, err protocol.Error)
}

type QuiddityOwner_Service_FindByUserID_Request interface {
	UserID() [32]byte
	Offset() uint64
	Limit() uint64
}
type QuiddityOwner_Service_FindByUserID_Response interface {
	QuiddityIDs() [][32]byte
	Numbers() uint64
}
