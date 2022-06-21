/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// FinancialAccountStatus indicate the domain record data fields.
type FinancialAccountStatus interface {
	AccountID() [16]byte             // financial-account domain
	Status() FinancialAccount_Status //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}

type FinancialAccountStatus_StorageServices interface {
	Save(fas FinancialAccountStatus) (numbers uint64, err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fas FinancialAccountStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status FinancialAccount_Status, offset, limit uint64) (accountIDs [][16]byte, numbers uint64, err protocol.Error)
}

type FinancialAccount_Status Quiddity_Status

const (
	FinancialAccount_Status_Closed = FinancialAccount_Status(Quiddity_Status_FreeFlag << iota)
)
