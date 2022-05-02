/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type OrderDuration interface {
	OrderID() [16]byte           // order domain
	Epoch() OrderDuration_Epoch  //
	Duration() protocol.Duration // from Epoch()
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}

type OrderDuration_StorageServices interface {
	Save(od OrderDuration) (err protocol.Error)

	Count(orderID [16]byte) (numbers uint64, err protocol.Error)
	Get(orderID [16]byte, versionOffset uint64) (od OrderDuration, err protocol.Error)
	Last(orderID [16]byte) (od OrderDuration, numbers uint64, err protocol.Error)
}

type OrderDuration_Epoch uint8

const (
	OrderDuration_Epoch_Unset OrderDuration_Epoch = iota
	OrderDuration_Epoch_IssueDate
	OrderDuration_Epoch_LastDeal
)
