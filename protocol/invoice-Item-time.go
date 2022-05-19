/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type InvoiceItemTime interface {
	InvoiceID() [16]byte  // invoice-status domain
	ProductID() [16]byte  // product-status domain
	Start() protocol.Time //
	End() protocol.Time   //
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}

type InvoiceItemTime_StorageServices interface {
	Save(iit InvoiceItemTime) (err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iit InvoiceItemTime, err protocol.Error)

	// FindByProductID(productID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
}
