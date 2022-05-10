/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type DiscountInvoice interface {
	DiscountID() [16]byte             // discount domain
	MinPrice() protocol.AmountOfMoney // Minimum invoice price
	MinAmount() uint64                // Minimum invoice product numbers
	Time() protocol.Time              // Save time
	RequestID() [16]byte              // user-request domain
}
