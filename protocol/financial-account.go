/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// FinancialAccount indicate the domain record data fields.
type FinancialAccount interface {
	AccountID() [16]byte                // quiddity domain?
	Currency() [16]byte                 // financial-currency
	UserID() [16]byte                   // user domain
	MoneySettlementReference() [16]byte // user domain. can be any org or society user.
	Time() protocol.Time                // save time
	RequestID() [16]byte                // user-request domain
}

type FinancialAccount_StorageServices interface {
	Save(fa FinancialAccount) (numbers uint64, err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fa FinancialAccount, numbers uint64, err protocol.Error)

	FindByUserID(userID [16]byte, offset, limit uint64) (accountIDs [][16]byte, numbers uint64, err protocol.Error)

	ListUserCurrencies(userID [16]byte, offset, limit uint64) (currency [][16]byte, numbers uint64, err protocol.Error)
	ListUserSettlementReferences(userID [16]byte, offset, limit uint64) (moneySettlementReferences [][16]byte, numbers uint64, err protocol.Error)
}
