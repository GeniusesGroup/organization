/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

/*
	Storage Data Structure
*/

type FinancialAccount interface {
	AccountID() [16]byte            //
	UserID() [16]byte               // user-status domain
	SocietyID() [16]byte            // society domain
	Status() FinancialAccountStatus //
	Time() protocol.Time            // Set or update time
	RequestID() [16]byte            // user-request domain
}

type FinancialAccountServices interface {
	Save(fa FinancialAccount) (err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fa FinancialAccount, err protocol.Error)
	Last(accountID [16]byte) (fa FinancialAccount, err protocol.Error)

	ListUserAccounts(userID [16]byte, offset, limit uint64) (ids [][16]byte, err protocol.Error)
}

// FinancialAccountStatus indicate FinancialAccount record status
type FinancialAccountStatus uint8

// FinancialAccount status
const (
	FinancialAccountStatus_Unset = iota
	FinancialAccountStatus_Registered
	FinancialAccountStatus_Closed
	FinancialAccountStatus_Opened
	FinancialAccountStatus_Blocked
	FinancialAccountStatus_Unblocked
)
