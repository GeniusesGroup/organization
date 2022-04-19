/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/iso"
	"../libgo/protocol"
)

/*
	Storage Data Structure
*/

type BankAccount interface {
	AccountID() [16]byte
	UserID() [16]byte     // user-status domain
	IBAN() iso.IBAN       // International Bank Account Number
	CardNumber() iso.Card // fixed size card number without any dash or space e.g. 1234123412341234
	Currency() uint64
	Status() BankAccountStatus
	Time() protocol.Time // Set or update time
	RequestID() [16]byte // user-request domain
}

type BankAccountServices interface {
	Save(ba BankAccount) (err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (ba BankAccount, err protocol.Error)
	Last(accountID [16]byte) (ba BankAccount, err protocol.Error)

	FindAccountByIBAN(iban string) (accountID [16]byte, err protocol.Error)
	FindAccountByCardNumber(cardNumber string) (accountID [16]byte, err protocol.Error)

	ListUserAccounts(userID [16]byte, offset, limit uint64) (ids [][16]byte, err protocol.Error)
}

// BankAccountStatus indicate BankAccount record status
type BankAccountStatus uint8

// BankAccount status
const (
	BankAccountStatus_Unset BankAccountStatus = 0b00000000
	BankAccountStatus_Block BankAccountStatus = 0b00000001
	BankAccountStatus_Close BankAccountStatus = 0b00000010
)
