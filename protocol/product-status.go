/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductStatus indicate the domain record data fields.
type ProductStatus interface {
	ProductID() [16]byte    // product domain
	Status() Product_Status //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ProductStatus_StorageServices interface {
	Save(ps ProductStatus) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (ps ProductStatus, err protocol.Error)
	Last(productID [16]byte) (ps ProductStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status Product_Status, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)

	protocol.EventTarget
}

type Product_Status Quiddity_Status

const (
// Product_Status_ = Product_Status(Quiddity_Status_FreeFlag << iota)
)
