/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// DiscountProduct or product sale indicate the domain record data fields.
type DiscountProduct interface {
	DiscountID() [16]byte             // discount domain
	ProductID() [16]byte              // product domain
	MinAmount() uint64                // Minimum number to buy to use discount
	MinPrice() protocol.AmountOfMoney // Minimum price to buy to use discount
	MinStock() uint64                 // 0 for unlimited until related product exist to sell
	Time() protocol.Time              // save time
	RequestID() [16]byte              // user-request domain
}

type DiscountProduct_StorageServices interface {
	Save(dp DiscountProduct) (numbers uint64, err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (dp DiscountProduct, numbers uint64, err protocol.Error)

	FindByProduct(productID [16]byte, offset, limit uint64) (discountIDs [][16]byte, numbers uint64, err protocol.Error)
}
