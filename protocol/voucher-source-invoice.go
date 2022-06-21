/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// VoucherSourceInvoice indicate the domain record data fields.
type VoucherSourceInvoice interface {
	VoucherID() [16]byte                    // voucher domain
	InvoiceID() [16]byte                    // invoice domain
	InvoiceType() VoucherSourceInvoice_Type //
	Time() protocol.Time                    // save time
	RequestID() [16]byte                    // user-request domain
}

type VoucherSourceInvoice_StorageServices interface {
	Save(vsi VoucherSourceInvoice) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vsi VoucherSourceInvoice, numbers uint64, err protocol.Error)

	FindByInvoice(invoiceID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, numbers uint64, err protocol.Error)
}

type VoucherSourceInvoice_Type uint8

const (
	VoucherSourceInvoice_Type_Unset VoucherSourceInvoice_Type = iota
	VoucherSourceInvoice_Type_PastCompleted
	VoucherSourceInvoice_Type_HurryToBuying
	VoucherSourceInvoice_Type_BackToBuy
)
