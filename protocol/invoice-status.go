/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// InvoiceStatus indicate the domain record data fields.
// Also usually known as shopping-cart, ... in GUI part of applications
type InvoiceStatus interface {
	InvoiceID() [16]byte    // invoice-status domain
	Status() Invoice_Status //
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

type InvoiceStatus_StorageServices interface {
	Save(is InvoiceStatus) (err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (is InvoiceStatus, err protocol.Error)
	Last(invoiceID [16]byte) (is InvoiceStatus, numbers uint64, err protocol.Error)

	// GetIDsByDateTime(time protocol.Time, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)

	FindByStatus(offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
}

type Invoice_Status uint8

const (
	Invoice_Status_Unset Invoice_Status = iota // or Draft
	Invoice_Status_Blocked
	// deleted : This state indicates that the invoice has been deleted.
	// You can only delete an invoice in the draft state and not in any other state.
	Invoice_Status_Hidden // use for Deleted
	// cancelled : Indicates that you have cancelled the invoice. Though you can view the cancelled invoice,
	// the customer cannot view the invoice or make payments against the invoice.
	// You can only cancel an unpaid issued invoice.
	Invoice_Status_Voided // invoice is no longer valid, and not payable by client or refund

	// expired : Indicates that an invoice has expired. Once expired,
	// the customer cannot make any payments against the invoice,
	// and you cannot cancel or delete the invoice. You can view an expired invoice.
	// Invoice_Status_Expired

	// partially paid : Indicates that the customer has made a partial payment against an invoice.
	// After a payment has been made, you can neither delete nor cancel the invoice.
	// You can only view the invoice and add internal notes to it.
	Invoice_Status_PartiallyPaid

	// paid : Indicates that the customer has fully paid the invoice amount.
	// Paid is the final state of an invoice. A paid invoice can only be viewed.
	// You can neither delete nor cancel it.
	Invoice_Status_Paid
	Invoice_Status_Committed
	Invoice_Status_PendingReview
	Invoice_Status_Confirmed // Approved
)
