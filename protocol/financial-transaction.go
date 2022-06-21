/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// FinancialTransaction indicate the domain record data fields.
// each FinancialTransaction is an immutable record and so use version mechanism to chain FinancialTransactions in AccountID way.
type FinancialTransaction interface {
	AccountID() [16]byte                    // financial-account domain
	AccountSideID() [16]byte                // financial-account domain. or AccountPartyID
	Reference() [16]byte                    // Any ID that user can assign to track for any purpose. It can't be nil.
	ReferenceType() FinancialTransaction_RT //
	Amount() protocol.AmountOfMoney         // This transaction
	Balance() protocol.AmountOfMoney        // Account balance with this transaction
	Time() protocol.Time                    // save time
	RequestID() [16]byte                    // user-request domain
}

type FinancialTransaction_StorageServices interface {
	Save(ft FinancialTransaction) (numbers uint64, err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (ft FinancialTransaction, numbers uint64, err protocol.Error)

	// TODO::: is it worth to uncomment below service?
	// FindByAccountSideID(accountID, accountSideID [16]byte, offset, limit uint64) (versionOffsets []uint64, numbers uint64, err protocol.Error)
}

// FinancialTransaction_RT indicate FinancialTransaction record reference type
type FinancialTransaction_RT uint8

// FinancialTransaction types
const (
	FinancialTransaction_RT_Unset   FinancialTransaction_RT = iota
	FinancialTransaction_RT_Failed                          // Refund for any reason e.g. not same currency account, not reachable bank, ...
	FinancialTransaction_RT_Blocked                         // by justice
	FinancialTransaction_RT_TransactionFee
	FinancialTransaction_RT_Exchange
	FinancialTransaction_RT_Prize
	FinancialTransaction_RT_Donate
	FinancialTransaction_RT_Charity
	FinancialTransaction_RT_Tip
	FinancialTransaction_RT_Cash          // InvoiceID, PosID >> Due to digital era banking, prefer to not support cash transactions
	FinancialTransaction_RT_RevolvingFund // >> Due to digital era banking, prefer to not support this type
	// FinancialTransaction_RT_SameBankTransfer     >> Not the goal and just the way of transaction that can check by RequestID
	// FinancialTransaction_RT_ForeignBankTransfer  >> Not the goal and just the way of transaction that can check by RequestID
	// FinancialTransaction_RT_POSTransfer          >> Not the goal and just the way of transaction that can check by RequestID
	// FinancialTransaction_RT_WebTransfer          >> Not the goal and just the way of transaction that can check by RequestID
	FinancialTransaction_RT_Asset
	FinancialTransaction_RT_Stock
	FinancialTransaction_RT_RentalFee
	FinancialTransaction_RT_Bill
	FinancialTransaction_RT_Invoice // both buy or refund for any reason(ReturnInvoice)
	FinancialTransaction_RT_Tax
	FinancialTransaction_RT_Commission
	FinancialTransaction_RT_Salary // Reward
	FinancialTransaction_RT_Profit
	FinancialTransaction_RT_Credit
	FinancialTransaction_RT_Loan
	FinancialTransaction_RT_BadDebt // Debt default
	FinancialTransaction_RT_Installment
	FinancialTransaction_RT_Compensation
)
