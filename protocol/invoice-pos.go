/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// InvoicePOS indicate the domain record data fields.
type InvoicePOS interface {
	InvoiceID() [16]byte // invoice domain
	PosID() [16]byte     // pos domain
	StaffID() [16]byte   // staff domain. or OrdererID that is creator of the invoice e.g. restaurant garson, drug prescription, sales agent,
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type InvoicePOS_StorageServices interface {
	Save(ip InvoicePOS) protocol.Error

	Get(invoiceID [16]byte) (ip InvoicePOS, err protocol.Error)

	FindByPosID(posID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
}
