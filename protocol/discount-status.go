/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
)

// DiscountStatus or product sale
type DiscountStatus interface {
	DiscountID() [16]byte    // discount domain
	Status() Discount_Status //
	Time() protocol.Time     // Save time
	RequestID() [16]byte     // user-request domain
}

type Discount_Status uint8

const (
	Discount_Status_Unset Discount_Status = iota
	Discount_Status_Registered
	Discount_Status_Edited
	Discount_Status_Expired
	Discount_Status_Blocked
)
