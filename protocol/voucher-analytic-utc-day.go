/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
	"../libgo/time/utc"
)

// VoucherAnalyticUtcDay indicate the domain record data fields.
type VoucherAnalyticUtcDay interface {
	Day() utc.DayElapsed                  //
	Issued() uint64                       // Total
	IssuedAmount() protocol.AmountOfMoney // Total
	Time() protocol.Time                  // save time
}

type VoucherAnalyticUtcDay_StorageServices interface {
	Save(va VoucherAnalyticUtcDay) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (va VoucherAnalyticUtcDay, numbers uint64, err protocol.Error)
}
