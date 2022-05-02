/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// InvoiceSummary is summary data that calculated to reduce fetch and calculate on each client device.
type InvoiceSummary interface {
	InvoiceID() [16]byte                     // invoice domain
	Quantity() uint64                        // Total product count
	SuggestedPrice() protocol.AmountOfMoney  // Total should pay before any discount or auction
	DiscountedPrice() protocol.AmountOfMoney // user profit from this purchase by vouchers+auctions
	VoucherPrice() protocol.AmountOfMoney    // the sum of all vouchers apply for this invoice
	PayablePrice() protocol.AmountOfMoney    // Total price user pay to organization
	VAT() protocol.AmountOfMoney             // Total VAT user pay to its country government
	Time() protocol.Time                     // Save time
}

type InvoiceSummary_StorageServices interface {
	Save(is InvoiceSummary) (err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (is InvoiceSummary, err protocol.Error)
	Last(invoiceID [16]byte) (is InvoiceSummary, numbers uint64, err protocol.Error)
}
