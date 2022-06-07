/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// InvoiceItemReserveTime indicate the domain record data fields.
type InvoiceItemReserveTime interface {
	InvoiceID() [16]byte  // invoice domain
	ProductID() [16]byte  // product domain
	Start() protocol.Time //
	End() protocol.Time   //
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type InvoiceItemReserveTime_StorageServices interface {
	Save(itr InvoiceItemReserveTime) (err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (itr InvoiceItemReserveTime, err protocol.Error)
	Last(invoiceID [16]byte) (itr InvoiceItemReserveTime, numbers uint64, err protocol.Error)
}
