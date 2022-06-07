/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductTimingInterval indicate the domain record data fields.
// e.g. flight duration, doctor visit, carwash, ...
type ProductTimingInterval interface {
	ProductID() [16]byte         // product domain
	Minimum() protocol.Duration  // Minimum time duration for this service
	Expected() protocol.Duration // Expected time duration for this service
	Maximum() protocol.Duration  // Maximum time duration for this service
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type ProductTimingInterval_StorageServices interface {
	Save(pt ProductTimingInterval) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pt ProductTimingInterval, err protocol.Error)
	Last(productID [16]byte) (pt ProductTimingInterval, numbers uint64, err protocol.Error)
}
