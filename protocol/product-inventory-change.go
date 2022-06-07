/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductInventoryChange indicate the domain record data fields.
type ProductInventoryChange interface {
	ProductInventoryChangeID() [16]byte //
	ProductID() [16]byte                // product domain
	From() [16]byte                     // building-location domain
	To() [16]byte                       // building-location domain. Can be zero for voiding products.
	Amount() int64                      // this transaction
	Priority() uint64                   //
	Time() protocol.Time                // save time
	RequestID() [16]byte                // user-request domain
}

type ProductInventoryChange_StorageServices interface {
	Save(pi ProductInventoryChange) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pi ProductInventoryChange, err protocol.Error)
	Last(productID [16]byte) (pi ProductInventoryChange, numbers uint64, err protocol.Error)

	FindByProduct(productID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}
