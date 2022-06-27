/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/math"
	"../libgo/protocol"
)

// DiscountValue indicate the domain record data fields.
type DiscountValue interface {
	DiscountID() [16]byte          // discount domain
	Quantity() math.PerMyriad      // max discount percentage for each invoices, products, ... price
	Price() protocol.AmountOfMoney // max discount price can use for each invoices, product, ... price
	Time() protocol.Time           // save time
	RequestID() [16]byte           // user-request domain
}

type DiscountValue_StorageServices interface {
	Save(dv DiscountValue) (numbers uint64, err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (dv DiscountValue, numbers uint64, err protocol.Error)
}
