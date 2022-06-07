/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductBestBefore indicate the domain record data fields.
// https://en.wikipedia.org/wiki/Expiration_date
type ProductBestBefore interface {
	ProductID() [16]byte         // product domain
	Duration() protocol.Duration // from production time
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type ProductBestBefore_StorageServices interface {
	Save(pb ProductBestBefore) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pb ProductBestBefore, err protocol.Error)
	Last(productID [16]byte) (pb ProductBestBefore, numbers uint64, err protocol.Error)
}
