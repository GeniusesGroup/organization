/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// Category or Topic
type Category interface {
	QuiddityID() [16]byte // CategoryID
	ParentID() [16]byte   // CategoryID, Exist if it isn't top category.
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}

// TODO::: cache strategy?? 1 day or what?
type Category_StorageServices interface {
	Save(c Category) (err protocol.Error)

	GetIDs(offset, limit uint64) (quiddityIDs [][16]byte, numbers uint64, err protocol.Error)

	Delete(quiddityID [16]byte) (err protocol.Error)
}
