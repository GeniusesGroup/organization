/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type FinancialAccount interface {
	AccountID() [16]byte             //
	UserID() [16]byte                // user-status domain
	SocietyID() [16]byte             // society domain
	Status() FinancialAccount_Status //
	Time() protocol.Time             // Save time
	RequestID() [16]byte             // user-request domain
}

type FinancialAccount_StorageServices interface {
	Save(fa FinancialAccount) (err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fa FinancialAccount, err protocol.Error)
	Last(accountID [16]byte) (fa FinancialAccount, err protocol.Error)

	FindByUserID(userID [16]byte, offset, limit uint64) (accountIDs [][16]byte, numbers uint64, err protocol.Error)

	ListUserSocieties(userID [16]byte, offset, limit uint64) (societyIDs [][16]byte, numbers uint64, err protocol.Error)
}

// FinancialAccount_Status indicate FinancialAccount record status
type FinancialAccount_Status uint8

// FinancialAccount status
const (
	FinancialAccount_Status_Unset = iota
	FinancialAccount_Status_Registered
	FinancialAccount_Status_Closed
	FinancialAccount_Status_Opened
	FinancialAccount_Status_Blocked
	FinancialAccount_Status_Unblocked
)
