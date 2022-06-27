/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// DiscountInvoice indicate the domain record data fields.
type DiscountInvoice interface {
	DiscountID() [16]byte             // discount domain
	MinPrice() protocol.AmountOfMoney // Minimum invoice price
	MinAmount() uint64                // Minimum invoice product numbers
	Time() protocol.Time              // save time
	RequestID() [16]byte              // user-request domain
}

type DiscountInvoice_StorageServices interface {
	Save(di DiscountInvoice) (numbers uint64, err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (di DiscountInvoice, numbers uint64, err protocol.Error)
}
