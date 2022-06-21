/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// Voucher indicate the domain record data fields.
// Voucher or gift voucher is a bond of the redeemable transaction type which is worth a certain monetary value and
// which may be spent only for specific reasons or on specific products
type Voucher interface {
	VoucherID() [16]byte //
	Type() Voucher_Type  //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type Voucher_StorageServices interface {
	Save(v Voucher) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (v Voucher, numbers uint64, err protocol.Error)
}

type Voucher_Type uint32

const (
	Voucher_Type_Unset Voucher_Type = 0
	Voucher_Type_Area  Voucher_Type = (1 << iota)
	Voucher_Type_Campaign
	Voucher_Type_Category
	Voucher_Type_Duration
	Voucher_Type_Group
	Voucher_Type_Invoice
	Voucher_Type_POS
	Voucher_Type_Product
	Voucher_Type_SourceInvoice
	Voucher_Type_Usage    // limit to usage number
	Voucher_Type_UserType // limit to user type
	Voucher_Type_User     // limit to specific users
)
