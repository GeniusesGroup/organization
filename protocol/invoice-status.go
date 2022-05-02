/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// InvoiceStatus indicate the invoice-status domain record data fields.
type InvoiceStatus interface {
	InvoiceID() [16]byte          // invoice domain
	Status() InvoiceStatus_Status //
	Time() protocol.Time          // Save time
	RequestID() [16]byte          // user-request domain
}

type InvoiceStatus_StorageServices interface {
	Save(is InvoiceStatus) (err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (is InvoiceStatus, err protocol.Error)
	Last(invoiceID [16]byte) (is InvoiceStatus, numbers uint64, err protocol.Error)
}

type InvoiceStatus_Status uint8

const (
	InvoiceStatus_Status_Unset InvoiceStatus_Status = iota // or Draft
	InvoiceStatus_Status_Blocked
	// deleted : This state indicates that the invoice has been deleted.
	// You can only delete an invoice in the draft state and not in any other state.
	InvoiceStatus_Status_Hidden // use for Deleted
	// cancelled : Indicates that you have cancelled the invoice. Though you can view the cancelled invoice,
	// the customer cannot view the invoice or make payments against the invoice.
	// You can only cancel an unpaid issued invoice.
	InvoiceStatus_Status_Voided // invoice is no longer valid, and not payable by client or refund

	// expired : Indicates that an invoice has expired. Once expired,
	// the customer cannot make any payments against the invoice,
	// and you cannot cancel or delete the invoice. You can view an expired invoice.
	// InvoiceStatus_Status_Expired

	// partially paid : Indicates that the customer has made a partial payment against an invoice.
	// After a payment has been made, you can neither delete nor cancel the invoice.
	// You can only view the invoice and add internal notes to it.
	InvoiceStatus_Status_PartiallyPaid

	// paid : Indicates that the customer has fully paid the invoice amount.
	// Paid is the final state of an invoice. A paid invoice can only be viewed.
	// You can neither delete nor cancel it.
	InvoiceStatus_Status_Paid
	InvoiceStatus_Status_Committed
	InvoiceStatus_Status_PendingReview
	InvoiceStatus_Status_Confirmed // Approved
)
