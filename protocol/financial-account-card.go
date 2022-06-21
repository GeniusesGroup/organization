/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/iso"
	"../libgo/protocol"
)

// FinancialAccountCard indicate the domain record data fields.
type FinancialAccountCard interface {
	AccountID() [16]byte       // financial-account domain
	CardNumber() iso.Card      //
	ExpireDate() protocol.Time //
	CVC() uint16               //
	Time() protocol.Time       // save time
	RequestID() [16]byte       // user-request domain
}

type FinancialAccountCard_StorageServices interface {
	Save(fac FinancialAccountCard) (numbers uint64, err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fac FinancialAccountCard, numbers uint64, err protocol.Error)

	FindByCardNumber(cardNumber iso.Card) (accountID [16]byte, err protocol.Error)
}

type FinancialAccountCard_Service_Register_Request interface {
	AccountID() [16]byte
	CardNumber() iso.Card
	ExpireDate() protocol.Time
	CVC() uint16
}

type FinancialAccountCard_Service_Get_Request interface {
	AccountID() [16]byte
	VersionOffset() uint64
}
type FinancialAccountCard_Service_Get_Response interface {
	FinancialAccountCard
	Numbers() uint64
}
