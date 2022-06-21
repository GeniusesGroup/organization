/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
	"../libgo/time/utc"
)

// VoucherAnalyticUtcMinute indicate the domain record data fields.
type VoucherAnalyticUtcMinute interface {
	Minute() utc.MinuteElapsed            //
	Issued() uint64                       // Total
	IssuedAmount() protocol.AmountOfMoney // Total
	Time() protocol.Time                  // save time
}

type VoucherAnalyticUtcMinute_StorageServices interface {
	Save(va VoucherAnalyticUtcMinute) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (va VoucherAnalyticUtcMinute, numbers uint64, err protocol.Error)
}
