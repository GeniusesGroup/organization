/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// DiscountUsage indicate the domain record data fields.
type DiscountUsage interface {
	DiscountID() [16]byte   // product domain
	MaximumInvoice() uint64 //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type DiscountUsage_StorageServices interface {
	Save(du DiscountUsage) (numbers uint64, err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (du DiscountUsage, numbers uint64, err protocol.Error)
}
