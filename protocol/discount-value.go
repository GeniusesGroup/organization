/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
)

type DiscountValue interface {
	DiscountID() [16]byte          // discount domain
	Percentage() math.PerMyriad    // max discount percentage for each invoices, products, ... price
	Price() protocol.AmountOfMoney // max discount price can use for each invoices, product, ... price
	Time() protocol.Time           // Save time
	RequestID() [16]byte           // user-request domain
}
