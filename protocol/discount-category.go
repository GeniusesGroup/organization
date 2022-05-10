/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
)

type DiscountCategory interface {
	DiscountID() [16]byte             // discount domain
	CategoryID() [16]byte             // category domain
	MinAmount() uint64                // Minimum number to buy to use discount
	MinPrice() protocol.AmountOfMoney // Minimum price to buy to use discount
	MinStock() uint64                 // 0 for unlimited until related product exist to sell
	Time() protocol.Time              // Save time
	RequestID() [16]byte              // user-request domain
}
