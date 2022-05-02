/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type InvoiceVoucher interface {
	InvoiceID() [16]byte                // invoice domain
	VoucherID() [16]byte                // voucher domain
	Discounted() protocol.AmountOfMoney //
	Time() protocol.Time                // Save time
	RequestID() [16]byte                // user-request domain
}

type InvoiceVoucher_StorageServices interface {
	Save(iv InvoiceVoucher) (err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iv InvoiceVoucher, err protocol.Error)
	Last(invoiceID [16]byte) (iv InvoiceVoucher, numbers uint64, err protocol.Error)

	FindByVoucherID(voucherID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
}
