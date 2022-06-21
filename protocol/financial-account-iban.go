/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/iso"
	"../libgo/protocol"
)

// FinancialAccountIBAN indicate the domain record data fields.
type FinancialAccountIBAN interface {
	AccountID() [16]byte // financial-account domain
	IBAN() iso.IBAN      //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type FinancialAccountIBAN_StorageServices interface {
	Save(fai FinancialAccountIBAN) (numbers uint64, err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fai FinancialAccountIBAN, numbers uint64, err protocol.Error)

	FindByIBAN(iban iso.IBAN) (accountID [16]byte, err protocol.Error)
}

type FinancialAccountIBAN_Service_Register_Request interface {
	AccountID() [16]byte
	IBAN() iso.IBAN
}

type FinancialAccountIBAN_Service_Get_Request interface {
	AccountID() [16]byte
	VersionOffset() uint64
}
type FinancialAccountIBAN_Service_Get_Response interface {
	FinancialAccountIBAN
	Numbers() uint64
}
