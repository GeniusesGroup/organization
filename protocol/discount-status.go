/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// DiscountStatus indicate the domain record data fields.
type DiscountStatus interface {
	DiscountID() [16]byte    // discount domain
	Status() Discount_Status //
	Time() protocol.Time     // save time
	RequestID() [16]byte     // user-request domain
}

type DiscountStatus_StorageServices interface {
	Save(ds DiscountStatus) (numbers uint64, err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (ds DiscountStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status Discount_Status, offset, limit uint64) (discountIDs [][16]byte, numbers uint64, err protocol.Error)

	protocol.EventTarget
}

type Discount_Status Quiddity_Status

const (
	Discount_Status_Expired = Discount_Status(Quiddity_Status_FreeFlag << iota)
	Discount_Status_End
)
