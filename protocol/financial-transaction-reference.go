/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// FinancialTransactionReference indicate the domain record data fields.
// FinancialTransactionReference store the secondary index that complicated than regular indexes.
// Use version to save multi transactions for the same reference.
type FinancialTransactionReference interface {
	ReferenceID() [16]byte         // any domain base on FinancialTransaction_Type
	SenderAccountID() [16]byte     // financial-account domain
	SenderAccountOffset() uint64   //
	ReceiverAccountID() [16]byte   // financial-account domain. or AccountPartyID
	ReceiverAccountOffset() uint64 //
	Time() protocol.Time           // save time
}

type FinancialTransactionReference_StorageServices interface {
	Save(ftr FinancialTransactionReference) (numbers uint64, err protocol.Error)

	Count(referenceID [16]byte) (numbers uint64, err protocol.Error)
	Get(referenceID [16]byte, versionOffset uint64) (ftr FinancialTransactionReference, numbers uint64, err protocol.Error)
}
