/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductTag indicate the domain record data fields.
type ProductTag interface {
	Tag() string         //
	ProductID() [16]byte // product domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductTag_StorageServices interface {
	Save(pt ProductTag) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pt ProductTag, err protocol.Error)
	Last(productID [16]byte) (pt ProductTag, numbers uint64, err protocol.Error)

	FindByProduct(productID [16]byte, offset, limit uint64) (tags []string, numbers uint64, err protocol.Error)
}
