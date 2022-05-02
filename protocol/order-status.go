/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// OrderStatus indicate the order-status domain record data fields.
type OrderStatus interface {
	OrderID() [16]byte          // order domain
	Status() OrderStatus_Status //
	Time() protocol.Time        // Save time
	RequestID() [16]byte        // user-request domain
}

type OrderStatus_Status uint8

const (
	OrderStatus_Status_Unset OrderStatus_Status = iota
	OrderStatus_Status_New
	OrderStatus_Status_Pending // Wait to approve
	OrderStatus_Status_PartiallyDeal
	OrderStatus_Status_FullyDeal
	OrderStatus_Status_Expired
	OrderStatus_Status_Cancelled // Voided
	OrderStatus_Status_Blocked
)
