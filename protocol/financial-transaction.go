/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

/*
	Storage Data Structure
*/

type FinancialTransaction interface {
	AccountID() [16]byte             // bank-account ||  domain
	AccountSideID() [16]byte         // or AccountPartyID
	Reference() [16]byte             // Any ID that user can assign to track for any purpose
	Amount() protocol.AmountOfMoney  // This transaction
	Balance() protocol.AmountOfMoney // Account balance with this transaction
	Type() FinancialTransaction_Type //
	Time() protocol.Time             // Save time
	RequestID() [16]byte             // user-request domain
}

type FinancialTransaction_StorageServices interface {
	Save(ft FinancialTransaction) (err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (ft FinancialTransaction, err protocol.Error)
	Last(accountID [16]byte) (ft FinancialTransaction, err protocol.Error)
}

// FinancialTransaction_Type indicate FinancialTransaction record type
type FinancialTransaction_Type uint8

// FinancialTransaction types
const (
	FinancialTransaction_Type_Unset  FinancialTransaction_Type = iota
	FinancialTransaction_Type_Failed                           // Refund for any reason e.g. not same currency account, not reachable bank, ...
	FinancialTransaction_Type_Blocked
	FinancialTransaction_Type_Exchange
	FinancialTransaction_Type_Prize
	FinancialTransaction_Type_Donate
	FinancialTransaction_Type_Charity
	FinancialTransaction_Type_Tip
	FinancialTransaction_Type_Cash          // >> Due to digital era banking, prefer to not support cash transactions
	FinancialTransaction_Type_RevolvingFund // >> Due to digital era banking, prefer to not support this type
	// FinancialTransaction_Type_SameBankTransfer     >> Not the goal and just the way of transaction that can check by RequestID
	// FinancialTransaction_Type_ForeignBankTransfer  >> Not the goal and just the way of transaction that can check by RequestID
	// FinancialTransaction_Type_POSTransfer          >> Not the goal and just the way of transaction that can check by RequestID
	// FinancialTransaction_Type_WebTransfer          >> Not the goal and just the way of transaction that can check by RequestID
	FinancialTransaction_Type_Asset
	FinancialTransaction_Type_Stock
	FinancialTransaction_Type_RentalFee
	FinancialTransaction_Type_Bill
	FinancialTransaction_Type_Invoice
	FinancialTransaction_Type_ReturnInvoice
	FinancialTransaction_Type_Tax
	FinancialTransaction_Type_Commission
	FinancialTransaction_Type_Salary // Reward
	FinancialTransaction_Type_Profit
	FinancialTransaction_Type_Loan
	FinancialTransaction_Type_BadDebt // Debt default
	FinancialTransaction_Type_Installment
	FinancialTransaction_Type_Compensation
)
