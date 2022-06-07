/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductInventoryChangeStatus indicate the domain record data fields.
type ProductInventoryChangeStatus interface {
	ProductInventoryChangeID() [16]byte    // product-inventory-change domain
	Status() ProductInventoryChange_Status // overall
	Time() protocol.Time                   // save time
	RequestID() [16]byte                   // user-request domain
}

type ProductInventoryChangeStatus_StorageServices interface {
	Save(ps ProductInventoryChangeStatus) protocol.Error

	Count(productInventoryChangeID [16]byte) (numbers uint64, err protocol.Error)
	Get(productInventoryChangeID [16]byte, versionOffset uint64) (ps ProductInventoryChangeStatus, err protocol.Error)
	Last(productInventoryChangeID [16]byte) (ps ProductInventoryChangeStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status ProductInventoryChange_Status, offset, limit uint64) (productInventoryChangeIDs [][16]byte, numbers uint64, err protocol.Error)

	protocol.EventTarget
}

type ProductInventoryChange_Status Quiddity_Status

// DC abbreviation for distribution-center or Warehouse
const (
	ProductInventoryChange_Status_OutcomeDCRequested = ProductInventoryChange_Status(Quiddity_Status_FreeFlag << iota)
	ProductInventoryChange_Status_IncomeDCRequested
	ProductInventoryChange_Status_OutcomeDCApproved
	ProductInventoryChange_Status_IncomeDCApproved
	ProductInventoryChange_Status_Void
)
