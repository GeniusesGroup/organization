/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type InvoiceItemProvider interface {
	InvoiceID() [16]byte // invoice-status domain
	ProductID() [16]byte // product-status domain
	StaffID() [16]byte   // staff-status domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

type InvoiceItemProvider_StorageServices interface {
	Save(iip InvoiceItemProvider) (err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iip InvoiceItemProvider, err protocol.Error)

	// FindByProductID(productID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
	// FindByUserID(userID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
}
