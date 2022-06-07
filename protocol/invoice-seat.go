/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// InvoiceSeat indicate the domain record data fields.
type InvoiceSeat interface {
	InvoiceID() [16]byte        // invoice domain
	ProductID() [16]byte        // product domain
	SeatID() uint64             // seat number
	Status() InvoiceSeat_Status //
	Time() protocol.Time        // save time
	RequestID() [16]byte        // user-request domain
}

type InvoiceSeat_StorageServices interface {
	Save(is InvoiceSeat) (err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (is InvoiceSeat, err protocol.Error)
	Last(invoiceID [16]byte) (is InvoiceSeat, numbers uint64, err protocol.Error)
}

type InvoiceSeat_Status Quiddity_Status

const (
	InvoiceSeat_Status_Voided = InvoiceSeat_Status(Quiddity_Status_FreeFlag << iota)
)

// service: reserve seat
