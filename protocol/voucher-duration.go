/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// VoucherDuration indicate the domain record data fields.
type VoucherDuration interface {
	VoucherID() [16]byte          // voucher domain
	Each() uint8                  // Each time use
	Epoch() VoucherDuration_Epoch //
	Duration() protocol.Duration  // from Epoch()
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type VoucherDuration_StorageServices interface {
	Save(vd VoucherDuration) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vd VoucherDuration, numbers uint64, err protocol.Error)
}

type VoucherDuration_Epoch uint8

const (
	VoucherDuration_Epoch_Unset VoucherDuration_Epoch = iota
	VoucherDuration_Epoch_IssueDate
	VoucherDuration_Epoch_LastUse
)
