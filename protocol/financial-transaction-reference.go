/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// FinancialTransactionReference store the secondary index that complicated than regular indexes.
// Use version to save multi transactions for the same reference.
type FinancialTransactionReference interface {
	ReferenceID() [16]byte         // any domain base on FinancialTransaction_Type
	SenderAccountID() [16]byte     // financial-account || financial-bank-account domain
	SenderAccountOffset() uint64   //
	ReceiverAccountID() [16]byte   // financial-account || financial-bank-account domain. or AccountPartyID
	ReceiverAccountOffset() uint64 //
	Time() protocol.Time           // Save time
}

type FinancialTransactionReference_StorageServices interface {
	Save(ftr FinancialTransactionReference) (err protocol.Error)

	Count(referenceID [16]byte) (numbers uint64, err protocol.Error)
	Get(referenceID [16]byte, versionOffset uint64) (ftr FinancialTransactionReference, err protocol.Error)
	Last(referenceID [16]byte) (ftr FinancialTransactionReference, numbers uint64, err protocol.Error)
}
