/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// InvoiceItemProvider indicate the domain record data fields.
type InvoiceItemProvider interface {
	InvoiceID() [16]byte // invoice domain
	ProductID() [16]byte // product domain
	StaffID() [16]byte   // staff domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type InvoiceItemProvider_StorageServices interface {
	Save(iip InvoiceItemProvider) (err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iip InvoiceItemProvider, err protocol.Error)
	Last(invoiceID [16]byte) (iip InvoiceItemProvider, numbers uint64, err protocol.Error)

	FindByStaffID(staffID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
}
