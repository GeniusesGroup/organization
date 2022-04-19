/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

/*
	Storage Data Structure
*/

type FinancialTransaction interface {
	AccountID() [16]byte            // bank-account ||  domain
	AccountSideID() [16]byte        // or AccountPartyID
	Reference() [16]byte            // Any ID that user can assign to track for any purpose
	Amount() int64                  // This transaction
	Balance() int64                 // Account balance with this transaction
	Type() FinancialTransactionType //
	Time() protocol.Time            // Set or update time
	RequestID() [16]byte            // user-request domain
}

type FinancialTransactionStorageServices interface {
	Save(ft FinancialTransaction) (err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (ft FinancialTransaction, err protocol.Error)
	Last(accountID [16]byte) (ft FinancialTransaction, err protocol.Error)
}

// FinancialTransactionType indicate FinancialTransaction record type
type FinancialTransactionType uint8

// FinancialTransaction types
const (
	FinancialTransactionType_Unset  FinancialTransactionType = iota
	FinancialTransactionType_Failed                          // Refund for any reason e.g. not same currency account, not reachable bank, ...
	FinancialTransactionType_Blocked
	FinancialTransactionType_Exchange
	FinancialTransactionType_Prize
	FinancialTransactionType_Donate
	FinancialTransactionType_Charity
	FinancialTransactionType_Tip
	FinancialTransactionType_Cash          // >> Due to digital era banking, prefer to not support cash transactions
	FinancialTransactionType_RevolvingFund // >> Due to digital era banking, prefer to not support this type
	// FinancialTransactionType_SameBankTransfer     >> Not the goal and just the way of transaction that can check by RequestID
	// FinancialTransactionType_ForeignBankTransfer  >> Not the goal and just the way of transaction that can check by RequestID
	// FinancialTransactionType_POSTransfer          >> Not the goal and just the way of transaction that can check by RequestID
	// FinancialTransactionType_WebTransfer          >> Not the goal and just the way of transaction that can check by RequestID
	FinancialTransactionType_Asset
	FinancialTransactionType_Stock
	FinancialTransactionType_RentalFee
	FinancialTransactionType_Bill
	FinancialTransactionType_Invoice
	FinancialTransactionType_ReturnInvoice
	FinancialTransactionType_Tax
	FinancialTransactionType_Commission
	FinancialTransactionType_Salary // Reward
	FinancialTransactionType_Profit
	FinancialTransactionType_Loan
	FinancialTransactionType_BadDebt // Debt default
	FinancialTransactionType_Installment
	FinancialTransactionType_Compensation
)
